package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateChatRoom creates new chatRoom
func (m *Model) CreateChatRoom(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error) {

	// check if chatRoom exist
	_, err := m.chatRoomDAO.Get(ctx, chatRoom.ID)

	// only can create chatRoom if not found
	if err != nil && status.Code(err) == codes.Unknown {
		// create chatRoom
		s, err := m.chatRoomDAO.Create(ctx, chatRoom)
		if err != nil {
			return nil, err
		}

		// return result
		return s, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.ChatRoomAlreadyExistError
}

// UpdateChatRoom updates chatRoom
func (m *Model) UpdateChatRoom(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error) {
	// check if chatRoom exist
	s, err := m.chatRoomDAO.Get(ctx, chatRoom.ID)
	if err != nil {
		return nil, err
	}

	// patch chatRoom
	s.Blocked = chatRoom.Blocked

	// update chatRoom
	_, err = m.chatRoomDAO.Update(ctx, s)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// UpdateChatRooms update chatRooms
func (m *Model) UpdateChatRooms(ctx context.Context, chatRoom *dto.ChatRoom, ids []string) ([]string, error) {
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	chatRoom.ID = ids[0]
	s, err := m.UpdateChatRoom(ctx, chatRoom)
	if err != nil {
		return nil, err
	}

	return []string{s.ID}, nil
}

// GetChatRoom gets chatRoom by ID
func (m *Model) GetChatRoom(ctx context.Context, id string) (*dto.ChatRoom, error) {
	return m.chatRoomDAO.Get(ctx, id)
}

// BatchGetChatRooms get chatRooms by slice of IDs
func (m *Model) BatchGetChatRooms(ctx context.Context, ids []string) ([]*dto.ChatRoom, error) {
	return m.chatRoomDAO.BatchGet(ctx, ids)
}

// QueryChatRooms queries chatRooms by sort, range, filter
func (m *Model) QueryChatRooms(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatRoom, error) {
	return m.chatRoomDAO.Query(ctx, sort, itemsRange, filter)
}

// QueryByUsers ...
func (m *Model) QueryByUsers(ctx context.Context, users []string) ([]*dto.ChatRoom, error) {
	return m.chatRoomDAO.QueryByUsers(ctx, users)
}

// DeleteChatRoom deletes chatRoom by ID
func (m *Model) DeleteChatRoom(ctx context.Context, id string) (*dto.ChatRoom, error) {
	// check if chatRoom exist
	s, err := m.chatRoomDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete all messages related
	_, ms, err := m.chatMessageDAO.Query(ctx, nil, nil, map[string]interface{}{
		constants.RoomID: id,
	})
	if err == nil {
		for _, mm := range ms {
			_ = m.chatMessageDAO.Delete(ctx, mm.ID)
		}
	}

	// delete chatRoom
	err = m.chatRoomDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// DeleteChatRooms delete chatRooms by IDs
func (m *Model) DeleteChatRooms(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		s, err := m.DeleteChatRoom(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, s.ID)
	}

	return deletedIDs, nil
}
