package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"time"
)

type ClientGetPatientsHandler struct {
	Model model.IModel
}

func (s *ClientGetPatientsHandler) ClientGetPatients(ctx context.Context, req *pb.ClientGetPatientsRequest) (*pb.ClientGetPatientsResponse, error) {
	s.processReq(req)

	// for monitoring date, return the list of patients who signed consent on the date
	if req.Day != 0 {

		// get 12 am's time from day n
		now := utility.MalaysiaTime(time.Now())
		daySelected, err := utility.DateStringToTime(utility.TimeToDateString(now.Add(time.Duration(-req.Day+1) * 24 * time.Hour)))
		if err != nil {
			logger.Log.Error("ClientGetPatients: " + err.Error())
		}
		daySelectedEnd, err := utility.DateStringToTime(utility.TimeToDateString(now.Add(time.Duration(-req.Day+2) * 24 * time.Hour)))
		if err != nil {
			logger.Log.Error("ClientGetPatients: " + err.Error())
		}

		patients, err := s.Model.ClientGetPatientsByConsentTime(ctx, utility.TimeToMilli(daySelected), utility.TimeToMilli(daySelectedEnd))
		if err != nil {
			logger.Log.Error("ClientGetPatients: " + err.Error())
			return nil, constants.InternalError
		}
		return s.patientsToResp(patients), nil
	}

	// if no param is given, return all patients
	if req.TelegramId == "" && req.PhoneNumber == "" && req.Id == "" {
		_, patients, err := s.Model.QueryPatients(ctx, nil, nil, nil)
		if err != nil {
			return nil, constants.InternalError
		}
		return s.patientsToResp(patients), nil
	}

	filter := map[string]interface{}{}
	if req.Id != "" {
		filter[constants.ID] = req.Id
	} else if req.TelegramId != "" {
		filter[constants.TelegramID] = req.TelegramId
	} else {
		filter[constants.PhoneNumber] = req.PhoneNumber
	}

	_, patients, err := s.Model.QueryPatients(ctx, nil, nil, filter)
	if err != nil {
		return nil, constants.InternalError
	}

	return s.patientsToResp(patients), nil
}

func (s *ClientGetPatientsHandler) patientsToResp(patients []*dto.Patient) *pb.ClientGetPatientsResponse {
	var resps []*pb.Patient
	for _, patient := range patients {
		resp := &pb.Patient{
			Id:                 patient.ID,
			TelegramId:         patient.TelegramID,
			Name:               patient.Name,
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
			UserId:             patient.UserID,
			DaySinceMonitoring: patient.DaySinceMonitoring,
			HasCompleted:       patient.HasCompleted,
			MentalStatus:       patient.MentalStatus,
			Type:               patient.Type,
			SwabDate:           patient.SwabDate,
			SwabResult:         patient.SwabResult,
			StressStatus:       patient.StressStatus,
			PtsdStatus:         patient.PtsdStatus,
			DepressionStatus:   patient.DepressionStatus,
			AnxietyStatus:      patient.AnxietyStatus,
		}
		resps = append(resps, resp)
	}
	rslt := &pb.ClientGetPatientsResponse{
		Patients: resps,
	}
	return rslt
}

func (s *ClientGetPatientsHandler) processReq(req *pb.ClientGetPatientsRequest) {
	req.PhoneNumber = utility.NormalizePhoneNumber(req.PhoneNumber, "")
	req.Id = utility.NormalizeID(req.Id)
}
