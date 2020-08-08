package meeting

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

type UpdateMeetingHandler struct {
	Model model.IModel
}

func (s *UpdateMeetingHandler) UpdateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	meeting := s.reqToMeeting(req)

	v, err := s.Model.UpdateMeeting(ctx, meeting)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeetingNotFoundError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeetingToResponse(v)
	return resp, nil
}

func (s *UpdateMeetingHandler) reqToMeeting(req *pb.CommonMeetingRequest) *dto.Meeting {
	meeting := &dto.Meeting{
		ID:           utility.RemoveZeroWidth(req.Id),
		PatientID:    req.Data.PatientId,
		ConsultantID: req.Data.ConsultantId,
		Time:         req.Data.Time,
		Status:       req.Data.Status,
	}
	return meeting
}
