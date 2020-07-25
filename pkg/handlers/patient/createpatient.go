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

func (s *CreatePatientHandler) CreatePatient(ctx context.Context, req *pb.CommonPatientRequest, user *dto.User) (*pb.CommonPatientResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}

	patient, err := s.validateAndProcessReq(req)
	if err != nil {
		logger.Log.Error("CreatePatientHandler: " + err.Error())
		return nil, err
	}

	// set patient type based on patient scope
	if scope := constants.UserPatientMap[user.Role]; scope != constants.AllPatients {
		patient.Type = scope
	}

	v, err := s.Model.CreatePatient(ctx, patient, user)
	if err != nil {
		logger.Log.Error("CreatePatientHandler: " + err.Error())
		if status.Code(err) == codes.AlreadyExists {
			return nil, err
		}
		return nil, constants.InternalError
	}

	// create swab for the patient
	swab := &dto.Swab{
		PatientID:           utility.RemoveZeroWidth(patient.ID),
		Status:              req.Data.SwabStatus,
		Date:                utility.RemoveZeroWidth(req.Data.SwabDate),
		Location:            req.Data.SwabLocation,
		IsOtherSwabLocation: req.Data.IsOtherSwabLocation,
	}
	_, err = s.Model.CreateSwab(ctx, swab, constants.UserPatientMap[user.Role], user)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.SwabAlreadyExistError
		}
		return nil, constants.InternalError
	}

	resp := utility.PatientToResponse(v)
	return resp, nil
}

func (s *CreatePatientHandler) validateAndProcessReq(req *pb.CommonPatientRequest) (*dto.Patient, error) {
	patient := &dto.Patient{
		ID:               utility.RemoveZeroWidth(req.Id),
		Name:             utility.RemoveZeroWidth(req.Data.Name),
		TelegramID:       utility.RemoveZeroWidth(req.Data.TelegramId),
		PhoneNumber:      utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:            utility.RemoveZeroWidth(req.Data.Email),
		Status:           req.Data.Status,
		Type:             req.Data.Type,
		ExposureDate:     utility.RemoveZeroWidth(req.Data.ExposureDate),
		ExposureSource:   utility.RemoveZeroWidth(req.Data.ExposureSource),
		RegistrationNum:  utility.RemoveZeroWidth(req.Data.RegistrationNum),
		AlternateContact: utility.RemoveZeroWidth(req.Data.AlternateContact),
		IsolationAddress: utility.RemoveZeroWidth(req.Data.IsolationAddress),
		SymptomDate:      utility.RemoveZeroWidth(req.Data.SymptomDate),
		Localization:     req.Data.Localization,
		FeverStartDate:   utility.RemoveZeroWidth(req.Data.FeverStartDate),
		HomeAddress:      req.Data.HomeAddress,
		IsSameAddress:    req.Data.IsSameAddress,
	}

	patient.PhoneNumber = utility.NormalizePhoneNumber(patient.PhoneNumber, "")
	patient.ID = utility.NormalizeID(patient.ID)
	patient.Name = utility.NormalizeName(patient.Name)
	var err error
	if patient.FeverStartDate != "" {
		patient.FeverStartDate, err = utility.NormalizeDate(patient.FeverStartDate)
		if err != nil {
			return nil, constants.InvalidDateError
		}
	}
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
	if patient.Localization > 4 {
		return nil, constants.InvalidLanguageError
	}
	if patient.Type == 0 {
		return nil, constants.InvalidPatientTypeError
	}
	if patient.IsSameAddress {
		patient.IsolationAddress = patient.HomeAddress
	}
	if patient.HomeAddress == "" || patient.IsolationAddress == "" {
		return nil, constants.InvalidAddressError
	}
	return patient, nil
}
