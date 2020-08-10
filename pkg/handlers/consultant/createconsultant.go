package consultant

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

type CreateConsultantHandler struct {
	Model model.IModel
}

func (s *CreateConsultantHandler) CreateConsultant(ctx context.Context, req *pb.CommonConsultantRequest) (*pb.CommonConsultantResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	consultant := &dto.Consultant{
		ID:          uuid.NewV4().String(),
		UserID:      req.Data.UserId,
		Name:        req.Data.Name,
		PhoneNumber: utility.NormalizePhoneNumber(req.Data.PhoneNumber, ""),
		Email:       req.Data.Email,
		TakenSlots:  req.Data.TakenSlots,
	}

	rslt, err := s.Model.CreateConsultant(ctx, consultant)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.ConsultantAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.ConsultantToResponse(rslt)
	return resp, nil
}
