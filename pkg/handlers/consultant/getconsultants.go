package consultant

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

type GetConsultantsHandler struct {
	Model model.IModel
}

func (s *GetConsultantsHandler) GetConsultants(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonConsultantsResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData
	filter := map[string]interface{}{}

	// If the request is batch get, call batch get model
	if len(req.Ids) > 0 {
		req.Ids = s.processReq(req.Ids)
		consultants, err := s.Model.BatchGetConsultants(ctx, req.Ids)
		if err != nil {
			return nil, constants.InternalError
		}
		resp := utility.ConsultantsToResponse(consultants)
		resp.Total = int64(len(consultants))
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

	total, consultants, err := s.Model.QueryConsultants(ctx, sort, itemsRange, filter)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ConsultantNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.ConsultantsToResponse(consultants)
	resp.Total = total
	return resp, nil
}

func (s *GetConsultantsHandler) processReq(ids []string) []string {
	// Ids is actually just ONE long string stored in a slice. The length of ids will always be 1
	// Protobuf doesn't know to split and what delimiter you use. So, split manually
	split := strings.Split(ids[0], ",")

	var normalised []string
	for _, id := range split {
		normalised = append(normalised, utility.NormalizeID(id))
	}

	return normalised
}
