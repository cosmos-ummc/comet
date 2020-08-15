package game

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteGamesHandler struct {
	Model model.IModel
}

func (s *DeleteGamesHandler) DeleteGames(ctx context.Context, req *pb.CommonDeletesRequest) (*pb.CommonIdsResponse, error) {
	req.Ids = s.processReq(req.Ids)
	ids, err := s.Model.DeleteGames(ctx, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.GameNotFoundError
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *DeleteGamesHandler) processReq(ids []string) []string {
	res := []string{}
	for _, id := range ids {
		split := strings.Split(id, ",")
		res = append(res, split...)
	}
	return res
}
