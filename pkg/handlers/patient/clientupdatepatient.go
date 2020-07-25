package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientUpdatePatientHandler struct {
	Model model.IModel
}

func (s *ClientUpdatePatientHandler) ClientUpdatePatient(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*empty.Empty, error) {
	if utf8.RuneCountInString(req.Remarks) > 200 {
		return &empty.Empty{}, constants.RemarksTooLongError
	}

	patient := s.processReq(req)

	_, err := s.Model.ClientUpdatePatient(ctx, patient)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}
	return &empty.Empty{}, nil
}

func (s *ClientUpdatePatientHandler) processReq(req *pb.ClientUpdatePatientRequest) *dto.Patient {
	patient := &dto.Patient{
		ID:            utility.RemoveZeroWidth(req.Id),
		Name:          utility.RemoveZeroWidth(req.Name),
		TelegramID:    utility.RemoveZeroWidth(req.TelegramId),
		PhoneNumber:   utility.RemoveZeroWidth(req.PhoneNumber),
		Email:         utility.RemoveZeroWidth(req.Email),
		Status:        req.Status,
		LastDeclared:  req.LastDeclared,
		Remarks:       utility.RemoveZeroWidth(req.Remarks),
		Localization:  req.Localization,
		Episode:       req.Episode,
		Consent:       req.Consent,
		PrivacyPolicy: req.PrivacyPolicy,
	}

	patient.PhoneNumber = utility.NormalizePhoneNumber(patient.PhoneNumber, "")
	patient.ID = utility.NormalizeID(patient.ID)
	patient.Name = utility.NormalizeName(patient.Name)

	return patient
}
