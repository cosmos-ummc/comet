package feed

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

type UpdateFeedHandler struct {
	Model model.IModel
}

func (s *UpdateFeedHandler) UpdateFeed(ctx context.Context, req *pb.CommonFeedRequest) (*pb.CommonFeedResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	feed := s.reqToFeed(req)

	v, err := s.Model.UpdateFeed(ctx, feed)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.FeedNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.FeedToResponse(v)
	return resp, nil
}

func (s *UpdateFeedHandler) reqToFeed(req *pb.CommonFeedRequest) *dto.Feed {
	feed := &dto.Feed{
		ID:          utility.RemoveZeroWidth(req.Id),
		Title:       req.Data.Title,
		Description: req.Data.Description,
		Link:        req.Data.Link,
		ImgPath:     req.Data.ImgPath,
		Type:        req.Data.Type,
	}
	return feed
}
