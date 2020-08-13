package user

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"context"
)

type ClientCheckCompletedHandler struct {
	Model model.IModel
}

func (s *ClientCheckCompletedHandler) ClientCheckCompleted(ctx context.Context, req *pb.ClientCheckCompletedRequest) (*pb.ClientCheckCompletedResponse, error) {
	_, patients, err := s.Model.QueryPatients(ctx, nil, nil, map[string]interface{}{
		constants.UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, constants.InvalidArgumentError
	}
	patient := patients[0]

	return &pb.ClientCheckCompletedResponse{
		Completed: patient.HasCompleted,
	}, nil
}
