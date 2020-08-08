package chatRoom

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetChatRoomHandler struct {
	Model model.IModel
}

func (s *GetChatRoomHandler) GetChatRoom(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonChatRoomResponse, error) {
	chatRoom, err := s.Model.GetChatRoom(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatRoomNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatRoomToResponse(chatRoom)
	return resp, nil
}
