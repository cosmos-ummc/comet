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

type DeleteMeditationHandler struct {
	Model model.IModel
}

func (s *DeleteMeditationHandler) DeleteMeditation(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonMeditationResponse, error) {
	rslt, err := s.Model.DeleteMeditation(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeditationNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeditationToResponse(rslt)
	return resp, nil
}
