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
		ok = true
		user = targetUser
		break
	}

	if !ok {
		return &pb.ClientNewMatchResponse{Msg: string(body)}, nil
	}

	// match! Create ChatRoom for both people
	c, err := s.Model.CreateChatRoom(ctx, &dto.ChatRoom{
		ID:             uuid.NewV4().String(),
		ParticipantIDs: []string{u.ID, user.ID},
		Timestamp:      utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
	})
	if err != nil {
		return nil, err
	}

	return &pb.ClientNewMatchResponse{
		User:     utility.UserToPb(user),
		ChatRoom: utility.ChatRoomToPb(c),
		Msg:      string(body),
	}, nil
}
