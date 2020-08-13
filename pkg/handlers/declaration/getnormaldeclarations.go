package declaration

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

type GetNormalDeclarationsHandler struct {
	Model model.IModel
}

func (s *GetNormalDeclarationsHandler) GetNormalDeclarations(ctx context.Context, req *pb.CommonGetsRequest) (*pb.CommonDeclarationsResponse, error) {
	var sort *dto.SortData
	var itemsRange *dto.RangeData

	// If the request is batch get, call batch get model
	if len(req.Ids) > 0 {
		req.Ids = s.processReq(req.Ids)
		declarations, err := s.Model.BatchGetDeclarations(ctx, req.Ids)
		if err != nil {
			if status.Code(err) == codes.Unknown {
				return nil, constants.DeclarationNotFoundError
			}
			return nil, constants.InternalError
		}
		resp := utility.DeclarationsToResponse(declarations)
		resp.Total = int64(len(declarations))
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

	total, declarations, err := s.Model.QueryDeclarationsByCategories(ctx, sort, itemsRange, req.FilterValue, []string{constants.DASS, constants.IESR})
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.DeclarationNotFoundError
		}
		return nil, constants.InternalError
	}

	resp := utility.DeclarationsToResponse(declarations)
	resp.Total = total
	return resp, nil
}

func (s *GetNormalDeclarationsHandler) processReq(ids []string) []string {
	// Ids is actually just ONE long string stored in a slice. The length of ids will always be 1
	// Protobuf doesn't know to split and what delimiter you use. So, split manually
	split := strings.Split(ids[0], ",")

	var normalised []string
	for _, id := range split {
		normalised = append(normalised, utility.NormalizeID(id))
	}

	return normalised
}
