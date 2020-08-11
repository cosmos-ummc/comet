package game

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetGameHandler struct {
	Model model.IModel
}

func (s *GetGameHandler) GetGame(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonGameResponse, error) {
	game, err := s.Model.GetGame(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.GameNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.GameToResponse(game)
	return resp, nil
}
