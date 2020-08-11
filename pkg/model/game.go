package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateGame creates new game
func (m *Model) CreateGame(ctx context.Context, game *dto.Game) (*dto.Game, error) {

	// check if game exist
	_, err := m.gameDAO.Get(ctx, game.ID)

	// only can create game if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create game
		s, err := m.gameDAO.Create(ctx, game)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.GameAlreadyExistError
}

// UpdateGame updates game
func (m *Model) UpdateGame(ctx context.Context, game *dto.Game) (*dto.Game, error) {
	// check if game exist
	s, err := m.gameDAO.Get(ctx, game.ID)
	if err != nil {
		return nil, err
	}

	// patch game
	s.Type = game.Type
	s.ImgPath = game.ImgPath
	s.Link = game.Link

	// update game
	_, err = m.gameDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateGames update games
func (m *Model) UpdateGames(ctx context.Context, game *dto.Game, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	game.ID = ids[0]
	s, err := m.UpdateGame(ctx, game)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetGame gets game by ID
func (m *Model) GetGame(ctx context.Context, id string) (*dto.Game, error) {
	return m.gameDAO.Get(ctx, id)
}

// BatchGetGames get games by slice of IDs
func (m *Model) BatchGetGames(ctx context.Context, ids []string) ([]*dto.Game, error) {
	return m.gameDAO.BatchGet(ctx, ids)
}

// QueryGames queries games by sort, range, filter
func (m *Model) QueryGames(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Game, error) {
	return m.gameDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteGame deletes game by ID
func (m *Model) DeleteGame(ctx context.Context, id string) (*dto.Game, error) {
	// check if game exist
	s, err := m.gameDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete game
	err = m.gameDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteGames delete games by IDs
func (m *Model) DeleteGames(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteGame(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}
