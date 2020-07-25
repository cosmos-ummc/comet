package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUsersHandler struct {
	Model model.IModel
}

func (s *DeleteUsersHandler) DeleteUsers(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	var ids []string

	// remove users from firebase auth
	for _, id := range req.Ids {
		u, err := s.Model.DeleteUser(ctx, id)
		if err != nil {
			if status.Code(err) == codes.Unknown {
				return nil, constants.UserNotFoundError
			}
			return nil, constants.InternalError
		}

		// add user into deleted user IDs
		ids = append(ids, u.ID)
	}

	return &pb.CommonIdsResponse{Data: ids}, nil
}
