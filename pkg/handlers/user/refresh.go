package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"context"
	"strings"

	"google.golang.org/grpc/metadata"
)

type RefreshHandler struct {
	Model model.IModel
}

func (s *RefreshHandler) Refresh(ctx context.Context) (*pb.RefreshResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, constants.MetadataNotFoundError
	}
	tokenSlice := md.Get("authorization")
	user, err := s.Model.Refresh(ctx, strings.Join(tokenSlice, " "))
	if err != nil {
		logger.Log.Error("Refresh: " + err.Error())
		return nil, err
	}

	resp := s.userToResp(user)
	return resp, nil
}

func (s *RefreshHandler) userToResp(user *dto.User) *pb.RefreshResponse {
	return &pb.RefreshResponse{
		AccessToken:  user.AccessToken,
		RefreshToken: user.RefreshToken,
	}
}
