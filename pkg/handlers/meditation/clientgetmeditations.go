package meditation

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetMeditationsHandler struct {
	Model model.IModel
}

func (s *ClientGetMeditationsHandler) GetMeditations(ctx context.Context, req *pb.ClientGetMeditationsRequest) (*pb.CommonMeditationsResponse, error) {
	// do query
	_, meditations, err := s.Model.QueryMeditations(ctx, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	// shuffle meditations and select 2
	utility.ShuffleMeditations(meditations)
	if len(meditations) > 2 {
		meditations = meditations[0:2]
	}

	resp := utility.MeditationsToResponse(meditations)
	resp.Total = int64(len(meditations))
	return resp, nil
}
