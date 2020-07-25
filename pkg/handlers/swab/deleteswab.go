package swab

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

type DeleteSwabHandler struct {
	Model model.IModel
}

func (s *DeleteSwabHandler) DeleteSwab(ctx context.Context, req *pb.CommonDeleteRequest, user *dto.User) (*pb.CommonSwabResponse, error) {
	rslt, err := s.Model.DeleteSwab(ctx, req.Id, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.SwabNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.SwabToResponse(rslt)
	return resp, nil
}
