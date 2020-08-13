package user

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"context"
)

type ClientSetNotFirstTimeHandler struct {
	Model model.IModel
}

func (s *ClientSetNotFirstTimeHandler) ClientSetNotFirstTime(ctx context.Context, req *pb.ClientSetNotFirstTimeRequest) (*pb.ClientSetNotFirstTimeResponse, error) {
	u, err := s.Model.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	u.NotFirstTimeChat = true

	_, err = s.Model.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &pb.ClientSetNotFirstTimeResponse{
		Ok: true,
	}, nil
}
