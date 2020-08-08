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

type UpdateChatMessageHandler struct {
	Model model.IModel
}

func (s *UpdateChatMessageHandler) UpdateChatMessage(ctx context.Context, req *pb.CommonChatMessageRequest) (*pb.CommonChatMessageResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	chatMessage := s.reqToChatMessage(req)

	v, err := s.Model.UpdateChatMessage(ctx, chatMessage)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatMessageNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatMessageToResponse(v)
	return resp, nil
}

func (s *UpdateChatMessageHandler) reqToChatMessage(req *pb.CommonChatMessageRequest) *dto.ChatMessage {
	chatMessage := &dto.ChatMessage{
		ID:       utility.RemoveZeroWidth(req.Id),
		RoomID:    req.Data.RoomId,
		SenderID:  req.Data.SenderId,
		Content:   req.Data.Content,
		Timestamp: utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
	}
	return chatMessage
}
