package chatMessage

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateChatMessagesHandler struct {
	Model model.IModel
}

func (s *UpdateChatMessagesHandler) UpdateChatMessages(ctx context.Context, req *pb.CommonChatMessagesRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	chatMessage := s.reqToChatMessage(req)

	ids, err := s.Model.UpdateChatMessages(ctx, chatMessage, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatMessageNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateChatMessagesHandler) reqToChatMessage(req *pb.CommonChatMessagesRequest) *dto.ChatMessage {
	chatMessage := &dto.ChatMessage{
		ID:        req.Ids[0],
		RoomID:    req.Data.RoomId,
		SenderID:  req.Data.SenderId,
		Content:   req.Data.Content,
		Timestamp: utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
	}
	return chatMessage
}
