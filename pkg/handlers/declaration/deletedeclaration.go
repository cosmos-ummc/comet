package declaration

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

type DeleteDeclarationHandler struct {
	Model model.IModel
}

func (s *DeleteDeclarationHandler) DeleteDeclaration(ctx context.Context, req *pb.CommonDeleteRequest, user *dto.User) (*pb.CommonDeclarationResponse, error) {
	rslt, err := s.Model.DeleteDeclaration(ctx, req.Id, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.DeclarationNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.DeclarationToResponse(rslt)
	return resp, nil
}
