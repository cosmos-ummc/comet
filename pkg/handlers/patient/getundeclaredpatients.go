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

type GetUndeclaredPatientsHandler struct {
	Model model.IModel
}

func (s *GetUndeclaredPatientsHandler) GetUndeclaredPatients(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonPatientsResponse, error) {
	timeString := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	t, err := utility.DateStringToTime(timeString)
	if err != nil {
		return nil, err
	}

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

	total, patients, err := s.Model.GetUndeclaredPatientsByTime(ctx, utility.TimeToMilli(t), sortData, rangeData)
	if err != nil {
		return nil, constants.InternalError
	}

	resp := utility.PatientsToResponse(patients)
	resp.Total = total
	return resp, nil
}
