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

type GetMeetingHandler struct {
	Model model.IModel
}

func (s *GetMeetingHandler) GetMeeting(ctx context.Context, req *pb.CommonGetRequest) (*pb.CommonMeetingResponse, error) {
	meeting, err := s.Model.GetMeeting(ctx, req.Id)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeetingNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeetingToResponse(meeting)
	return resp, nil
}
