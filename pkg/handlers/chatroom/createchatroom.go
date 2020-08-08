package chatRoom

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateChatRoomHandler struct {
	Model model.IModel
}

func (s *CreateChatRoomHandler) CreateChatRoom(ctx context.Context, req *pb.CommonChatRoomRequest) (*pb.CommonChatRoomResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	chatRoom := &dto.ChatRoom{
		ID:             uuid.NewV4().String(),
		ParticipantIDs: req.Data.ParticipantIds,
		Blocked:        req.Data.Blocked,
	}

	rslt, err := s.Model.CreateChatRoom(ctx, chatRoom)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.ChatRoomAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.ChatRoomToResponse(rslt)
	return resp, nil
}
