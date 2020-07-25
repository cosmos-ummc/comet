package patient

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"strings"
)

type GetPatientsHandler struct {
	Model model.IModel
}

func (s *GetPatientsHandler) GetPatients(ctx context.Context, req *pb.CommonGetsRequest, user *dto.User) (*pb.CommonPatientsResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData
	filter := map[string]interface{}{}

	// If the request is batch get, call batch get model
	if len(req.Ids) > 0 && req.Ids[0] != "" {
		req.Ids = s.processReq(req.Ids)

		patients, err := s.Model.BatchGetPatients(ctx, req.Ids, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, constants.InternalError
		}

		resp := utility.PatientsToResponse(patients)
		resp.Total = int64(len(patients))
		return resp, nil
	}

	if req.Item != "" && req.Order != "" {
		sort = &dto.SortData{
			Item:  req.Item,
			Order: req.Order,
		}
	}

	if req.To != 0 {
		itemsRange = &dto.RangeData{
			From: int(req.From),
			To:   int(req.To),
		}
	}

	if req.FilterItem != "" && req.FilterValue != "" {
		filter[req.FilterItem] = req.FilterValue
	}

	total, patients, err := s.Model.QueryPatients(ctx, sort, itemsRange, filter, constants.UserPatientMap[user.Role])
	if err != nil {
		return nil, constants.InternalError
	}

	resp := utility.PatientsToResponse(patients)
	resp.Total = total
	return resp, nil
}

func (s *GetPatientsHandler) processReq(ids []string) []string {
	// Ids is actually just ONE long string stored in a slice. The length of ids will always be 1
	// Protobuf doesn't know to split and what delimiter you use. So, split manually
	split := strings.Split(ids[0], ",")

	var normalised []string
	for _, id := range split {
		normalised = append(normalised, utility.NormalizeID(id))
	}

	return normalised
}
