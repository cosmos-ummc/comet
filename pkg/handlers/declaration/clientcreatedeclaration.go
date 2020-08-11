package declaration

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"
)

type ClientCreateDeclarationHandler struct {
	Model model.IModel
}

func (s *ClientCreateDeclarationHandler) ClientCreateDeclaration(ctx context.Context, req *pb.ClientCreateDeclarationRequest) (*pb.ClientCreateDeclarationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	req.Data.Id = uuid.NewV4().String()
	declaration := utility.PbToDeclaration(req.Data)

	_, err := s.Model.ClientCreateDeclaration(ctx, declaration)
	if err != nil {
		return nil, err
	}
	resp := &pb.ClientCreateDeclarationResponse{
		HasSymptom: constants.DeclarationSevere,
	}
	return resp, nil
}
