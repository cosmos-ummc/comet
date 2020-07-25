package swab

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetSwabsHandler struct {
	Model model.IModel
}

func (s *GetSwabsHandler) GetSwabs(ctx context.Context, req *pb.CommonGetsRequest, user *dto.User) (*pb.CommonSwabsResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData
	filter := map[string]interface{}{}

	// If the request is batch get, call batch get model
	if len(req.Ids) > 0 {
		req.Ids = s.processReq(req.Ids)
		swabs, err := s.Model.BatchGetSwabs(ctx, req.Ids, constants.UserPatientMap[user.Role])
		if err != nil {
			return nil, constants.InternalError
		}
		resp := utility.SwabsToResponse(swabs)
		resp.Total = int64(len(swabs))
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

	total, swabs, err := s.Model.QuerySwabs(ctx, sort, itemsRange, filter, constants.UserPatientMap[user.Role])
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.SwabNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.SwabsToResponse(swabs)
	resp.Total = total
	return resp, nil
}

func (s *GetSwabsHandler) processReq(ids []string) []string {
	// Ids is actually just ONE long string stored in a slice. The length of ids will always be 1
	// Protobuf doesn't know to split and what delimiter you use. So, split manually
	split := strings.Split(ids[0], ",")

	var normalised []string
	for _, id := range split {
		normalised = append(normalised, utility.NormalizeID(id))
	}

	return normalised
}
