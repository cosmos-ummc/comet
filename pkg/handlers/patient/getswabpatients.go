package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetSwabPatientsHandler struct {
	Model model.IModel
}

func (s *GetSwabPatientsHandler) GetSwabPatients(ctx context.Context, req *pb.CommonGetsRequest, user *dto.User) (*pb.CommonPatientsResponse, error) {
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

	total, patients, err := s.Model.GetSwabPatients(ctx, sortData, rangeData, constants.UserPatientMap[user.Role])
	if err != nil {
		return nil, constants.InternalError
	}

	resp := utility.PatientsToResponse(patients)
	resp.Total = total
	return resp, nil
}
