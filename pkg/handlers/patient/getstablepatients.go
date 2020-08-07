package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"sort"
	"time"
)

type GetStablePatientsHandler struct {
	Model model.IModel
}

func (s *GetStablePatientsHandler) GetStablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	t, err := utility.DateStringToTime(utility.TimeToDateString(utility.MalaysiaTime(time.Now())))
	if err != nil {
		return nil, constants.InternalError
	}
	timeMilli := utility.TimeToMilli(t)

	_, declarations, err := s.Model.GetStableDeclarations(ctx, timeMilli)
	if err != nil {
		return nil, constants.InternalError
	}

	// get no yet called patients
	patients := []*dto.Patient{}
	for _, declaration := range declarations {
		// get patient
		p, err := s.Model.GetPatient(ctx, declaration.PatientID)
		if err != nil {
			// if type not allowed, skip the record
			continue
		}

		patients = append(patients, p)
	}

	// sort if needed
	if req.Item != "" && req.Order != "" {
		s.sortPatients(patients, req.Item, req.Order)
	}

	// get total
	total := int64(len(patients))

	index := int64(0)
	var rslt []*dto.Patient

	for _, patient := range patients {
		if index < req.From {
			index += 1
			continue
		}
		if req.To != 0 && index > req.To {
			index += 1
			continue
		}
		index += 1
		rslt = append(rslt, patient)
	}

	resp := utility.PatientsToResponse(rslt)
	resp.Total = total
	return resp, nil
}

func (s *GetStablePatientsHandler) sortPatients(patients []*dto.Patient, field string, order string) {

	switch field {
	case "id":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].ID < patients[j].ID
		})
	case "name":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Name < patients[j].Name
		})
	case "telegramId":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].TelegramID < patients[j].TelegramID
		})
	case "phoneNumber":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].PhoneNumber < patients[j].PhoneNumber
		})
	case "email":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Email < patients[j].Email
		})
	case "status":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Status < patients[j].Status
		})
	case "isolationAddress":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].IsolationAddress < patients[j].IsolationAddress
		})
	case "consent":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Consent < patients[j].Consent
		})
	case "privacyPolicy":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].PrivacyPolicy < patients[j].PrivacyPolicy
		})
	case "homeAddress":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].HomeAddress < patients[j].HomeAddress
		})
	default:
	}

	if order == "DESC" {
		// reverse slice
		for i, j := 0, len(patients)-1; i < j; i, j = i+1, j-1 {
			patients[i], patients[j] = patients[j], patients[i]
		}
	}
}
