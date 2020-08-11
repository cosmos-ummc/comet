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

type GetFeedHandler struct {
	Model model.IModel
}

func (s *GetFeedHandler) GetFeed(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonFeedResponse, error) {
	feed, err := s.Model.GetFeed(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.FeedNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.FeedToResponse(feed)
	return resp, nil
}
