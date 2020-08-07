package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"time"
)

type ClientGetUndeclaredPatientsHandler struct {
	Model model.IModel
}

func (s *ClientGetUndeclaredPatientsHandler) ClientGetUndeclaredPatients(ctx context.Context, req *pb.ClientGetUndeclaredPatientsRequest) (*pb.ClientGetUndeclaredPatientsResponse, error) {
	timeString := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	t, err := utility.DateStringToTime(timeString)
	if err != nil {
		return nil, err
	}

	patients, err := s.Model.ClientGetUndeclaredPatientsByTime(ctx, utility.TimeToMilli(t))
	if err != nil {
		return nil, constants.InternalError
	}

	return s.patientsToResp(patients), nil
}

func (s *ClientGetUndeclaredPatientsHandler) patientsToResp(patients []*dto.Patient) *pb.ClientGetUndeclaredPatientsResponse {
	var resps []*pb.Patient
	for _, patient := range patients {
		resp := &pb.Patient{
			Id:                 patient.ID,
			TelegramId:         patient.TelegramID,
			Name:               patient.Name,
			Status:             patient.Status,
			PhoneNumber:        patient.PhoneNumber,
			Email:              patient.Email,
			IsolationAddress:   patient.IsolationAddress,
			Remarks:            patient.Remarks,
			Consent:            patient.Consent,
			PrivacyPolicy:      patient.PrivacyPolicy,
			HomeAddress:        patient.HomeAddress,
			LastDassTime:       patient.LastDassTime,
			LastIesrTime:       patient.LastIesrTime,
			LastDassResult:     patient.LastDassResult,
			LastIesrResult:     patient.LastIesrResult,
			RegistrationStatus: patient.RegistrationStatus,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.ClientGetUndeclaredPatientsResponse{
		Patients: resps,
	}
	return rslt
}
