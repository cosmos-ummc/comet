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

type UpdateConsultantHandler struct {
	Model model.IModel
}

func (s *UpdateConsultantHandler) UpdateConsultant(ctx context.Context, req *pb.CommonConsultantRequest) (*pb.CommonConsultantResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	consultant := s.reqToConsultant(req)

	v, err := s.Model.UpdateConsultant(ctx, consultant)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ConsultantNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ConsultantToResponse(v)
	return resp, nil
}

func (s *UpdateConsultantHandler) reqToConsultant(req *pb.CommonConsultantRequest) *dto.Consultant {
	consultant := &dto.Consultant{
		ID:          utility.RemoveZeroWidth(req.Id),
		UserID:      req.Data.UserId,
		Name:        req.Data.Name,
		PhoneNumber: utility.NormalizePhoneNumber(req.Data.PhoneNumber, ""),
		Email:       req.Data.Email,
		TakenSlots:  req.Data.TakenSlots,
	}
	return consultant
}
