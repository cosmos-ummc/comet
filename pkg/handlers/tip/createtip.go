package tip

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CreateTipHandler struct {
	Model model.IModel
}

func (s *CreateTipHandler) CreateTip(ctx context.Context, req *pb.CommonTipRequest) (*pb.CommonTipResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	tip := &dto.Tip{
		ID:          uuid.NewV4().String(),
		Title:       req.Data.Title,
		Description: req.Data.Description,
	}

	rslt, err := s.Model.CreateTip(ctx, tip)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.TipAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.TipToResponse(rslt)
	return resp, nil
}
