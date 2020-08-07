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
	// GetDeclaredByTime gets declared patients of the specified type in given from timestamp
	GetDeclaredByTime(ctx context.Context, from int64) ([]*dto.Patient, error)
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
	// QueryNoCall queries patients who have declared but no yet call
	QueryNoCall(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error)
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
	// Query queries declarations by time
	QueryByTime(ctx context.Context, from int64) (int64, []*dto.Declaration, error)
	// QueryByCallingStatusAndTime queries declarations by time
	QueryByCallingStatusAndTime(ctx context.Context, callingStatus []int64, from int64) (int64, []*dto.Declaration, error)
	// BatchGet gets declarations
	BatchGet(ctx context.Context, ids []string) ([]*dto.Declaration, error)
	// BatchDelete deletes declarations
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
	// QueryStableDeclarations query stable declarations
	QueryStableDeclarations(ctx context.Context, from int64) (int64, []*dto.Declaration, error)
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
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter *dto.FilterData) (int64, []*dto.User, error)
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
