package declaration

import (
	pb "comet/pkg/api"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientCreateDeclarationHandler struct {
	Model model.IModel
}

func (s *ClientCreateDeclarationHandler) ClientCreateDeclaration(ctx context.Context, req *pb.ClientCreateDeclarationRequest) (*pb.ClientCreateDeclarationResponse, error) {
	declaration := &dto.Declaration{
		PatientID: utility.RemoveZeroWidth(req.PatientId),
		Cough:     req.Cough,
		Throat:    req.Throat,
		Fever:     req.Fever,
		Breathe:   req.Breathe,
		Chest:     req.Chest,
		Blue:      req.Blue,
		Drowsy:    req.Drowsy,
		Loss:      req.Loss,
	}

	rslt, err := s.Model.ClientCreateDeclaration(ctx, declaration)
	if err != nil {
		return nil, err
	}
	resp := &pb.ClientCreateDeclarationResponse{
		HasSymptom: rslt.HasSymptom,
	}
	return resp, nil
}
