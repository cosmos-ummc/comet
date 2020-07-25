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

type GetPatientHandler struct {
	Model model.IModel
}

func (s *GetPatientHandler) GetPatient(ctx context.Context, req *pb.CommonGetRequest, user *dto.User) (*pb.CommonPatientResponse, error) {
	s.processReq(req)

	patient, err := s.Model.GetPatient(ctx, req.Id, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.PatientToResponse(patient)
	return resp, nil
}

func (s *GetPatientHandler) processReq(req *pb.CommonGetRequest) {
	req.Id = utility.NormalizeID(req.Id)
}
