package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
)

type ClientTutorialHandler struct {
	Model model.IModel
}

func (s *ClientTutorialHandler) ClientTutorial(ctx context.Context, req *pb.TutorialRequest) (*empty.Empty, error) {
	p, err := s.Model.GetPatient(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if p.TutorialStage >= 4 {
		return &empty.Empty{}, nil
	}

	if p.TutorialStage == 0 {
		err = utility.SendBotNotification(p.TelegramID, constants.PhoneLinkMessage)
	}
	if p.TutorialStage == 1 {
		err = utility.SendBotNotification(p.TelegramID, constants.FirstDassMessage)
	}
	if p.TutorialStage == 2 {
		err = utility.SendBotNotification(p.TelegramID, constants.FirstDailyMessage)
	}
	if p.TutorialStage == 3 {
		err = utility.SendBotNotification(p.TelegramID, constants.FirstIesrMessage)
	}
	p.TutorialStage += 1

	return &empty.Empty{}, nil
}
