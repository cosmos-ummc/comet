package meeting

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

type CreateMeetingHandler struct {
	Model model.IModel
}

func (s *CreateMeetingHandler) CreateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}
	meeting := &dto.Meeting{
		ID:           uuid.NewV4().String(),
		PatientID:    req.Data.PatientId,
		ConsultantID: req.Data.ConsultantId,
		Time:         req.Data.Time,
		Status:       req.Data.Status,
	}

	rslt, err := s.Model.CreateMeeting(ctx, meeting)
	if err != nil {
		if status.Code(err) == codes.AlreadyExists {
			return nil, constants.MeetingAlreadyExistError
		}
		return nil, constants.InternalError
	}
	resp := utility.MeetingToResponse(rslt)
	return resp, nil
}
