package feed

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteFeedHandler struct {
	Model model.IModel
}

func (s *DeleteFeedHandler) DeleteFeed(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonFeedResponse, error) {
	rslt, err := s.Model.DeleteFeed(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.FeedNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.FeedToResponse(rslt)
	return resp, nil
}
