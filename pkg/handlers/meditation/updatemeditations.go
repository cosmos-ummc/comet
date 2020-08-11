package meditation

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateMeditationsHandler struct {
	Model model.IModel
}

func (s *UpdateMeditationsHandler) UpdateMeditations(ctx context.Context, req *pb.CommonMeditationsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	meditation := s.reqToMeditation(req)

	ids, err := s.Model.UpdateMeditations(ctx, meditation, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeditationNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateMeditationsHandler) reqToMeditation(req *pb.CommonMeditationsRequest) *dto.Meditation {
	meditation := &dto.Meditation{
		Link: req.Data.Link,
	}
	return meditation
}
