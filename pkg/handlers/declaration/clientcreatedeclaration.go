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

	// compute message by status
	severeList := []int64{constants.DeclarationModerate, constants.DeclarationSevere, constants.DeclarationExtremelySevere}
	var list []string
	r := int64(1)

	if d.Category == constants.DASS {
		if utility.IntInSlice(d.DepressionStatus, severeList) {
			list = append(list, "depression")
		}
		if utility.IntInSlice(d.AnxietyStatus, severeList) {
			list = append(list, "anxiety")
		}
		if utility.IntInSlice(d.StressStatus, severeList) {
			list = append(list, "stress")
		}
		if len(list) == 0 {
			r = 1
		}
		if len(list) == 3 {
			r = 9
		}
		if len(list) == 1 {
			if utility.StringInSlice("depression", list) {
				r = 3
			} else if utility.StringInSlice("anxiety", list) {
				r = 4
			} else if utility.StringInSlice("stress", list) {
				r = 5
			}
		} else {
			if utility.StringInSlice("depression", list) && utility.StringInSlice("anxiety", list) {
				r = 6
			} else if utility.StringInSlice("depression", list) && utility.StringInSlice("stress", list) {
				r = 7
			} else {
				r = 8
			}
		}
	} else if d.Category == constants.IESR {
		if utility.IntInSlice(d.PtsdStatus, severeList) {
			r = 2
		} else {
			r = 1
		}
	} else {
		if utility.IntInSlice(d.DailyStatus, severeList) {
			r = 2
		} else {
			r = 1
		}
	}

	resp := &pb.ClientCreateDeclarationResponse{
		HasSymptom: r,
	}
	return resp, nil
}
