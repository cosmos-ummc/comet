package activity

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
)

type GetActivitiesHandler struct {
	Model model.IModel
}

func (s *GetActivitiesHandler) GetActivities(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonActivitiesResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData
	filter := map[string]interface{}{}

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

	total, activities, err := s.Model.QueryActivities(ctx, sort, itemsRange, filter)
	if err != nil {
		logger.Log.Error("GetActivitiesHandler: " + err.Error())
		return nil, constants.InternalError
	}

	resp := utility.ActivitiesToResponses(activities)
	resp.Total = total
	return resp, nil
}
