package feed

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetFeedsHandler struct {
	Model model.IModel
}

func (s *ClientGetFeedsHandler) GetFeeds(ctx context.Context, req *pb.ClientGetFeedsRequest) (*pb.CommonFeedsResponse, error) {
	// if id is 1, do uncategorized query
	if req.Id == "1" {
		var filter map[string]interface{}
		filter = map[string]interface{}{
			constants.Type: constants.NoMentalIssue,
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
	var resultFeeds []*dto.Feed

	// default is depression
	if patient.MentalStatus == constants.NoMentalIssue {
		patient.MentalStatus = constants.Depression
	}

	// do first query
	_, feeds, err := s.Model.QueryFeeds(ctx, nil, nil, map[string]interface{}{
		constants.Type: patient.MentalStatus,
	})
	if err != nil {
		return nil, err
	}
	utility.ShuffleFeeds(feeds)
	if len(feeds) > 3 {
		feeds = feeds[0:3]
	}
	resultFeeds = append(resultFeeds, feeds...)

	// do second query
	if patient.MentalStatus != constants.Depression {
		_, feeds, err = s.Model.QueryFeeds(ctx, nil, nil, map[string]interface{}{
			constants.Type: constants.Depression,
		})
		if err != nil {
			return nil, err
		}
		utility.ShuffleFeeds(feeds)
		if len(feeds) > 3 {
			feeds = feeds[0:3]
		}
		resultFeeds = append(resultFeeds, feeds...)
	}

	// do third query
	if patient.MentalStatus != constants.Anxiety {
		_, feeds, err = s.Model.QueryFeeds(ctx, nil, nil, map[string]interface{}{
			constants.Type: constants.Anxiety,
		})
		if err != nil {
			return nil, err
		}
		utility.ShuffleFeeds(feeds)
		if len(feeds) > 3 {
			feeds = feeds[0:3]
		}
		resultFeeds = append(resultFeeds, feeds...)
	}

	// do forth query
	if patient.MentalStatus != constants.Stress {
		_, feeds, err = s.Model.QueryFeeds(ctx, nil, nil, map[string]interface{}{
			constants.Type: constants.Stress,
		})
		if err != nil {
			return nil, err
		}
		utility.ShuffleFeeds(feeds)
		if len(feeds) > 3 {
			feeds = feeds[0:3]
		}
		resultFeeds = append(resultFeeds, feeds...)
	}

	// do fifth query
	if patient.MentalStatus != constants.PTSD {
		_, feeds, err = s.Model.QueryFeeds(ctx, nil, nil, map[string]interface{}{
			constants.Type: constants.PTSD,
		})
		if err != nil {
			return nil, err
		}
		utility.ShuffleFeeds(feeds)
		if len(feeds) > 3 {
			feeds = feeds[0:3]
		}
		resultFeeds = append(resultFeeds, feeds...)
	}

	resp := utility.FeedsToResponse(resultFeeds)
	resp.Total = int64(len(resultFeeds))
	return resp, nil
}
