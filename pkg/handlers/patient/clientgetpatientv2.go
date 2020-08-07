package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetPatientV2Handler struct {
	Model model.IModel
}

func (s *ClientGetPatientV2Handler) ClientGetPatientV2(ctx context.Context, req *pb.ClientGetPatientV2Request) (*pb.ClientGetPatientV2Response, error) {
	s.processReq(req)

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

func (s *ClientGetPatientV2Handler) patientsToResp(patients []*dto.Patient) *pb.ClientGetPatientV2Response {

	// if no patient found, return 0
	if len(patients) < 1 {
		return &pb.ClientGetPatientV2Response{
			Data: &pb.Patient{RegistrationStatus: constants.NotFound},
		}
	}

	patient := patients[0]

	// if telegram id not found, return 2
	if patient.TelegramID == "" {
		patient.RegistrationStatus = constants.Incomplete
		return &pb.ClientGetPatientV2Response{
			Data: utility.PatientToPb(patient),
		}
	}

	patient.RegistrationStatus = constants.Complete
	return &pb.ClientGetPatientV2Response{
		Data: utility.PatientToPb(patient),
	}
}

func (s *ClientGetPatientV2Handler) processReq(req *pb.ClientGetPatientV2Request) {
	req.PhoneNumber = utility.NormalizePhoneNumber(req.PhoneNumber, "")
	req.Id = utility.NormalizeID(req.Id)
}
