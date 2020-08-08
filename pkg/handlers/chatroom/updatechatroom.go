package chatRoom

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

type UpdateChatRoomHandler struct {
	Model model.IModel
}

func (s *UpdateChatRoomHandler) UpdateChatRoom(ctx context.Context, req *pb.CommonChatRoomRequest) (*pb.CommonChatRoomResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	chatRoom := s.reqToChatRoom(req)

	v, err := s.Model.UpdateChatRoom(ctx, chatRoom)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatRoomNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatRoomToResponse(v)
	return resp, nil
}

func (s *UpdateChatRoomHandler) reqToChatRoom(req *pb.CommonChatRoomRequest) *dto.ChatRoom {
	chatRoom := &dto.ChatRoom{
		ID:             utility.RemoveZeroWidth(req.Id),
		ParticipantIDs: req.Data.ParticipantIds,
		Blocked:        req.Data.Blocked,
	}
	return chatRoom
}
