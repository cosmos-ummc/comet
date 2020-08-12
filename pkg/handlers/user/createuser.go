package user

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

type CreateUserHandler struct {
	Model model.IModel
}

func (s *CreateUserHandler) CreateUser(ctx context.Context, req *pb.CommonUserRequest) (*pb.CommonUserResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	user := &dto.User{
		ID:               uuid.NewV4().String(),
		Role:             constants.Superuser,
		Name:             req.Data.Name,
		PhoneNumber:      utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:            utility.RemoveZeroWidth(req.Data.Email),
		Password:         utility.RemoveZeroWidth(req.Data.Password),
		BlockList:        req.Data.BlockList,
		Visible:          req.Data.Visible,
		Disabled:         req.Data.Disabled,
		NotFirstTimeChat: req.Data.NotFirstTimeChat,
		InvitedToMeeting: req.Data.InvitedToMeeting,
	}
	err := s.validateAndProcessReq(user)
	if err != nil {
		return nil, err
	}

	// check if phone number exist
	count, _, err := s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.PhoneNumber,
		Value: user.PhoneNumber,
	}, false)
	if count > 0 {
		return nil, constants.PhoneNumberAlreadyExistError
	}

	// check if email exist
	count, _, err = s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.Email,
		Value: user.Email,
	}, false)
	if count > 0 {
		return nil, constants.EmailAlreadyExistError
	}

	rslt, err := s.Model.CreateUser(ctx, user)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.UserAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.UserToResponse(rslt)
	return resp, nil
}

func (s *CreateUserHandler) validateAndProcessReq(user *dto.User) error {
	user.Name = utility.NormalizeName(user.Name)
	user.PhoneNumber = utility.NormalizePhoneNumber(user.PhoneNumber, "")
	user.Email = utility.NormalizeEmail(user.Email)
	user.Role = utility.NormalizeRole(user.Role)
	if len(user.Password) < 6 {
		return constants.InvalidPasswordError
	}
	if user.PhoneNumber == "" {
		return constants.InvalidPhoneNumberError
	}
	if user.Email == "" {
		return constants.InvalidEmailError
	}
	if user.Role == "" {
		return constants.InvalidRoleError
	}

	return nil
}
