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
			Id:                patient.ID,
			TelegramId:        patient.TelegramID,
			Name:              patient.Name,
			Status:            patient.Status,
			PhoneNumber:       patient.PhoneNumber,
			Email:             patient.Email,
			LastDeclared:      patient.LastDeclared,
			SwabCount:         patient.SwabCount,
			Episode:           patient.Episode,
			Type:              patient.Type,
			TypeChangeDate:    patient.TypeChangeDate,
			LastDeclareResult: patient.LastDeclareResult,
			ExposureDate:      patient.ExposureDate,
			ExposureSource:    patient.ExposureSource,
			DaysSinceExposure: patient.DaysSinceExposure,
			RegistrationNum:   patient.RegistrationNum,
			AlternateContact:  patient.AlternateContact,
			IsolationAddress:  patient.IsolationAddress,
			SymptomDate:       patient.SymptomDate,
			SwabDate:          patient.SwabDate,
			Remarks:           patient.Remarks,
			Localization:      patient.Localization,
			Consent:           patient.Consent,
			PrivacyPolicy:     patient.PrivacyPolicy,
			FeverStartDate:    patient.FeverStartDate,
			FeverContDay:      patient.FeverContDay,
			DaysSinceSwab:     patient.DaysSinceSwab,
			HomeAddress:       patient.HomeAddress,
			IsSameAddress:     patient.IsSameAddress,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.ClientGetUndeclaredPatientsResponse{
		Patients: resps,
	}
	return rslt
}
