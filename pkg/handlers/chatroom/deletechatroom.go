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

type DeleteChatRoomHandler struct {
	Model model.IModel
}

func (s *DeleteChatRoomHandler) DeleteChatRoom(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonChatRoomResponse, error) {
	rslt, err := s.Model.DeleteChatRoom(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatRoomNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatRoomToResponse(rslt)
	return resp, nil
}
