package dao

import (
	"context"

	"comet/pkg/dto"
)

// IPatientDAO ...
type IPatientDAO interface {
	// Create creates new patient
	Create(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// Get gets patient by ID
	Get(ctx context.Context, id string) (*dto.Patient, error)
	// BatchGet gets patients by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.Patient, error)
	// Query queries patients by sort, range, type and a filter to match any number of fields
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Patient, error)
	// GetByStatus gets patients by statuses and type
	GetByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error)
	// GetUndeclaredByTime gets undeclared patients of the specified type given from timestamp
	GetUndeclaredByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error)
	// Update updates patient
	Update(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// Delete deletes patient by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes patients by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
	// ClientGetUndeclaredByTime gets undeclared patients given from timestamp (with telegramID)
	ClientGetUndeclaredByTime(ctx context.Context, from int64) ([]*dto.Patient, error)
	// GetByConsentTime gets patients given from and to consent timestamp
	GetByConsentTime(ctx context.Context, from int64, to int64) ([]*dto.Patient, error)
}

// IQuestionDAO ...
type IQuestionDAO interface {
	// Create creates new question
	Create(ctx context.Context, question *dto.Question) (*dto.Question, error)
	// Update updates question
	Update(ctx context.Context, question *dto.Question) (*dto.Question, error)
	// Get gets question by ID
	Get(ctx context.Context, id string) (*dto.Question, error)
	// BatchGet gets questions by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.Question, error)
	// Query queries questions by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Question, error)
	// Delete deletes question by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes questions by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IFeedDAO ...
type IFeedDAO interface {
	// Create creates new feed
	Create(ctx context.Context, feed *dto.Feed) (*dto.Feed, error)
	// Update updates feed
	Update(ctx context.Context, feed *dto.Feed) (*dto.Feed, error)
	// Get gets feed by ID
	Get(ctx context.Context, id string) (*dto.Feed, error)
	// BatchGet gets feeds by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.Feed, error)
	// Query queries feeds by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Feed, error)
	// Delete deletes feed by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes feeds by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IDeclarationDAO ...
type IDeclarationDAO interface {
	// Create creates new declaration
	Create(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// Get gets declaration
	Get(ctx context.Context, declarationID string) (*dto.Declaration, error)
	// Update updates declaration
	Update(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// Delete deletes declaration by ID
	Delete(ctx context.Context, declarationID string) error
	// Query queries declarations
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Declaration, error)
	// BatchGet gets declarations
	BatchGet(ctx context.Context, ids []string) ([]*dto.Declaration, error)
	// BatchDelete deletes declarations
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IUserDAO ...
type IUserDAO interface {
	// Create creates new user
	Create(ctx context.Context, user *dto.User) (*dto.User, error)
	// Update updates user
	Update(ctx context.Context, user *dto.User) (*dto.User, error)
	// Get gets user by ID
	Get(ctx context.Context, id string) (*dto.User, error)
	// BatchGet gets users by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.User, error)
	// Query queries users by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter *dto.FilterData, superuserOnly bool) (int64, []*dto.User, error)
	// Delete deletes user by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes users by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

type IReportDAO interface {
	// Create creates report
	Create(ctx context.Context, date string, declarationReport *dto.Report) error
	// UpdateByFields decrease decrementField by 1 for given date and increase incrementField by 1
	UpdateByFields(ctx context.Context, date, decrementField, incrementField string) error
	// Get gets report by date
	Get(ctx context.Context, date string) (*dto.Report, error)
	// BatchGet gets report by dates
	BatchGet(ctx context.Context, dates []string) ([]*dto.Report, error)
}

type IAuthDAO interface {
	InitIndex(ctx context.Context) error
	// Create creates new auth token
	Create(ctx context.Context, auth *dto.AuthObject) (*dto.AuthObject, error)
	// Get gets auth token
	Get(ctx context.Context, token string) (*dto.AuthObject, error)
	// Delete deletes user by token
	Delete(ctx context.Context, token string) error
	// DeleteByID deletes tokens by userID
	DeleteByID(ctx context.Context, id string) error
}

// IChatMessageDAO ...
type IChatMessageDAO interface {
	// Create creates new chatMessage
	Create(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error)
	// Update updates chatMessage
	Update(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error)
	// Get gets chatMessage by ID
	Get(ctx context.Context, id string) (*dto.ChatMessage, error)
	// BatchGet gets chatMessages by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.ChatMessage, error)
	// Query queries chatMessages by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatMessage, error)
	// Delete deletes chatMessage by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes chatMessages by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IChatRoomDAO ...
type IChatRoomDAO interface {
	// Create creates new chatRoom
	Create(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error)
	// Update updates chatRoom
	Update(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error)
	// Get gets chatRoom by ID
	Get(ctx context.Context, id string) (*dto.ChatRoom, error)
	// BatchGet gets chatRooms by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.ChatRoom, error)
	// Query queries chatRooms by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatRoom, error)
	// Delete deletes chatRoom by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes chatRooms by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IConsultantDAO ...
type IConsultantDAO interface {
	// Create creates new consultant
	Create(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error)
	// Update updates consultant
	Update(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error)
	// Get gets consultant by ID
	Get(ctx context.Context, id string) (*dto.Consultant, error)
	// BatchGet gets consultants by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.Consultant, error)
	// Query queries consultants by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Consultant, error)
	// Delete deletes consultant by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes consultants by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IMeetingDAO ...
type IMeetingDAO interface {
	// Create creates new meeting
	Create(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error)
	// Update updates meeting
	Update(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error)
	// Get gets meeting by ID
	Get(ctx context.Context, id string) (*dto.Meeting, error)
	// BatchGet gets meetings by slice of IDs
	BatchGet(ctx context.Context, ids []string) ([]*dto.Meeting, error)
	// Query queries meetings by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meeting, error)
	// Delete deletes meeting by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes meetings by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}
