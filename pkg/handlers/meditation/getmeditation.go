package meditation

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetMeditationHandler struct {
	Model model.IModel
}

func (s *GetMeditationHandler) GetMeditation(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonMeditationResponse, error) {
	meditation, err := s.Model.GetMeditation(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeditationNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeditationToResponse(meditation)
	return resp, nil
}
