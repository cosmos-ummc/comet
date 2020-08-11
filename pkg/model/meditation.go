package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateMeditation creates new meditation
func (m *Model) CreateMeditation(ctx context.Context, meditation *dto.Meditation) (*dto.Meditation, error) {

	// check if meditation exist
	_, err := m.meditationDAO.Get(ctx, meditation.ID)

	// only can create meditation if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create meditation
		s, err := m.meditationDAO.Create(ctx, meditation)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.MeditationAlreadyExistError
}

// UpdateMeditation updates meditation
func (m *Model) UpdateMeditation(ctx context.Context, meditation *dto.Meditation) (*dto.Meditation, error) {
	// check if meditation exist
	s, err := m.meditationDAO.Get(ctx, meditation.ID)
	if err != nil {
		return nil, err
	}

	// patch meditation
	s.Link = meditation.Link

	// update meditation
	_, err = m.meditationDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateMeditations update meditations
func (m *Model) UpdateMeditations(ctx context.Context, meditation *dto.Meditation, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	meditation.ID = ids[0]
	s, err := m.UpdateMeditation(ctx, meditation)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetMeditation gets meditation by ID
func (m *Model) GetMeditation(ctx context.Context, id string) (*dto.Meditation, error) {
	return m.meditationDAO.Get(ctx, id)
}

// BatchGetMeditations get meditations by slice of IDs
func (m *Model) BatchGetMeditations(ctx context.Context, ids []string) ([]*dto.Meditation, error) {
	return m.meditationDAO.BatchGet(ctx, ids)
}

// QueryMeditations queries meditations by sort, range, filter
func (m *Model) QueryMeditations(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meditation, error) {
	return m.meditationDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteMeditation deletes meditation by ID
func (m *Model) DeleteMeditation(ctx context.Context, id string) (*dto.Meditation, error) {
	// check if meditation exist
	s, err := m.meditationDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete meditation
	err = m.meditationDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteMeditations delete meditations by IDs
func (m *Model) DeleteMeditations(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteMeditation(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}
