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

type GetSwabHandler struct {
	Model model.IModel
}

func (s *GetSwabHandler) GetSwab(ctx context.Context, req *pb.CommonGetRequest, user *dto.User) (*pb.CommonSwabResponse, error) {
	swab, err := s.Model.GetSwab(ctx, req.Id, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.SwabNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.SwabToResponse(swab)
	return resp, nil
}
