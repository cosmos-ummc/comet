package tip

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateTipsHandler struct {
	Model model.IModel
}

func (s *UpdateTipsHandler) UpdateTips(ctx context.Context, req *pb.CommonTipsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	tip := s.reqToTip(req)

	ids, err := s.Model.UpdateTips(ctx, tip, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.TipNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateTipsHandler) reqToTip(req *pb.CommonTipsRequest) *dto.Tip {
	tip := &dto.Tip{
		Title:       req.Data.Title,
		Description: req.Data.Description,
	}
	return tip
}
