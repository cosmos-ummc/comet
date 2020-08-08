package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"context"
)

type GetStablePatientsHandler struct {
	Model model.IModel
}

func (s *GetStablePatientsHandler) GetStablePatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	// TODO: Get patients more than given scores can already
	return &pb.CommonPatientsResponse{}, nil
}
