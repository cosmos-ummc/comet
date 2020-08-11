package patient

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

type UpdatePatientsHandler struct {
	Model model.IModel
}

func (s *UpdatePatientsHandler) UpdatePatients(ctx context.Context, req *pb.CommonPatientsRequest) (*pb.CommonIdsResponse, error) {
	// TODO: Support Batch Updates
	if req.Data == nil || len(req.Ids) != 1 {
		return nil, constants.InvalidArgumentError
	}

	patient, err := s.validateAndProcessReq(req)
	if err != nil {
		return nil, err
	}

	ids, err := s.Model.UpdatePatients(ctx, patient, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}

	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdatePatientsHandler) validateAndProcessReq(req *pb.CommonPatientsRequest) (*dto.Patient, error) {
	if utf8.RuneCountInString(req.Data.Remarks) > 200 {
		return nil, constants.RemarksTooLongError
	}

	patient := &dto.Patient{
		ID:                 req.Ids[0],
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
