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

type GetCallPatientsHandler struct {
	Model model.IModel
}

func (s *GetCallPatientsHandler) GetCallPatients(ctx context.Context, req *pb.CommonGetsRequest, user *dto.User) (*pb.CommonPatientsResponse, error) {
	t, err := utility.DateStringToTime(utility.TimeToDateString(utility.MalaysiaTime(time.Now())))
	if err != nil {
		return nil, constants.InternalError
	}
	timeMilli := utility.TimeToMilli(t)

	_, declarations, err := s.Model.QueryDeclarationsByCallingStatusAndTime(ctx, []int64{constants.UMMCCalled, constants.PatientCalled}, timeMilli, constants.UserPatientMap[user.Role])
	if err != nil {
		return nil, constants.InternalError
	}

	// get called patients
	var patients []*dto.Patient
	for _, declaration := range declarations {
		// get patient
		p, err := s.Model.GetPatient(ctx, declaration.PatientID, constants.UserPatientMap[user.Role])
		if err != nil {
			// if type not allowed, skip the record
			continue
		}

		// skip patients that are other status
		if p.Status != constants.Symptomatic && p.Status != constants.Asymptomatic && p.Status != constants.ConfirmedButNotAdmitted {
			continue
		}

		// put report data in
		p.CallingStatus = declaration.CallingStatus

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

func (s *GetCallPatientsHandler) sortPatients(patients []*dto.Patient, field string, order string) {

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
	case "lastDeclared":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].LastDeclared < patients[j].LastDeclared
		})
	case "swabCount":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].SwabCount < patients[j].SwabCount
		})
	case "episode":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Episode < patients[j].Episode
		})
	case "type":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Type < patients[j].Type
		})
	case "lastDeclareResult":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].LastDeclareResult
		})
	case "exposureDate":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].ExposureDate < patients[j].ExposureDate
		})
	case "exposureSource":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].ExposureSource < patients[j].ExposureSource
		})
	case "daysSinceExposure":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].DaysSinceExposure < patients[j].DaysSinceExposure
		})
	case "registrationNum":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].RegistrationNum < patients[j].RegistrationNum
		})
	case "alternateContact":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].AlternateContact < patients[j].AlternateContact
		})
	case "isolationAddress":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].IsolationAddress < patients[j].IsolationAddress
		})
	case "symptomDate":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].SymptomDate < patients[j].SymptomDate
		})
	case "swabDate":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].SwabDate < patients[j].SwabDate
		})
	case "callingStatus":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].CallingStatus < patients[j].CallingStatus
		})
	case "consent":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].Consent < patients[j].Consent
		})
	case "privacyPolicy":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].PrivacyPolicy < patients[j].PrivacyPolicy
		})
	case "feverStartDate":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].FeverStartDate < patients[j].FeverStartDate
		})
	case "feverContDay":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].FeverContDay < patients[j].FeverContDay
		})
	case "daysSinceSwab":
		sort.Slice(patients, func(i, j int) bool {
			return patients[i].DaysSinceSwab < patients[j].DaysSinceSwab
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
