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

type DeleteGameHandler struct {
	Model model.IModel
}

func (s *DeleteGameHandler) DeleteGame(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonGameResponse, error) {
	rslt, err := s.Model.DeleteGame(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.GameNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.GameToResponse(rslt)
	return resp, nil
}
