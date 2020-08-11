package feed

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

type CreateFeedHandler struct {
	Model model.IModel
}

func (s *CreateFeedHandler) CreateFeed(ctx context.Context, req *pb.CommonFeedRequest) (*pb.CommonFeedResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	feed := &dto.Feed{
		ID:          uuid.NewV4().String(),
		Title:       req.Data.Title,
		Description: req.Data.Description,
		Link:        req.Data.Link,
		ImgPath:     req.Data.ImgPath,
		Type:        req.Data.Type,
	}

	rslt, err := s.Model.CreateFeed(ctx, feed)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.FeedAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.FeedToResponse(rslt)
	return resp, nil
}
