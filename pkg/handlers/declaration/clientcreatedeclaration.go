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
		ID:                 "",
		PatientID:          utility.RemoveZeroWidth(req.PatientId),
		PatientName:        "",
		PatientPhoneNumber: "",
		Result:             nil,
		Category:           "",
		Score:              0,
		Status:             0,
		SubmittedAt:        0,
		DoctorRemarks:      "",
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
