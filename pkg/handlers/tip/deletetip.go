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

type DeleteTipHandler struct {
	Model model.IModel
}

func (s *DeleteTipHandler) DeleteTip(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonTipResponse, error) {
	rslt, err := s.Model.DeleteTip(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.TipNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.TipToResponse(rslt)
	return resp, nil
}
