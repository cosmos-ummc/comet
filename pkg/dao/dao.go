package dao

import (
	"context"

	"comet/pkg/dto"
)

// IPatientDAO ...
type IPatientDAO interface {
	// Create creates new patient
	Create(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// Get gets patient by ID and type
	Get(ctx context.Context, id string, patientType int64) (*dto.Patient, error)
	// BatchGet gets patients by slice of IDs and type
	BatchGet(ctx context.Context, ids []string, patientType int64) ([]*dto.Patient, error)
	// Query queries patients by sort, range, type and a filter to match any number of fields
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Patient, error)
	// GetByStatus gets patients by statuses and type
	GetByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
	// GetSwabPatients gets swab patients of the specified type >= 14 days (ONLY RETURN PATIENTS WITH STATUS 1, 2, 3)
	GetSwabPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
	// GetDeclaredByTime gets declared patients of the specified type in given from timestamp
	GetDeclaredByTime(ctx context.Context, from int64, patientType int64) ([]*dto.Patient, error)
	// GetUndeclaredByTime gets undeclared patients of the specified type given from timestamp
	GetUndeclaredByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
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
	QueryNoCall(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
}

// ISwabDAO ...
type ISwabDAO interface {
	// Create creates new swab
	Create(ctx context.Context, swab *dto.Swab) (*dto.Swab, error)
	// Update updates swab
	Update(ctx context.Context, swab *dto.Swab) (*dto.Swab, error)
	// Get gets swab by ID
	Get(ctx context.Context, id string, patientType int64) (*dto.Swab, error)
	// BatchGet gets swabs by slice of IDs
	BatchGet(ctx context.Context, ids []string, patientType int64) ([]*dto.Swab, error)
	// Query queries swabs by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Swab, error)
	// Delete deletes swab by ID
	Delete(ctx context.Context, id string) error
	// BatchDelete deletes swabs by IDs
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
}

// IDeclarationDAO ...
type IDeclarationDAO interface {
	// Create creates new declaration
	Create(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// Get gets declaration
	Get(ctx context.Context, declarationID string, patientType int64) (*dto.Declaration, error)
	// Update updates declaration
	Update(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// Delete deletes declaration by ID
	Delete(ctx context.Context, declarationID string) error
	// Query queries declarations
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Declaration, error)
	// Query queries declarations by time
	QueryByTime(ctx context.Context, from int64, patientType int64) (int64, []*dto.Declaration, error)
	// QueryByCallingStatusAndTime queries declarations by time
	QueryByCallingStatusAndTime(ctx context.Context, callingStatus []int64, from int64, patientType int64) (int64, []*dto.Declaration, error)
	// BatchGet gets declarations
	BatchGet(ctx context.Context, ids []string, patientType int64) ([]*dto.Declaration, error)
	// BatchDelete deletes declarations
	BatchDelete(ctx context.Context, ids []string) ([]string, error)
	// QueryStableDeclarations query stable declarations
	QueryStableDeclarations(ctx context.Context, from int64, patientType int64) (int64, []*dto.Declaration, error)
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

type IPatientStatusReport interface {
	// Create creates status report
	Create(ctx context.Context, date string, patientStatusReport *dto.PatientStatusReport) error
	// UpdateByFields decrease decrementField by 1 and increase incrementField by 1 (based on patient type)
	UpdateByFields(ctx context.Context, date, decrementField, incrementField string, patientType int64) error
	// Get gets status report by date
	Get(ctx context.Context, date string) (*dto.PatientStatusReport, error)
	// BatchGet gets status report by dates
	BatchGet(ctx context.Context, dates []string) ([]*dto.PatientStatusReport, error)
}

type ICallingReport interface {
	// Create creates calling report
	Create(ctx context.Context, date string, callingReport *dto.CallingReport) error
	// UpdateByFields decrease decrementField by 1 and increase incrementField by 1 (based on patient type)
	UpdateByFields(ctx context.Context, date, decrementField, incrementField string, patientType int64) error
	// Get gets calling report by date
	Get(ctx context.Context, date string) (*dto.CallingReport, error)
	// BatchGet gets calling report by dates
	BatchGet(ctx context.Context, dates []string) ([]*dto.CallingReport, error)
}

type IDeclarationReport interface {
	// Create creates declaration report
	Create(ctx context.Context, date string, declarationReport *dto.DeclarationReport) error
	// UpdateByFields decrease decrementField by 1 for given date and increase incrementField by 1 (based on patient type)
	UpdateByFields(ctx context.Context, date, decrementField, incrementField string, patientType int64) error
	// Get gets declaration report by date
	Get(ctx context.Context, date string) (*dto.DeclarationReport, error)
	// BatchGet gets declaration report by dates
	BatchGet(ctx context.Context, dates []string) ([]*dto.DeclarationReport, error)
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

// IActivityDAO ...
type IActivityDAO interface {
	// Create creates new activity
	Create(ctx context.Context, activity *dto.Activity) (*dto.Activity, error)
	// Query queries activities by sort, range, filter
	Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Activity, error)
	// Get gets activity by ID
	Get(ctx context.Context, id string) (*dto.Activity, error)
}
