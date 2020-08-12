package chatRoom

import (
	pb "comet/pkg/api"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientGetChatRoomsHandler struct {
	Model model.IModel
}

func (s *ClientGetChatRoomsHandler) ClientGetChatRooms(ctx context.Context, req *pb.ClientGetChatRoomsRequest) (*pb.CommonChatRoomsResponse, error) {

	rooms, err := s.Model.QueryByUsers(ctx, []string{req.Id})
	if err != nil {
		return nil, err
	}

	resp := utility.ChatRoomsToResponse(rooms)
	resp.Total = int64(len(rooms))
	return resp, nil
}
