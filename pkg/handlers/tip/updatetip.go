package tip

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateTipHandler struct {
	Model model.IModel
}

func (s *UpdateTipHandler) UpdateTip(ctx context.Context, req *pb.CommonTipRequest) (*pb.CommonTipResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	tip := s.reqToTip(req)

	v, err := s.Model.UpdateTip(ctx, tip)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.TipNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.TipToResponse(v)
	return resp, nil
}

func (s *UpdateTipHandler) reqToTip(req *pb.CommonTipRequest) *dto.Tip {
	tip := &dto.Tip{
		ID:          utility.RemoveZeroWidth(req.Id),
		Title:       req.Data.Title,
		Description: req.Data.Description,
	}
	return tip
}
