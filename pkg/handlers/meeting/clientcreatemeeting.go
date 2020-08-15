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

type ClientCreateMeetingHandler struct {
	Model model.IModel
}

func (s *ClientCreateMeetingHandler) ClientCreateMeeting(ctx context.Context, req *pb.CommonMeetingRequest) (*pb.CommonMeetingResponse, error) {
	if req.Data == nil {
		return nil, constants.InvalidArgumentError
	}

	// get patient by id
	_, patients, err := s.Model.QueryPatients(ctx, nil, nil, map[string]interface{}{
		constants.UserId: req.Id,
	})
	if err != nil {
		return nil, err
	}
	if len(patients) == 0 {
		return nil, constants.InvalidArgumentError
	}
	patient := patients[0]

	// look for suitable consultant
	_, consultants, err := s.Model.QueryConsultants(ctx, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	utility.ShuffleConsultants(consultants)

	found := false
	var c *dto.Consultant
	for _, consultant := range consultants {
		if !utility.StringInSlice(req.Data.Time, consultant.TakenSlots) {
			consultant.TakenSlots = append(consultant.TakenSlots, req.Data.Time)
			c, err = s.Model.UpdateConsultant(ctx, consultant)
			if err != nil {
				return nil, err
			}
			found = true
			break
		}
	}

	if found {
		meeting := &dto.Meeting{
			ID:                    uuid.NewV4().String(),
			PatientID:             patient.ID,
			PatientPhoneNumber:    patient.PhoneNumber,
			PatientName:           patient.Name,
			ConsultantID:          c.ID,
			ConsultantName:        c.Name,
			ConsultantPhoneNumber: c.PhoneNumber,
			Time:                  req.Data.Time,
			Status:                req.Data.Status,
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
	return &pb.CommonMeetingResponse{}, nil

}
