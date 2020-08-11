package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"
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

	// create user first then only can create patient
	user := &dto.User{
		ID:          patient.UserID,
		Role:        constants.Consultant,
		Name:        patient.Name,
		PhoneNumber: patient.PhoneNumber,
		Email:       patient.Email,
		Password:    utility.RemoveZeroWidth(req.Data.Password),
	}

	// check if phone number exist
	count, _, err := s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.PhoneNumber,
		Value: user.PhoneNumber,
	}, false)
	if count > 0 {
		return nil, constants.PhoneNumberAlreadyExistError
	}

	// check if email exist
	count, _, err = s.Model.QueryUsers(ctx, nil, nil, &dto.FilterData{
		Item:  constants.Email,
		Value: user.Email,
	}, false)
	if count > 0 {
		return nil, constants.EmailAlreadyExistError
	}

	_, err = s.Model.CreateUser(ctx, user)
	if err != nil {
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
		ID:                 utility.RemoveZeroWidth(req.Id),
		UserID:             uuid.NewV4().String(),
		Name:               utility.RemoveZeroWidth(req.Data.Name),
		TelegramID:         utility.RemoveZeroWidth(req.Data.TelegramId),
		PhoneNumber:        utility.RemoveZeroWidth(req.Data.PhoneNumber),
		Email:              utility.RemoveZeroWidth(req.Data.Email),
		Status:             req.Data.Status,
		LastDassTime:       req.Data.LastDassTime,
		LastIesrTime:       req.Data.LastIesrTime,
		LastDassResult:     req.Data.LastDassResult,
		LastIesrResult:     req.Data.LastIesrResult,
		Remarks:            req.Data.Remarks,
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
	if patient.ID == "" {
		return nil, constants.InvalidPatientIDError
	}
	if patient.Name == "" {
		return nil, constants.InvalidPatientNameError
	}
	return patient, nil
}
