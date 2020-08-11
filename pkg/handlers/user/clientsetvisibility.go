package user

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"context"
)

type ClientSetVisibilityHandler struct {
	Model model.IModel
}

func (s *ClientSetVisibilityHandler) ClientSetVisibility(ctx context.Context, req *pb.ClientSetVisibilityRequest) (*pb.ClientSetVisibilityResponse, error) {
	u, err := s.Model.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	u.Visible = req.Visible

	_, err = s.Model.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	return &pb.ClientSetVisibilityResponse{
		Ok: true,
	}, nil
}
