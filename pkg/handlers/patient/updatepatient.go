package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"unicode/utf8"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdatePatientHandler struct {
	Model model.IModel
}

func (s *UpdatePatientHandler) UpdatePatient(ctx context.Context, req *pb.CommonPatientRequest) (*pb.CommonPatientResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}

	patient, err := s.validateAndProcessReq(req)
	if err != nil {
		logger.Log.Error("UpdatePatientHandler: " + err.Error())
		return nil, err
	}

	v, err := s.Model.UpdatePatient(ctx, patient)
	if err != nil {
		logger.Log.Error("UpdatePatientHandler: " + err.Error())
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		if status.Code(err) == codes.AlreadyExists {
			return nil, err
		}
		return nil, constants.InternalError
	}

	resp := utility.PatientToResponse(v)
	return resp, nil
}

func (s *UpdatePatientHandler) validateAndProcessReq(req *pb.CommonPatientRequest) (*dto.Patient, error) {
	if utf8.RuneCountInString(req.Data.Remarks) > 200 {
		return nil, constants.RemarksTooLongError
	}

	patient := &dto.Patient{
		ID:                 utility.RemoveZeroWidth(req.Id),
		Name:               utility.RemoveZeroWidth(req.Data.Name),
		TelegramID:         utility.RemoveZeroWidth(req.Data.TelegramId),
		PhoneNumber:        utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:              utility.RemoveZeroWidth(req.Data.Email),
		Status:             req.Data.Status,
		Remarks:            utility.RemoveZeroWidth(req.Data.Remarks),
		HomeAddress:        req.Data.HomeAddress,
		IsolationAddress:   utility.RemoveZeroWidth(req.Data.IsolationAddress),
		DaySinceMonitoring: req.Data.DaySinceMonitoring,
		HasCompleted:       req.Data.HasCompleted,
		MentalStatus:       req.Data.MentalStatus,
		Type:               req.Data.Type,
		SwabDate:           req.Data.SwabDate,
		SwabResult:         req.Data.SwabResult,
	}

	patient.PhoneNumber = utility.NormalizePhoneNumber(patient.PhoneNumber, "")
	patient.ID = utility.NormalizeID(patient.ID)
	patient.Name = utility.NormalizeName(patient.Name)
	if patient.PhoneNumber == "" {
		return nil, constants.InvalidPhoneNumberError
	}
	if patient.Name == "" {
		return nil, constants.InvalidPatientNameError
	}
	return patient, nil
}
