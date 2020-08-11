package meditation

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateMeditationHandler struct {
	Model model.IModel
}

func (s *UpdateMeditationHandler) UpdateMeditation(ctx context.Context, req *pb.CommonMeditationRequest) (*pb.CommonMeditationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	meditation := s.reqToMeditation(req)

	v, err := s.Model.UpdateMeditation(ctx, meditation)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeditationNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeditationToResponse(v)
	return resp, nil
}

func (s *UpdateMeditationHandler) reqToMeditation(req *pb.CommonMeditationRequest) *dto.Meditation {
	meditation := &dto.Meditation{
		ID:   utility.RemoveZeroWidth(req.Id),
		Link: req.Data.Link,
	}
	return meditation
}
