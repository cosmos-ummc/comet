package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMeeting creates new meeting
func (m *Model) CreateMeeting(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error) {

	// check if meeting exist
	_, err := m.meetingDAO.Get(ctx, meeting.ID)

	// only can create meeting if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create meeting
		s, err := m.meetingDAO.Create(ctx, meeting)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.MeetingAlreadyExistError
}

// UpdateMeeting updates meeting
func (m *Model) UpdateMeeting(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error) {
	// check if meeting exist
	s, err := m.meetingDAO.Get(ctx, meeting.ID)
	if err != nil {
		return nil, err
	}

	// patch meeting
	s.Status = meeting.Status
	s.Time = meeting.Time
	s.ConsultantID = meeting.ConsultantID

	// update meeting
	_, err = m.meetingDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateMeetings update meetings
func (m *Model) UpdateMeetings(ctx context.Context, meeting *dto.Meeting, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	meeting.ID = ids[0]
	s, err := m.UpdateMeeting(ctx, meeting)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetMeeting gets meeting by ID
func (m *Model) GetMeeting(ctx context.Context, id string) (*dto.Meeting, error) {
	return m.meetingDAO.Get(ctx, id)
}

// BatchGetMeetings get meetings by slice of IDs
func (m *Model) BatchGetMeetings(ctx context.Context, ids []string) ([]*dto.Meeting, error) {
	return m.meetingDAO.BatchGet(ctx, ids)
}

// QueryMeetings queries meetings by sort, range, filter
func (m *Model) QueryMeetings(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meeting, error) {
	return m.meetingDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteMeeting deletes meeting by ID
func (m *Model) DeleteMeeting(ctx context.Context, id string) (*dto.Meeting, error) {
	// check if meeting exist
	s, err := m.meetingDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete meeting
	err = m.meetingDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteMeetings delete meetings by IDs
func (m *Model) DeleteMeetings(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteMeeting(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}

// QueryMeetingsByPatientID ...
func (m *Model) QueryMeetingsByPatientID(ctx context.Context, id string) ([]*dto.Meeting, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	_, meetings, err := m.meetingDAO.Query(ctx, nil, nil, filter)
	return meetings, err
}
