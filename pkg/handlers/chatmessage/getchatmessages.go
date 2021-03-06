package chatMessage

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

type GetChatMessagesHandler struct {
	Model model.IModel
}

func (s *GetChatMessagesHandler) GetChatMessages(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonChatMessagesResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData
	filter := map[string]interface{}{}

	// If the request is batch get, call batch get model
	if len(req.Ids) > 0 {
		req.Ids = s.processReq(req.Ids)
		chatMessages, err := s.Model.BatchGetChatMessages(ctx, req.Ids)
		if err != nil {
			return nil, constants.InternalError
		}
		resp := utility.ChatMessagesToResponse(chatMessages)
		resp.Total = int64(len(chatMessages))
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

	total, chatMessages, err := s.Model.QueryChatMessages(ctx, sort, itemsRange, filter)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ChatMessageNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.ChatMessagesToResponse(chatMessages)
	resp.Total = total
	return resp, nil
}

func (s *GetChatMessagesHandler) processReq(ids []string) []string {
	// Ids is actually just ONE long string stored in a slice. The length of ids will always be 1
	// Protobuf doesn't know to split and what delimiter you use. So, split manually
	split := strings.Split(ids[0], ",")

	var normalised []string
	for _, id := range split {
		normalised = append(normalised, utility.NormalizeID(id))
	}

	return normalised
}
