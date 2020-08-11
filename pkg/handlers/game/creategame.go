package game

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateGameHandler struct {
	Model model.IModel
}

func (s *CreateGameHandler) CreateGame(ctx context.Context, req *pb.CommonGameRequest) (*pb.CommonGameResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	game := &dto.Game{
		ID:         uuid.NewV4().String(),
		LinkAdr:    req.Data.LinkAdr,
		LinkIos:    req.Data.LinkIos,
		ImgPathAdr: req.Data.ImgPathAdr,
		ImgPathIos: req.Data.ImgPathIos,
	}

	rslt, err := s.Model.CreateGame(ctx, game)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.GameAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.GameToResponse(rslt)
	return resp, nil
}
