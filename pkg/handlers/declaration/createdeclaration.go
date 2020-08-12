package declaration

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateDeclarationHandler struct {
	Model model.IModel
}

func (s *CreateDeclarationHandler) CreateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest) (*pb.CommonDeclarationResponse, error) {
	declaration := &dto.Declaration{
		ID:               uuid.NewV4().String(),
		PatientID:        utility.RemoveZeroWidth(req.Data.PatientId),
		Result:           utility.PbToQuestions(req.Data.Result),
		Category:         req.Data.Category,
		Score:            req.Data.Score,
		Depression:       req.Data.Depression,
		Anxiety:          req.Data.Anxiety,
		Stress:           req.Data.Stress,
		DepressionStatus: req.Data.DepressionStatus,
		StressStatus:     req.Data.StressStatus,
		AnxietyStatus:    req.Data.AnxietyStatus,
		PtsdStatus:       req.Data.PtsdStatus,
		SubmittedAt:      req.Data.SubmittedAt,
		DoctorRemarks:    utility.RemoveZeroWidth(req.Data.DoctorRemarks),
	}

	// create new declaration
	rslt, err := s.Model.CreateDeclaration(ctx, declaration)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.DeclarationAlreadyExistError
		}
		return nil, constants.InternalError
	}

	resp := utility.DeclarationToResponse(rslt)
	return resp, nil
}
