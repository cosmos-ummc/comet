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

type DeleteChatMessageHandler struct {
	Model model.IModel
}

func (s *DeleteChatMessageHandler) DeleteChatMessage(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonChatMessageResponse, error) {
	rslt, err := s.Model.DeleteChatMessage(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatMessageNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatMessageToResponse(rslt)
	return resp, nil
}
