package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"context"
)

type LoginHandler struct {
	Model model.IModel
}

func (s *LoginHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user, err := s.Model.Login(ctx, req.Email, req.Password)
	if err != nil {
		logger.Log.Error("Login: " + err.Error())
		return nil, constants.InvalidPasswordVerifyError
	}
	resp := s.userToResp(user)
	return resp, nil
}

func (s *LoginHandler) userToResp(user *dto.User) *pb.LoginResponse {
	return &pb.LoginResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
		DisplayName:  user.Name,
		Role:         user.Role,
		Id:           user.ID,
	}
}
