package chatMessage

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"fmt"
	"github.com/twinj/uuid"
	"net/http"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateChatMessageHandler struct {
	Model model.IModel
}

func (s *CreateChatMessageHandler) CreateChatMessage(ctx context.Context, req *pb.CommonChatMessageRequest) (*pb.CommonChatMessageResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	chatMessage := &dto.ChatMessage{
		ID:        uuid.NewV4().String(),
		RoomID:    req.Data.RoomId,
		SenderID:  req.Data.SenderId,
		Content:   req.Data.Content,
		Timestamp: utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
	}

	rslt, err := s.Model.CreateChatMessage(ctx, chatMessage)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.ChatMessageAlreadyExistError
		}
		return nil, constants.InternalError
	}

	// trigger general socket event for hot reload
	_, err = http.Get(fmt.Sprintf("https://chat.quaranteams.tk/message?roomid=%s", chatMessage.RoomID))
	if err != nil {
		return nil, err
	}

	resp := utility.ChatMessageToResponse(rslt)
	return resp, nil
}
