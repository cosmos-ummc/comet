package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"context"

	"github.com/golang/protobuf/ptypes/empty"
)

type UpdatePasswordHandler struct {
	Model model.IModel
}

func (s *UpdatePasswordHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*empty.Empty, error) {
	// get user id by token
	id, err := s.Model.GetUserIDByToken(ctx, req.Token)
	if err != nil {
		logger.Log.Error("UpdatePassword: " + err.Error())
		return nil, constants.UserNotFoundError
	}

	// prepare user payload
	user := &dto.User{
		ID:       id,
		Password: req.Password,
	}

	// update user password
	_, err = s.Model.UpdateUserPassword(ctx, user)
	if err != nil {
		logger.Log.Error("UpdatePassword: " + err.Error())
		return nil, constants.InternalError
	}

	// revoke all user tokens
	err = s.Model.RevokeTokensByUserID(ctx, id)
	if err != nil {
		logger.Log.Error("UpdatePassword: " + err.Error())
		return nil, constants.InternalError
	}

	return &empty.Empty{}, nil
}
