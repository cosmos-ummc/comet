package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"time"
)

type GetNoCallPatientsHandler struct {
	Model model.IModel
}

func (s *GetNoCallPatientsHandler) GetNoCallPatients(ctx context.Context, req *pb.CommonGetsRequest, user *dto.User) (*pb.CommonPatientsResponse, error) {
	t, err := utility.DateStringToTime(utility.TimeToDateString(utility.MalaysiaTime(time.Now())))
	if err != nil {
		return nil, constants.InternalError
	}
	timeMilli := utility.TimeToMilli(t)
	var sortData *dto.SortData
	var rangeData *dto.RangeData
	if req.Item != "" && req.Order != "" {
		sortData = &dto.SortData{
			Item:  req.Item,
			Order: req.Order,
		}
	}
	if req.To != 0 {
		rangeData = &dto.RangeData{
			From: int(req.From),
			To:   int(req.To),
		}
	}

	count, patients, err := s.Model.GetNoCallPatients(ctx, timeMilli, sortData, rangeData, constants.UserPatientMap[user.Role])

	resp := utility.PatientsToResponse(patients)
	resp.Total = count
	return resp, nil
}
