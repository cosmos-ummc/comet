package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientUpdatePatientHandler struct {
	Model model.IModel
}

func (s *ClientUpdatePatientHandler) ClientUpdatePatient(ctx context.Context, req *pb.ClientUpdatePatientRequest) (*empty.Empty, error) {
	patient := s.processReq(req)

	_, err := s.Model.ClientUpdatePatient(ctx, patient)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}
	return &empty.Empty{}, nil
}

func (s *ClientUpdatePatientHandler) processReq(req *pb.ClientUpdatePatientRequest) *dto.Patient {
	patient := &dto.Patient{
		ID:            utility.RemoveZeroWidth(req.Id),
		TelegramID:    utility.RemoveZeroWidth(req.TelegramId),
		Consent:       req.Consent,
		PrivacyPolicy: req.PrivacyPolicy,
	}

	return patient
}
