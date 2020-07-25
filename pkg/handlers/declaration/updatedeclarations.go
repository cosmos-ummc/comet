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

type UpdateDeclarationsHandler struct {
	Model model.IModel
}

func (s *UpdateDeclarationsHandler) UpdateDeclarations(ctx context.Context, req *pb.CommonDeclarationsRequest, user *dto.User) (*pb.CommonIdsResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	declaration := s.reqToDeclaration(req)
	err := s.processReq(declaration)
	if err != nil {
		return nil, err
	}

	ids, err := s.Model.UpdateDeclarations(ctx, declaration, req.Ids, constants.UserPatientMap[user.Role])
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

func (s *UpdateDeclarationsHandler) processReq(declaration *dto.Declaration) error {
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
