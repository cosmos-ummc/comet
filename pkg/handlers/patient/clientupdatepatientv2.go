package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientUpdatePatientV2Handler struct {
	Model model.IModel
}

func (s *ClientUpdatePatientV2Handler) ClientUpdatePatientV2(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*pb.ClientUpdatePatientV2Response, error) {
	patient := s.processReq(req)

	_, err := s.Model.ClientUpdatePatient(ctx, patient)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}

	// if telegram ID not found, return incomplete
	if patient.TelegramID == "" {
		return &pb.ClientUpdatePatientV2Response{
			RegistrationStatus: constants.Incomplete,
		}, nil
	}

	return &pb.ClientUpdatePatientV2Response{
		RegistrationStatus: constants.Complete,
	}, nil
}

func (s *ClientUpdatePatientV2Handler) processReq(req *pb.ClientUpdatePatientRequest) *dto.Patient {
	patient := &dto.Patient{
		ID:            utility.RemoveZeroWidth(req.Id),
		TelegramID:    utility.RemoveZeroWidth(req.TelegramId),
		Consent:       req.Consent,
		PrivacyPolicy: req.PrivacyPolicy,
	}
	return patient
}
