package game

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateGamesHandler struct {
	Model model.IModel
}

func (s *UpdateGamesHandler) UpdateGames(ctx context.Context, req *pb.CommonGamesRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	game := s.reqToGame(req)

	ids, err := s.Model.UpdateGames(ctx, game, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.GameNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateGamesHandler) reqToGame(req *pb.CommonGamesRequest) *dto.Game {
	game := &dto.Game{
		LinkAdr:    req.Data.LinkAdr,
		LinkIos:    req.Data.LinkIos,
		ImgPathAdr: req.Data.ImgPathAdr,
		ImgPathIos: req.Data.ImgPathIos,
	}
	return game
}
