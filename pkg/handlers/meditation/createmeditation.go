package meditation

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateMeditationHandler struct {
	Model model.IModel
}

func (s *CreateMeditationHandler) CreateMeditation(ctx context.Context, req *pb.CommonMeditationRequest) (*pb.CommonMeditationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	meditation := &dto.Meditation{
		ID:   uuid.NewV4().String(),
		Link: req.Data.Link,
	}

	rslt, err := s.Model.CreateMeditation(ctx, meditation)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.MeditationAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeditationToResponse(rslt)
	return resp, nil
}
