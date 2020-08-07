package user

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

type UpdateUsersHandler struct {
	Model model.IModel
}

func (s *UpdateUsersHandler) UpdateUsers(ctx context.Context, req *pb.CommonUsersRequest) (*pb.CommonIdsResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	user := s.reqToUser(req)
	err := s.validateAndProcessReq(user)
	if err != nil {
		return nil, err
	}

	// get user
	u, err := s.Model.GetUser(ctx, req.Ids[0])
	if err != nil {
		return nil, err
	}

	if u.PhoneNumber != user.PhoneNumber {
		// check if phone number exist
		count, _, err := s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
			Item:  constants.PhoneNumber,
			Value: user.PhoneNumber,
		})
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, constants.PhoneNumberAlreadyExistError
		}
	}

	if u.Email != user.Email {
		// check if email exist
		count, _, err := s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
			Item:  constants.Email,
			Value: user.Email,
		})
		if err != nil {
			return nil, err
		}
		if count > 0 {
			return nil, constants.EmailAlreadyExistError
		}
	}

	ids, err := s.Model.UpdateUsers(ctx, user, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.UserNotFoundError
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateUsersHandler) reqToUser(req *pb.CommonUsersRequest) *dto.User {
	user := &dto.User{
		Role:        utility.RemoveZeroWidth(req.Data.Role),
		Name:        utility.RemoveZeroWidth(req.Data.Name),
		PhoneNumber: utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:       utility.RemoveZeroWidth(req.Data.Email),
		BlockList:   req.Data.BlockList,
		Password:    utility.RemoveZeroWidth(req.Data.Password),
	}
	return user
}

func (s *UpdateUsersHandler) validateAndProcessReq(user *dto.User) error {
	user.Name = utility.NormalizeName(user.Name)
	user.PhoneNumber = utility.NormalizePhoneNumber(user.PhoneNumber, "")
	user.Role = utility.NormalizeRole(user.Role)
	user.Email = utility.NormalizeEmail(user.Email)
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
