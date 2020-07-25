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

type DeletePatientHandler struct {
	Model model.IModel
}

func (s *DeletePatientHandler) DeletePatient(ctx context.Context, req *pb.CommonDeleteRequest, user *dto.User) (*pb.CommonPatientResponse, error) {
	s.processReq(req)

	patient, err := s.Model.DeletePatient(ctx, req.Id, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.PatientToResponse(patient)
	return resp, nil
}

func (s *DeletePatientHandler) processReq(req *pb.CommonDeleteRequest) {
	req.Id = utility.NormalizeID(req.Id)
}
