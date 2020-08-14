package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"time"
)

type ClientTutorialHandler struct {
	Model model.IModel
}

func delaySecond(telegramID string, tutorialStage int64, n time.Duration) {
	for _ = range time.Tick(n * time.Second) {
		if tutorialStage == 0 {
			_ = utility.SendBotNotification(telegramID, constants.PhoneLinkMessage)
		}
		if tutorialStage == 1 {
			_ = utility.SendBotNotification(telegramID, constants.FirstDassMessage)
		}
		if tutorialStage == 2 {
			_ = utility.SendBotNotification(telegramID, constants.FirstDailyMessage)
		}
		if tutorialStage == 3 {
			_ = utility.SendBotNotification(telegramID, constants.FirstIesrMessage)
		}
		break
	}
}

func (s *ClientTutorialHandler) ClientTutorial(ctx context.Context, req *pb.TutorialRequest) (*empty.Empty, error) {
	p, err := s.Model.GetPatient(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if p.TutorialStage >= 4 {
		return &empty.Empty{}, nil
	}

	go delaySecond(p.TelegramID, p.TutorialStage, 3) // delay first seconds

	p.TutorialStage += 1
	_, err = s.Model.UpdatePatient(ctx, p)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
