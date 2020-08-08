package meeting

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/model"
	"comet/pkg/utility"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteMeetingHandler struct {
	Model model.IModel
}

func (s *DeleteMeetingHandler) DeleteMeeting(ctx context.Context, req *pb.CommonDeleteRequest) (*pb.CommonMeetingResponse, error) {
	rslt, err := s.Model.DeleteMeeting(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeetingNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeetingToResponse(rslt)
	return resp, nil
}
