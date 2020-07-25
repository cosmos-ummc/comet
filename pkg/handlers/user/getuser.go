package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetUserHandler struct {
	Model model.IModel
}

func (s *GetUserHandler) GetUser(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonUserResponse, error) {
	user, err := s.Model.GetUser(ctx, req.Id)
	if err != nil {
		logger.Log.Error("GetUserHandler: " + err.Error())
		if status.Code(err) == codes.Unknown {
			return nil, constants.UserNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.UserToResponse(user)

	return resp, nil
}
