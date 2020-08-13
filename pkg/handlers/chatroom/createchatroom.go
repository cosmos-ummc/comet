package chatRoom

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"fmt"
	"github.com/twinj/uuid"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

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
		Timestamp:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
	}

	// generate opponent name
	r, err := http.Get("https://api.namefake.com/")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	bodyString = bodyString[9:]
	bodyString = bodyString[0:strings.Index(bodyString, "\"")]
	chatRoom.Name = "Anonymous " + bodyString

	rslt, err := s.Model.CreateChatRoom(ctx, chatRoom)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.ChatRoomAlreadyExistError
		}
		return nil, constants.InternalError
	}

	// trigger event to pop up explore
	_, err = http.Get(fmt.Sprintf("https://chat.quaranteams.tk/message?id=%s&id2=%s", chatRoom.ParticipantIDs[0], chatRoom.ParticipantIDs[1]))
	if err != nil {
		return nil, err
	}

	resp := utility.ChatRoomToResponse(rslt)
	return resp, nil
}
