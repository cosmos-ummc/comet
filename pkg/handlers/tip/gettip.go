package tip

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetTipHandler struct {
	Model model.IModel
}

func (s *GetTipHandler) GetTip(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonTipResponse, error) {
	tip, err := s.Model.GetTip(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.TipNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.TipToResponse(tip)
	return resp, nil
}
