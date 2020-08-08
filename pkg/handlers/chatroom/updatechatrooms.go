package chatRoom

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateChatRoomsHandler struct {
	Model model.IModel
}

func (s *UpdateChatRoomsHandler) UpdateChatRooms(ctx context.Context, req *pb.CommonChatRoomsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	chatRoom := s.reqToChatRoom(req)

	ids, err := s.Model.UpdateChatRooms(ctx, chatRoom, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatRoomNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateChatRoomsHandler) reqToChatRoom(req *pb.CommonChatRoomsRequest) *dto.ChatRoom {
	chatRoom := &dto.ChatRoom{
		ID:             req.Ids[0],
		ParticipantIDs: req.Data.ParticipantIds,
		Blocked:        req.Data.Blocked,
	}
	return chatRoom
}
