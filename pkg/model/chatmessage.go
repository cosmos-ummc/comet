package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateChatMessage creates new chatMessage
func (m *Model) CreateChatMessage(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error) {

	// check if chatMessage exist
	_, err := m.chatMessageDAO.Get(ctx, chatMessage.ID)

	// only can create chatMessage if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create chatMessage
		s, err := m.chatMessageDAO.Create(ctx, chatMessage)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.ChatMessageAlreadyExistError
}

// UpdateChatMessage updates chatMessage
func (m *Model) UpdateChatMessage(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error) {
	// check if chatMessage exist
	s, err := m.chatMessageDAO.Get(ctx, chatMessage.ID)
	if err != nil {
		return nil, err
	}

	// patch chatMessage
	s.Content = chatMessage.Content

	// update chatMessage
	_, err = m.chatMessageDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateChatMessages update chatMessages
func (m *Model) UpdateChatMessages(ctx context.Context, chatMessage *dto.ChatMessage, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	chatMessage.ID = ids[0]
	s, err := m.UpdateChatMessage(ctx, chatMessage)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetChatMessage gets chatMessage by ID
func (m *Model) GetChatMessage(ctx context.Context, id string) (*dto.ChatMessage, error) {
	return m.chatMessageDAO.Get(ctx, id)
}

// BatchGetChatMessages get chatMessages by slice of IDs
func (m *Model) BatchGetChatMessages(ctx context.Context, ids []string) ([]*dto.ChatMessage, error) {
	return m.chatMessageDAO.BatchGet(ctx, ids)
}

// QueryChatMessages queries chatMessages by sort, range, filter
func (m *Model) QueryChatMessages(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatMessage, error) {
	return m.chatMessageDAO.Query(ctx, sort, itemsRange, filter)
}

// DeleteChatMessage deletes chatMessage by ID
func (m *Model) DeleteChatMessage(ctx context.Context, id string) (*dto.ChatMessage, error) {
	// check if chatMessage exist
	s, err := m.chatMessageDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete chatMessage
	err = m.chatMessageDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteChatMessages delete chatMessages by IDs
func (m *Model) DeleteChatMessages(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteChatMessage(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}
