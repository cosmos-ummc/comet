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
		UserID:      uuid.NewV4().String(),
		Name:        req.Data.Name,
		PhoneNumber: utility.NormalizePhoneNumber(req.Data.PhoneNumber, ""),
		Email:       req.Data.Email,
		TakenSlots:  req.Data.TakenSlots,
	}

	// create user first then only can create consultant
	user := &dto.User{
		ID:          consultant.UserID,
		Role:        constants.Consultant,
		Name:        consultant.Name,
		PhoneNumber: consultant.PhoneNumber,
		Email:       consultant.Email,
		Password:    utility.RemoveZeroWidth(req.Data.Password),
	}

	// check if phone number exist
	count, _, err := s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.PhoneNumber,
		Value: user.PhoneNumber,
	})
	if count > 0 {
		return nil, constants.PhoneNumberAlreadyExistError
	}

	// check if email exist
	count, _, err = s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.Email,
		Value: user.Email,
	})
	if count > 0 {
		return nil, constants.EmailAlreadyExistError
	}

	_, err = s.Model.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// create consultant after user has been created
	rs, err := s.Model.CreateConsultant(ctx, consultant)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.ConsultantAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.ConsultantToResponse(rs)
	return resp, nil
}
