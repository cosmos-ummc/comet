package declaration

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"unicode/utf8"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateDeclarationHandler struct {
	Model model.IModel
}

func (s *CreateDeclarationHandler) CreateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest, user *dto.User) (*pb.CommonDeclarationResponse, error) {
	declaration := &dto.Declaration{
		PatientID:     utility.RemoveZeroWidth(req.Data.PatientId),
		Cough:         req.Data.Cough,
		Throat:        req.Data.Throat,
		Fever:         req.Data.Fever,
		Breathe:       req.Data.Breathe,
		Chest:         req.Data.Chest,
		Blue:          req.Data.Blue,
		Drowsy:        req.Data.Drowsy,
		SubmittedAt:   req.Data.SubmittedAt,
		Date:          utility.RemoveZeroWidth(req.Data.Date),
		DoctorRemarks: utility.RemoveZeroWidth(req.Data.DoctorRemarks),
		Loss:          req.Data.Loss,
	}
	err := s.processReq(declaration)
	if err != nil {
		return nil, err
	}

	// create new declaration
	rslt, err := s.Model.CreateDeclaration(ctx, declaration, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.DeclarationAlreadyExistError
		}
		return nil, constants.InternalError
	}

	resp := utility.DeclarationToResponse(rslt)
	return resp, nil
}

func (s *CreateDeclarationHandler) processReq(declaration *dto.Declaration) error {
	declaration.PatientID = utility.NormalizeID(declaration.PatientID)
	var err error
	declaration.Date, err = utility.NormalizeDate(declaration.Date)
	if err != nil {
		return err
	}
	if utf8.RuneCountInString(declaration.DoctorRemarks) > 200 {
		return constants.RemarksTooLongError
	}
	return nil
}
