package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type ClientMessageEventHandler struct {
	Model model.IModel
}

func (s *ClientMessageEventHandler) ClientMessageEvent(ctx context.Context, req *pb.ClientMessageEventRequest) (*pb.ClientMessageEventResponse, error) {
	// get patient
	p, err := s.Model.GetPatient(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if p.TelegramID == "" {
		return &pb.ClientMessageEventResponse{
			Ok: false,
		}, nil
	}

	if req.Daily {
		err = utility.SendBotNotification(p.TelegramID, constants.ReminderMessageDaily)
	} else {
		err = utility.SendBotNotification(p.TelegramID, constants.ReminderMessage)
	}
	
	if err != nil {
		return nil, err
	}
	return &pb.ClientMessageEventResponse{
		Ok: true,
	}, nil
}
