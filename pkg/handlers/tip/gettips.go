package tip

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

type GetTipsHandler struct {
	Model model.IModel
}

func (s *GetTipsHandler) GetTips(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonTipsResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData
	filter := map[string]interface{}{}

	// If the request is batch get, call batch get model
	if len(req.Ids) > 0 {
		req.Ids = s.processReq(req.Ids)
		tips, err := s.Model.BatchGetTips(ctx, req.Ids)
		if err != nil {
			return nil, constants.InternalError
		}
		resp := utility.TipsToResponse(tips)
		resp.Total = int64(len(tips))
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

	total, tips, err := s.Model.QueryTips(ctx, sort, itemsRange, filter)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.TipNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.TipsToResponse(tips)
	resp.Total = total
	return resp, nil
}

func (s *GetTipsHandler) processReq(ids []string) []string {
	split := strings.Split(ids[0], ",")
	return split
}
