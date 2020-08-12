package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateTip creates new tip
func (m *Model) CreateTip(ctx context.Context, tip *dto.Tip) (*dto.Tip, error) {

	// check if tip exist
	_, err := m.tipDAO.Get(ctx, tip.ID)

	// only can create tip if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create tip
		s, err := m.tipDAO.Create(ctx, tip)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.TipAlreadyExistError
}

// UpdateTip updates tip
func (m *Model) UpdateTip(ctx context.Context, tip *dto.Tip) (*dto.Tip, error) {
	// check if tip exist
	s, err := m.tipDAO.Get(ctx, tip.ID)
	if err != nil {
		return nil, err
	}

	// patch tip
	s.Title = tip.Title
	s.Description = tip.Description

	// update tip
	_, err = m.tipDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateTips update tips
func (m *Model) UpdateTips(ctx context.Context, tip *dto.Tip, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	tip.ID = ids[0]
	s, err := m.UpdateTip(ctx, tip)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetTip gets tip by ID
func (m *Model) GetTip(ctx context.Context, id string) (*dto.Tip, error) {
	return m.tipDAO.Get(ctx, id)
}

// BatchGetTips get tips by slice of IDs
func (m *Model) BatchGetTips(ctx context.Context, ids []string) ([]*dto.Tip, error) {
	return m.tipDAO.BatchGet(ctx, ids)
}

// QueryTips queries tips by sort, range, filter
func (m *Model) QueryTips(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Tip, error) {
	return m.tipDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteTip deletes tip by ID
func (m *Model) DeleteTip(ctx context.Context, id string) (*dto.Tip, error) {
	// check if tip exist
	s, err := m.tipDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete tip
	err = m.tipDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteTips delete tips by IDs
func (m *Model) DeleteTips(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteTip(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}
