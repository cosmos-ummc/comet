package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateFeed creates new feed
func (m *Model) CreateFeed(ctx context.Context, feed *dto.Feed) (*dto.Feed, error) {

	// check if feed exist
	_, err := m.feedDAO.Get(ctx, feed.ID)

	// only can create feed if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create feed
		s, err := m.feedDAO.Create(ctx, feed)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.FeedAlreadyExistError
}

// UpdateFeed updates feed
func (m *Model) UpdateFeed(ctx context.Context, feed *dto.Feed) (*dto.Feed, error) {
	// check if feed exist
	s, err := m.feedDAO.Get(ctx, feed.ID)
	if err != nil {
		return nil, err
	}

	// patch feed
	s.Type = feed.Type
	s.Title = feed.Title
	s.Description = feed.Description
	s.ImgPath = feed.ImgPath
	s.Link = feed.Link

	// update feed
	_, err = m.feedDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateFeeds update feeds
func (m *Model) UpdateFeeds(ctx context.Context, feed *dto.Feed, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	feed.ID = ids[0]
	s, err := m.UpdateFeed(ctx, feed)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetFeed gets feed by ID
func (m *Model) GetFeed(ctx context.Context, id string) (*dto.Feed, error) {
	return m.feedDAO.Get(ctx, id)
}

// BatchGetFeeds get feeds by slice of IDs
func (m *Model) BatchGetFeeds(ctx context.Context, ids []string) ([]*dto.Feed, error) {
	return m.feedDAO.BatchGet(ctx, ids)
}

// QueryFeeds queries feeds by sort, range, filter
func (m *Model) QueryFeeds(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Feed, error) {
	return m.feedDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteFeed deletes feed by ID
func (m *Model) DeleteFeed(ctx context.Context, id string) (*dto.Feed, error) {
	// check if feed exist
	s, err := m.feedDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete feed
	err = m.feedDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteFeeds delete feeds by IDs
func (m *Model) DeleteFeeds(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteFeed(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}
