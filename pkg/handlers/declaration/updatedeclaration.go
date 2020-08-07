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

type UpdateDeclarationHandler struct {
	Model model.IModel
}

func (s *UpdateDeclarationHandler) UpdateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	declaration := s.reqToDeclaration(req)

	// update declaration
	v, err := s.Model.UpdateDeclaration(ctx, declaration)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.DeclarationNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.DeclarationToResponse(v)
	return resp, nil
}

func (s *UpdateDeclarationHandler) reqToDeclaration(req *pb.CommonDeclarationRequest) *dto.Declaration {
	declaration := &dto.Declaration{
		ID:            utility.RemoveZeroWidth(req.Id),
		Score:         req.Data.Score,
		Status:        req.Data.Status,
		DoctorRemarks: utility.RemoveZeroWidth(req.Data.DoctorRemarks),
	}
	return declaration
}
