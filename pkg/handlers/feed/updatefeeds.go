package feed

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateFeedsHandler struct {
	Model model.IModel
}

func (s *UpdateFeedsHandler) UpdateFeeds(ctx context.Context, req *pb.CommonFeedsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	feed := s.reqToFeed(req)

	ids, err := s.Model.UpdateFeeds(ctx, feed, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.FeedNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateFeedsHandler) reqToFeed(req *pb.CommonFeedsRequest) *dto.Feed {
	feed := &dto.Feed{
		Title:       req.Data.Title,
		Description: req.Data.Description,
		Link:        req.Data.Link,
		ImgPath:     req.Data.ImgPath,
		Type:        req.Data.Type,
	}
	return feed
}
