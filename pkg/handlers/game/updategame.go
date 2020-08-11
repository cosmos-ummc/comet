package game

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

type UpdateGameHandler struct {
	Model model.IModel
}

func (s *UpdateGameHandler) UpdateGame(ctx context.Context, req *pb.CommonGameRequest) (*pb.CommonGameResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	game := s.reqToGame(req)

	v, err := s.Model.UpdateGame(ctx, game)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.GameNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.GameToResponse(v)
	return resp, nil
}

func (s *UpdateGameHandler) reqToGame(req *pb.CommonGameRequest) *dto.Game {
	game := &dto.Game{
		ID:      utility.RemoveZeroWidth(req.Id),
		Link:    req.Data.Link,
		ImgPath: req.Data.ImgPath,
		Type:    req.Data.Type,
	}
	return game
}
