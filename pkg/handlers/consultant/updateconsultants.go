package consultant

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

type UpdateConsultantsHandler struct {
	Model model.IModel
}

func (s *UpdateConsultantsHandler) UpdateConsultants(ctx context.Context, req *pb.CommonConsultantsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	consultant := s.reqToConsultant(req)

	ids, err := s.Model.UpdateConsultants(ctx, consultant, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ConsultantNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateConsultantsHandler) reqToConsultant(req *pb.CommonConsultantsRequest) *dto.Consultant {
	consultant := &dto.Consultant{
		ID:          req.Ids[0],
		UserID:      req.Data.UserId,
		Name:        req.Data.Name,
		PhoneNumber: utility.NormalizePhoneNumber(req.Data.PhoneNumber, ""),
		Email:       req.Data.Email,
	}
	return consultant
}
