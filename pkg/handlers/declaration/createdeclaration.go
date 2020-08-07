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
		ID:            uuid.NewV4().String(),
		PatientID:     utility.RemoveZeroWidth(req.Data.PatientId),
		Result:        utility.PbToQuestions(req.Data.Result),
		Category:      req.Data.Category,
		Score:         req.Data.Score,
		Status:        req.Data.Status,
		SubmittedAt:   req.Data.SubmittedAt,
		DoctorRemarks: utility.RemoveZeroWidth(req.Data.DoctorRemarks),
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
