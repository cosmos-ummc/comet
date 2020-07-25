package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetOtherPatientsHandler struct {
	Model model.IModel
}

func (s *GetOtherPatientsHandler) GetOtherPatients(ctx context.Context, req *pb.CommonGetsRequest, user *dto.User) (*pb.CommonPatientsResponse, error) {

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

	count, patients, err := s.Model.GetPatientsByStatus(ctx, []int64{constants.ConfirmedAndAdmitted, constants.Completed, constants.Quit, constants.Recovered, constants.PassedAway},
		sortData, rangeData, constants.UserPatientMap[user.Role])
	if err != nil {
		logger.Log.Error("GetOtherPatients: " + err.Error())
		return nil, constants.InternalError
	}

	resp := utility.PatientsToResponse(patients)
	resp.Total = count
	return resp, nil
}
