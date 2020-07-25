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

type UpdateDeclarationHandler struct {
	Model model.IModel
}

func (s *UpdateDeclarationHandler) UpdateDeclaration(ctx context.Context, req *pb.CommonDeclarationRequest, user *dto.User) (*pb.CommonDeclarationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	declaration := s.reqToDeclaration(req)

	err := s.processReq(declaration)
	if err != nil {
		return nil, err
	}

	// update declaration
	v, err := s.Model.UpdateDeclaration(ctx, declaration, constants.UserPatientMap[user.Role])
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
		PatientID:     utility.RemoveZeroWidth(req.Data.PatientId),
		Cough:         req.Data.Cough,
		Throat:        req.Data.Throat,
		Fever:         req.Data.Fever,
		Breathe:       req.Data.Breathe,
		Chest:         req.Data.Chest,
		Blue:          req.Data.Blue,
		Drowsy:        req.Data.Drowsy,
		HasSymptom:    req.Data.HasSymptom,
		SubmittedAt:   req.Data.SubmittedAt,
		CallingStatus: req.Data.CallingStatus,
		Date:          utility.RemoveZeroWidth(req.Data.Date),
		DoctorRemarks: utility.RemoveZeroWidth(req.Data.DoctorRemarks),
		Loss:          req.Data.Loss,
	}
	return declaration
}

func (s *UpdateDeclarationHandler) processReq(declaration *dto.Declaration) error {
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
