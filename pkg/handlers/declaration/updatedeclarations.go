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

type UpdateDeclarationsHandler struct {
	Model model.IModel
}

func (s *UpdateDeclarationsHandler) UpdateDeclarations(ctx context.Context, req *pb.CommonDeclarationsRequest) (*pb.CommonIdsResponse, error) {
	if req.Data == nil || len(req.Ids) != 1 {
		return nil, constants.InvalidArgumentError
	}
	declaration := s.reqToDeclaration(req)

	ids, err := s.Model.UpdateDeclarations(ctx, declaration, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.DeclarationNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateDeclarationsHandler) reqToDeclaration(req *pb.CommonDeclarationsRequest) *dto.Declaration {
	declaration := &dto.Declaration{
		ID:            utility.RemoveZeroWidth(req.Ids[0]),
		Score:         req.Data.Score,
		Status:        req.Data.Status,
		DoctorRemarks: utility.RemoveZeroWidth(req.Data.DoctorRemarks),
	}
	return declaration
}
