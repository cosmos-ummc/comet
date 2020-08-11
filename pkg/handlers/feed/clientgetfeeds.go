package feed

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetFeedsHandler struct {
	Model model.IModel
}

func (s *ClientGetFeedsHandler) GetFeeds(ctx context.Context, req *pb.ClientGetFeedsRequest) (*pb.CommonFeedsResponse, error) {
	// get patient
	total, patients, err := s.Model.QueryPatients(ctx, nil, nil, map[string]interface{}{
		constants.UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if total == 0 {
		return &pb.CommonFeedsResponse{}, nil
	}

	// setup filter for patient mental status
	patient := patients[0]
	var filter map[string]interface{}
	if patient.MentalStatus != constants.NoMentalIssue {
		filter = map[string]interface{}{
			constants.Type: patient.MentalStatus,
		}
	}

	// do query
	_, feeds, err := s.Model.QueryFeeds(ctx, nil, nil, filter)
	if err != nil {
		return nil, err
	}

	// shuffle feeds
	utility.ShuffleFeeds(feeds)
	if len(feeds) > 3 {
		feeds = feeds[0:3]
	}

	resp := utility.FeedsToResponse(feeds)
	resp.Total = int64(len(feeds))
	return resp, nil
}
