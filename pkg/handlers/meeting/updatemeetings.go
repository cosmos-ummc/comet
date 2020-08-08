package meeting

import (
	pb "comet/pkg/api"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/model"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UpdateMeetingsHandler struct {
	Model model.IModel
}

func (s *UpdateMeetingsHandler) UpdateMeetings(ctx context.Context, req *pb.CommonMeetingsRequest) (*pb.CommonIdsResponse, error) {
	if len(req.Ids) != 1 {
		return nil, constants.OperationUnsupportedError
	}
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	meeting := s.reqToMeeting(req)

	ids, err := s.Model.UpdateMeetings(ctx, meeting, req.Ids)
	if err != nil {
		if status.Code(err) == codes.Unknown {
			return nil, constants.MeetingNotFoundError
		}
		if status.Code(err) == codes.Unimplemented {
			return nil, err
		}
		return nil, constants.InternalError
	}
	return &pb.CommonIdsResponse{Data: ids}, nil
}

func (s *UpdateMeetingsHandler) reqToMeeting(req *pb.CommonMeetingsRequest) *dto.Meeting {
	meeting := &dto.Meeting{
		ID:           req.Ids[0],
		PatientID:    req.Data.PatientId,
		ConsultantID: req.Data.ConsultantId,
		Time:         req.Data.Time,
		Status:       req.Data.Status,
	}
	return meeting
}
