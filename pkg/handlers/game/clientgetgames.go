package game

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetGamesHandler struct {
	Model model.IModel
}

func (s *ClientGetGamesHandler) GetGames(ctx context.Context, req *pb.ClientGetGamesRequest) (*pb.CommonGamesResponse, error) {
	// do query
	_, games, err := s.Model.QueryGames(ctx, nil, nil, nil)
	if err != nil {
		return nil, err
	}

	// shuffle games and select 3
	utility.ShuffleGames(games)
	games = games[0:3]

	resp := utility.GamesToResponse(games)
	resp.Total = int64(len(games))
	return resp, nil
}
