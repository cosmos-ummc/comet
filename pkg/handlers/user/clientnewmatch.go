package user

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
)

type ClientNewMatchHandler struct {
	Model model.IModel
}

func (s *ClientNewMatchHandler) ClientNewMatch(ctx context.Context, req *pb.ClientNewMatchRequest) (*pb.ClientNewMatchResponse, error) {
	// get patient
	_, ps, err := s.Model.QueryPatients(ctx, nil, nil, map[string]interface{}{
		constants.UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(ps) == 0 {
		return nil, constants.InvalidArgumentError
	}

	p := ps[0]

	if p.HasCompleted {
		return &pb.ClientNewMatchResponse{}, nil
	}

	// get user
	u, err := s.Model.GetUser(ctx, p.UserID)
	if err != nil {
		return nil, err
	}

	// find patient IDs via similarity API
	resp, err := http.Get(fmt.Sprintf("https://chat.quaranteams.tk/similarusers?id=%s", p.ID))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(body)
	ids := strings.Split(bodyString[1:len(bodyString)-1], ",")
	i := 0
	for _, rs := range ids {
		ids[i] = rs[1 : len(rs)-1]
		i += 1
	}

	ok := false
	user := &dto.User{}

	// for each patient ID, figure out who are available
	for _, patientID := range ids {
		targetPatient, err := s.Model.GetPatient(ctx, patientID)
		if err != nil {
			continue
		}
		targetUser, err := s.Model.GetUser(ctx, targetPatient.UserID)
		if err != nil {
			continue
		}
		if !targetUser.Visible ||
			utility.StringInSlice(targetUser.ID, u.BlockList) ||
			utility.StringInSlice(u.ID, targetUser.BlockList) ||
			targetPatient.HasCompleted {
			continue
		}
		// check if chatroom already exist
		chatrooms, err := s.Model.QueryByUsers(ctx, []string{u.ID, targetUser.ID})
		if err != nil {
			continue
		}
		if len(chatrooms) != 0 {
			continue
		}

		ok = true
		user = targetUser
		break
	}

	if !ok {
		return &pb.ClientNewMatchResponse{Msg: string(body)}, nil
	}

	// generate opponent name
	r, err := http.Get("https://api.namefake.com/")
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	body2, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	bodyString2 := string(body2)
	bodyString2 = bodyString2[9:]
	bodyString2 = bodyString2[0:strings.Index(bodyString2, "\"")]

	// match! Create ChatRoom for both people
	c, err := s.Model.CreateChatRoom(ctx, &dto.ChatRoom{
		ID:             uuid.NewV4().String(),
		ParticipantIDs: []string{u.ID, user.ID},
		Timestamp:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
		Name:           "Anonymous " + bodyString2,
	})
	if err != nil {
		return nil, err
	}

	// trigger event to pop up explore
	_, err = http.Get(fmt.Sprintf("https://chat.quaranteams.tk/chatroom?id=%s&id2=%s", c.ParticipantIDs[0], c.ParticipantIDs[1]))
	if err != nil {
		return nil, err
	}

	return &pb.ClientNewMatchResponse{
		User:     utility.UserToPb(user),
		ChatRoom: utility.ChatRoomToPb(c),
		Msg:      string(body),
	}, nil
}
