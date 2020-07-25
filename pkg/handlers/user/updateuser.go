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

type UpdateUserHandler struct {
	Model model.IModel
}

func (s *UpdateUserHandler) UpdateUser(ctx context.Context, req *pb.CommonUserRequest) (*pb.CommonUserResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	user := s.reqToUser(req)

	u, err := s.Model.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// special case: if questions JSON is updated, only update the questions JSON
	if req.Data.FinalQuestionsJson != "" && req.Data.PhoneNumber == "" {
		u.FinalQuestionsJSON = req.Data.FinalQuestionsJson
		v, err := s.Model.UpdateUser(ctx, u)
		if err != nil {
			if status.Code(err) == codes.Unknown {
				return nil, constants.UserNotFoundError
			}
			return nil, constants.InternalError
		}
		resp := utility.UserToResponse(v)
		return resp, nil
	}

	// special case: if chart is updated, only update the chart
	if req.Data.Chart != "" && req.Data.PhoneNumber == "" {
		u.Chart = req.Data.Chart
		v, err := s.Model.UpdateUser(ctx, u)
		if err != nil {
			if status.Code(err) == codes.Unknown {
				return nil, constants.UserNotFoundError
			}
			return nil, constants.InternalError
		}
		resp := utility.UserToResponse(v)
		return resp, nil
	}

	err = s.validateAndProcessReq(user)
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

	v, err := s.Model.UpdateUser(ctx, user)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.UserNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.UserToResponse(v)
	return resp, nil
}

func (s *UpdateUserHandler) reqToUser(req *pb.CommonUserRequest) *dto.User {
	user := &dto.User{
		ID:                 utility.RemoveZeroWidth(req.Id),
		Role:               utility.RemoveZeroWidth(req.Data.Role),
		DisplayName:        utility.RemoveZeroWidth(req.Data.DisplayName),
		PhoneNumber:        utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:              utility.RemoveZeroWidth(req.Data.Email),
		Disabled:           false,
		Password:           utility.RemoveZeroWidth(req.Data.Password),
		FinalQuestionsJSON: req.Data.FinalQuestionsJson,
		Chart:              req.Data.Chart,
	}
	return user
}

func (s *UpdateUserHandler) validateAndProcessReq(user *dto.User) error {
	user.DisplayName = utility.NormalizeName(user.DisplayName)
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
