package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreatePatientHandler struct {
	Model model.IModel
}

func (s *CreatePatientHandler) CreatePatient(ctx context.Context, req *pb.CommonPatientRequest) (*pb.CommonPatientResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}

	patient, err := s.validateAndProcessReq(req)
	if err != nil {
		logger.Log.Error("CreatePatientHandler: " + err.Error())
		return nil, err
	}

	v, err := s.Model.CreatePatient(ctx, patient)
	if err != nil {
		logger.Log.Error("CreatePatientHandler: " + err.Error())
		if status.Code(err) == codes.AlreadyExists {
			return nil, err
		}
		return nil, constants.InternalError
	}

	resp := utility.PatientToResponse(v)
	return resp, nil
}

func (s *CreatePatientHandler) validateAndProcessReq(req *pb.CommonPatientRequest) (*dto.Patient, error) {
	patient := &dto.Patient{
		ID:               utility.RemoveZeroWidth(req.Id),
		UserID:           req.Data.UserId,
		Name:             utility.RemoveZeroWidth(req.Data.Name),
		TelegramID:       utility.RemoveZeroWidth(req.Data.TelegramId),
		PhoneNumber:      utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:            utility.RemoveZeroWidth(req.Data.Email),
		Status:           req.Data.Status,
		LastDassTime:     req.Data.LastDassTime,
		LastIesrTime:     req.Data.LastIesrTime,
		LastDassResult:   req.Data.LastDassResult,
		LastIesrResult:   req.Data.LastIesrResult,
		Remarks:          req.Data.Remarks,
		HomeAddress:      req.Data.HomeAddress,
		IsolationAddress: utility.RemoveZeroWidth(req.Data.IsolationAddress),
	}

	patient.PhoneNumber = utility.NormalizePhoneNumber(patient.PhoneNumber, "")
	patient.ID = utility.NormalizeID(patient.ID)
	patient.Name = utility.NormalizeName(patient.Name)
	if patient.PhoneNumber == "" {
		return nil, constants.InvalidPhoneNumberError
	}
	if patient.ID == "" {
		return nil, constants.InvalidPatientIDError
	}
	if patient.Name == "" {
		return nil, constants.InvalidPatientNameError
	}
	if patient.Status < 1 || patient.Status > 6 {
		return nil, constants.InvalidPatientStatusError
	}
	return patient, nil
}
