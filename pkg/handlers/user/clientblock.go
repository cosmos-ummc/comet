package user

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"context"
	"net/http"
)

type ClientBlockHandler struct {
	Model model.IModel
}

func (s *ClientBlockHandler) ClientBlock(ctx context.Context, req *pb.ClientBlockRequest) (*pb.ClientBlockResponse, error) {
	// get user
	u, err := s.Model.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// get target user
	targetU, err := s.Model.GetUser(ctx, req.TargetId)
	if err != nil {
		return nil, err
	}

	// update user block list
	u.BlockList = append(u.BlockList, targetU.ID)
	_, err = s.Model.UpdateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	// update target user block list
	targetU.BlockList = append(targetU.BlockList, u.ID)
	_, err = s.Model.UpdateUser(ctx, targetU)
	if err != nil {
		return nil, err
	}

	// query chatrooms
	chatrooms, err := s.Model.QueryByUsers(ctx, []string{u.ID, targetU.ID})
	if err != nil {
		return nil, err
	}
	for _, room := range chatrooms {
		// update each room to blocked
		room.Blocked = true
		_, err = s.Model.UpdateChatRoom(ctx, room)
		if err != nil {
			return nil, err
		}
	}

	// trigger event to refresh chatrooms
	_, err = http.Get("https://chat.quaranteams.tk/block")
	if err != nil {
		return nil, err
	}

	return &pb.ClientBlockResponse{
		Ok: true,
	}, nil
}
