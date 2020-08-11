package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ClientVerifyPatientCompleteHandler struct {
	Model model.IModel
}

func (s *ClientVerifyPatientCompleteHandler) ClientVerifyPatientComplete(ctx context.Context, req *pb.ClientVerifyPatientCompleteRequest) (*pb.ClientVerifyPatientCompleteResponse, error) {
	hasCompleted, err := s.Model.VerifyPatientComplete(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.PatientNotFoundError
		}
		return nil, constants.InternalError
	}
	return &pb.ClientVerifyPatientCompleteResponse{
		HasCompleted: hasCompleted,
	}, nil
}
