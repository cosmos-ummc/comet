package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateConsultant creates new consultant
func (m *Model) CreateConsultant(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error) {

	// check if consultant exist
	_, err := m.consultantDAO.Get(ctx, consultant.ID)

	// only can create consultant if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create consultant
		s, err := m.consultantDAO.Create(ctx, consultant)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.ConsultantAlreadyExistError
}

// UpdateConsultant updates consultant
func (m *Model) UpdateConsultant(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error) {
	// check if consultant exist
	s, err := m.consultantDAO.Get(ctx, consultant.ID)
	if err != nil {
		return nil, err
	}

	// patch consultant

	// update consultant
	_, err = m.consultantDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateConsultants update consultants
func (m *Model) UpdateConsultants(ctx context.Context, consultant *dto.Consultant, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	consultant.ID = ids[0]
	s, err := m.UpdateConsultant(ctx, consultant)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetConsultant gets consultant by ID
func (m *Model) GetConsultant(ctx context.Context, id string) (*dto.Consultant, error) {
	return m.consultantDAO.Get(ctx, id)
}

// BatchGetConsultants get consultants by slice of IDs
func (m *Model) BatchGetConsultants(ctx context.Context, ids []string) ([]*dto.Consultant, error) {
	return m.consultantDAO.BatchGet(ctx, ids)
}

// QueryConsultants queries consultants by sort, range, filter
func (m *Model) QueryConsultants(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Consultant, error) {
	return m.consultantDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteConsultant deletes consultant by ID
func (m *Model) DeleteConsultant(ctx context.Context, id string) (*dto.Consultant, error) {
	// check if consultant exist
	s, err := m.consultantDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete consultant
	err = m.consultantDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteConsultants delete consultants by IDs
func (m *Model) DeleteConsultants(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteConsultant(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}

// QueryConsultantsByPatientID ...
func (m *Model) QueryConsultantsByPatientID(ctx context.Context, id string) ([]*dto.Consultant, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	_, consultants, err := m.consultantDAO.Query(ctx, nil, nil, filter)
	return consultants, err
}
