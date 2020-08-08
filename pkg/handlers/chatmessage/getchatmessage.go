package chatMessage

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetChatMessageHandler struct {
	Model model.IModel
}

func (s *GetChatMessageHandler) GetChatMessage(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonChatMessageResponse, error) {
	chatMessage, err := s.Model.GetChatMessage(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatMessageNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatMessageToResponse(chatMessage)
	return resp, nil
}
