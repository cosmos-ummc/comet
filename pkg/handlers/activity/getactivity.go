package activity

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GetActivityHandler struct {
	Model model.IModel
}

func (s *GetActivityHandler) GetActivity(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonActivityResponse, error) {
	activity, err := s.Model.GetActivity(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.ActivityNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.ActivityToResponse(activity)
	return resp, nil
}
