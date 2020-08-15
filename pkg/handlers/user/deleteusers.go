package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUsersHandler struct {
	Model model.IModel
}

func (s *DeleteUsersHandler) DeleteUsers(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	res := []string{}
	for _, id := range req.Ids {
		split := strings.Split(id, ",")
		res = append(res, split...)
	}

	// remove users from firebase auth
	for _, id := range res {
		_, err := s.Model.DeleteUser(ctx, id)
		if err != nil {
			if status.Code(err) == codes.Unknown {
				return nil, constants.UserNotFoundError
			}
			return nil, constants.InternalError
		}
	}

	return &pb.CommonIdsResponse{Data: req.Ids}, nil
}
