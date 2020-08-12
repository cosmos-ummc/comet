package declaration

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"
	"time"
)

type ClientCreateDeclarationHandler struct {
	Model model.IModel
}

func (s *ClientCreateDeclarationHandler) ClientCreateDeclaration(ctx context.Context, req *pb.ClientCreateDeclarationRequest) (*pb.ClientCreateDeclarationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	req.Data.Id = uuid.NewV4().String()
	req.Data.SubmittedAt = utility.TimeToMilli(utility.MalaysiaTime(time.Now()))
	req.Data.PatientId = req.PatientId
	declaration := utility.PbToDeclaration(req.Data)

	d, err := s.Model.ClientCreateDeclaration(ctx, declaration)
	if err != nil {
		return nil, err
	}

	hasSymptom := false
	if d.StressStatus == constants.DeclarationSevere || d.StressStatus == constants.DeclarationExtremelySevere ||
		d.DepressionStatus == constants.DeclarationSevere || d.DepressionStatus == constants.DeclarationExtremelySevere ||
		d.AnxietyStatus == constants.DeclarationSevere || d.AnxietyStatus == constants.DeclarationExtremelySevere ||
		d.PtsdStatus == constants.DeclarationSevere || d.PtsdStatus == constants.DeclarationExtremelySevere {
		hasSymptom = true
	}

	resp := &pb.ClientCreateDeclarationResponse{
		HasSymptom: hasSymptom,
	}
	return resp, nil
}
