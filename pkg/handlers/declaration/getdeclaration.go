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

type GetDeclarationHandler struct {
	Model model.IModel
}

func (s *GetDeclarationHandler) GetDeclaration(ctx context.Context, req *pb.CommonGetRequest, user *dto.User) (*pb.CommonDeclarationResponse, error) {
	declaration, err := s.Model.GetDeclaration(ctx, req.Id, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.DeclarationNotFoundError
		}
		return nil, constants.InternalError
	}
	return utility.DeclarationToResponse(declaration), nil
}
