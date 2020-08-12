package tip

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetTipsHandler struct {
	Model model.IModel
}

func (s *ClientGetTipsHandler) GetTips(ctx context.Context, req *pb.ClientGetTipsRequest) (*pb.CommonTipsResponse, error) {
	// do query
	_, tips, err := s.Model.QueryTips(ctx, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	if len(tips) == 0 {
		return &pb.CommonTipsResponse{}, nil
	}

	// shuffle tips
	utility.ShuffleTips(tips)
	tips = tips[0:1]

	resp := utility.TipsToResponse(tips)
	resp.Total = 1
	return resp, nil
}
