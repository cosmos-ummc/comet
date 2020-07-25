package model

import (
	"comet/pkg/dto"
	"context"
)

// CreateActivity creates new activity
func (m *Model) CreateActivity(ctx context.Context, activity *dto.Activity) (*dto.Activity, error) {
	// create activity
	s, err := m.activityDAO.Create(ctx, activity)
	if err != nil {
		return nil, err
	}

	// return result
	return s, nil
}

// GetActivity gets activity by ID
func (m *Model) GetActivity(ctx context.Context, id string) (*dto.Activity, error) {
	return m.activityDAO.Get(ctx, id)
}

// QueryActivities queries activities by sort, range, filter
func (m *Model) QueryActivities(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Activity, error) {
	return m.activityDAO.Query(ctx, sort, itemsRange, filter)
}
