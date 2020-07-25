package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUserHandler struct {
	Model model.IModel
}

func (s *DeleteUserHandler) DeleteUser(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonUserResponse, error) {
	rslt, err := s.Model.DeleteUser(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.UserNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.UserToResponse(rslt)
	return resp, nil
}
