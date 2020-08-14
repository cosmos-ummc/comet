package declaration

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"
	"time"
)

type ClientCreateDeclarationHandler struct {
	Model model.IModel
}

func (s *ClientCreateDeclarationHandler) ClientCreateDeclaration(ctx context.Context, req *pb.ClientCreateDeclarationRequest) (*pb.ClientCreateDeclarationResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}

	var result []*dto.Question
	for k, v := range req.Data {
		q, err := s.Model.GetQuestion(ctx, k)
		if err != nil {
			continue
		}
		q.Score = v
		result = append(result, q)
	}

	declaration := &dto.Declaration{
		ID:          uuid.NewV4().String(),
		PatientID:   req.PatientId,
		Result:      result,
		SubmittedAt: utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
	}

	d, err := s.Model.ClientCreateDeclaration(ctx, declaration)
	if err != nil {
		return nil, err
	}

	// check if the declaration is the first
	total, _, err := s.Model.QueryDeclarationsByCategories(ctx, nil, nil, req.PatientId, []string{declaration.Category})
	if total == 1 {
		p, err := s.Model.GetPatient(ctx, req.PatientId)
		if err == nil {
			switch declaration.Category {
			case constants.DASS:
				_ = utility.SendBotNotification(p.TelegramID, constants.FirstDassMessage)
			case constants.IESR:
				_ = utility.SendBotNotification(p.TelegramID, constants.FirstIesrMessage)
			default:
				_ = utility.SendBotNotification(p.TelegramID, constants.FirstDailyMessage)
			}
		}
	}

	hasSymptom := false
	if d.StressStatus == constants.DeclarationSevere || d.StressStatus == constants.DeclarationExtremelySevere || d.StressStatus == constants.DeclarationModerate ||
		d.DepressionStatus == constants.DeclarationSevere || d.DepressionStatus == constants.DeclarationExtremelySevere || d.DepressionStatus == constants.DeclarationModerate ||
		d.AnxietyStatus == constants.DeclarationSevere || d.AnxietyStatus == constants.DeclarationExtremelySevere || d.AnxietyStatus == constants.DeclarationModerate ||
		d.PtsdStatus == constants.DeclarationSevere || d.PtsdStatus == constants.DeclarationExtremelySevere || d.PtsdStatus == constants.DeclarationModerate {
		hasSymptom = true
	}

	// if hasSymptom, notify user to create meeting
	if hasSymptom {
		// get patient
		p, err := s.Model.GetPatient(ctx, req.PatientId)
		if err != nil {
			return nil, err
		}

		u, err := s.Model.GetUser(ctx, p.UserID)
		if err != nil {
			return nil, err
		}

		u.InvitedToMeeting = true
		_, err = s.Model.UpdateUser(ctx, u)
		if err != nil {
			return nil, err
		}
	}

	r := int64(1)
	if hasSymptom {
		r = 2
	}

	resp := &pb.ClientCreateDeclarationResponse{
		HasSymptom: r,
	}
	return resp, nil
}
