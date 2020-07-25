package model

import (
	"context"

	"comet/pkg/dto"
)

// IModel ...
type IModel interface {
	///////////// Declaration models
	// ClientCreateDeclaration creates new declaration or updates existing declaration
	ClientCreateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// CreateDeclaration creates new declaration
	CreateDeclaration(ctx context.Context, declaration *dto.Declaration, patientType int64) (*dto.Declaration, error)
	// GetDeclaration gets declaration
	GetDeclaration(ctx context.Context, declarationID string, patientType int64) (*dto.Declaration, error)
	// UpdateDeclaration updates declaration
	UpdateDeclaration(ctx context.Context, declaration *dto.Declaration, patientType int64) (*dto.Declaration, error)
	// DeleteDeclaration deletes declaration
	DeleteDeclaration(ctx context.Context, declarationID string, patientType int64) (*dto.Declaration, error)
	// BatchGetDeclarations get declarations
	BatchGetDeclarations(ctx context.Context, declarationID []string, patientType int64) ([]*dto.Declaration, error)
	// UpdateDeclarations update declarations
	UpdateDeclarations(ctx context.Context, declaration *dto.Declaration, ids []string, patientType int64) ([]string, error)
	// DeleteDeclarations delete declarations
	DeleteDeclarations(ctx context.Context, declarationID []string, patientType int64) ([]string, error)
	// QueryDeclarations query declarations
	QueryDeclarations(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Declaration, error)
	// QueryDeclarationsByTime ...
	QueryDeclarationsByTime(ctx context.Context, from int64, patientType int64) (int64, []*dto.Declaration, error)
	// QueryDeclarationsByCallingStatusAndTime ...
	QueryDeclarationsByCallingStatusAndTime(ctx context.Context, callingStatus []int64, from int64, patientType int64) (int64, []*dto.Declaration, error)
	// QueryDeclarationsByPatientID ...
	QueryDeclarationsByPatientID(ctx context.Context, id string, patientType int64) (int64, []*dto.Declaration, error)
	// GetStableDeclarations ...
	GetStableDeclarations(ctx context.Context, from int64, patientType int64) (int64, []*dto.Declaration, error)
	/////////////

	///////////// Patient models
	// CreatePatient creates new patient
	CreatePatient(ctx context.Context, patient *dto.Patient, user *dto.User) (*dto.Patient, error)
	// GetPatient gets patient of specified type by ID
	GetPatient(ctx context.Context, id string, patientType int64) (*dto.Patient, error)
	// QueryPatients queries patients of specified type by sort, range, filter
	QueryPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Patient, error)
	// BatchGetPatients get patients of specified type by slice of IDs
	BatchGetPatients(ctx context.Context, ids []string, patientType int64) ([]*dto.Patient, error)
	// GetPatientsByStatus gets patients of specified type by list of status
	GetPatientsByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
	// GetSwabPatients get patients of specified type days since swab >= 14
	GetSwabPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
	// GetNoCallPatients queries patients who have declared but no yet call
	GetNoCallPatients(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
	// GetDeclaredByTime gets declared patients of specified type in given from timestamp
	GetDeclaredPatientsByTime(ctx context.Context, from int64, patientType int64) ([]*dto.Patient, error)
	// GetUndeclaredPatientsByTime gets undeclared patients of specified type given from timestamp
	GetUndeclaredPatientsByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error)
	// UpdatePatient updates patient
	UpdatePatient(ctx context.Context, patient *dto.Patient, patientType int64, user *dto.User) (*dto.Patient, error)
	// UpdatePatients update patients
	UpdatePatients(ctx context.Context, patient *dto.Patient, ids []string, patientType int64, user *dto.User) ([]string, error)
	// DeletePatient deletes patient by ID
	DeletePatient(ctx context.Context, id string, patientType int64) (*dto.Patient, error)
	// DeletePatients delete patients by IDs
	DeletePatients(ctx context.Context, ids []string, patientType int64) ([]string, error)
	// ClientGetPatientsByConsentTime gets patients given from and to timestamp
	ClientGetPatientsByConsentTime(ctx context.Context, from int64, to int64) ([]*dto.Patient, error)
	// ClientUpdatePatient updates patient from chatbot client
	ClientUpdatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// ClientGetUndeclaredPatientsByTime gets undeclared patients (with telegramID) given from timestamp
	ClientGetUndeclaredPatientsByTime(ctx context.Context, from int64) ([]*dto.Patient, error)
	/////////////

	///////////// Swab models
	// CreateSwab creates new swab
	CreateSwab(ctx context.Context, swab *dto.Swab, patientType int64, user *dto.User) (*dto.Swab, error)
	// UpdateSwab updates swab
	UpdateSwab(ctx context.Context, swab *dto.Swab, patientType int64, user *dto.User) (*dto.Swab, error)
	// UpdateSwabs update swabs
	UpdateSwabs(ctx context.Context, swab *dto.Swab, ids []string, patientType int64, user *dto.User) ([]string, error)
	// GetSwab gets swab by ID
	GetSwab(ctx context.Context, id string, patientType int64) (*dto.Swab, error)
	// BatchGetSwabs get swabs by slice of IDs
	BatchGetSwabs(ctx context.Context, ids []string, patientType int64) ([]*dto.Swab, error)
	// QuerySwabs queries swabs by sort, range, filter
	QuerySwabs(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Swab, error)
	// DeleteSwab deletes swab by ID
	DeleteSwab(ctx context.Context, id string, patientType int64) (*dto.Swab, error)
	// DeleteSwabs delete swabs by IDs
	DeleteSwabs(ctx context.Context, ids []string, patientType int64) ([]string, error)
	// QuerySwabsByPatientID ...
	QuerySwabsByPatientID(ctx context.Context, id string, patientType int64) ([]*dto.Swab, error)

	/////////////

	///////////// User models
	// CreateUser creates new user
	CreateUser(ctx context.Context, user *dto.User) (*dto.User, error)
	// UpdateUser updates user
	UpdateUser(ctx context.Context, user *dto.User) (*dto.User, error)
	// UpdateUserPassword updates user password only
	UpdateUserPassword(ctx context.Context, user *dto.User) (*dto.User, error)
	// CreateToken creates token with custom ttl
	CreateToken(ctx context.Context, auth *dto.AuthObject) (*dto.AuthObject, error)
	// RevokeTokensByUserID revokes all tokens by UserID
	RevokeTokensByUserID(ctx context.Context, id string) error
	// GetUserIDByToken gets userID by token
	GetUserIDByToken(ctx context.Context, token string) (string, error)
	// UpdateUsers update users
	UpdateUsers(ctx context.Context, user *dto.User, ids []string) ([]string, error)
	// GetUser gets user by ID
	GetUser(ctx context.Context, id string) (*dto.User, error)
	// BatchGetUsers get users by slice of IDs
	BatchGetUsers(ctx context.Context, ids []string) ([]*dto.User, error)
	// QueryUsers queries users by sort, range, filter
	QueryUsers(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter *dto.FilterData) (int64, []*dto.User, error)
	// DeleteUser deletes user by ID
	DeleteUser(ctx context.Context, id string) (*dto.User, error)
	// RevokeUserTokens revoke all user tokens
	RevokeUserTokens(ctx context.Context) error
	// DeleteUsers delete users by IDs
	DeleteUsers(ctx context.Context, ids []string) ([]string, error)
	// Login verifies user by email and password and return tokens
	Login(ctx context.Context, email string, password string) (*dto.User, error)
	// VerifyUser verifies user by header
	VerifyUser(ctx context.Context, header string) (*dto.User, error)
	// Logout logs user out from the system by header
	Logout(ctx context.Context, header string) error
	// Refresh returns new token to authorized user by header
	Refresh(ctx context.Context, header string) (*dto.User, error)
	/////////////

	///////////// Report models
	// GetDeclarationReport gets declaration report
	GetDeclarationReport(ctx context.Context, dateString string, patientType int64) (*dto.DeclarationReport, error)
	// GetDeclarationReports get declaration reports given dates inclusive
	GetDeclarationReports(ctx context.Context, from, to string, patientType int64) ([]*dto.DeclarationReport, error)
	// GetCallingReport gets calling report
	GetCallingReport(ctx context.Context, date string, patientType int64) (*dto.CallingReport, error)
	// GetCallingReports get calling reports given dates inclusive
	GetCallingReports(ctx context.Context, from, to string, patientType int64) ([]*dto.CallingReport, error)
	// GetPatientStatusReport gets patient status report
	GetPatientStatusReport(ctx context.Context, dateString string, patientType int64) (*dto.PatientStatusReport, error)
	// GetPatientStatusReports get patient status reports given dates inclusive
	GetPatientStatusReports(ctx context.Context, from, to string, patientType int64) ([]*dto.PatientStatusReport, error)
	// GenerateReport generates all reports based on latest data
	GenerateReport(ctx context.Context, date string) error
	// SyncDays syncs days
	SyncDays(ctx context.Context) error
	// SyncPatientReport sync patient record when changed type
	SyncPatientReport(ctx context.Context) error
	/////////////

	///////////// Activity models
	// CreateActivity creates new activity
	CreateActivity(ctx context.Context, activity *dto.Activity) (*dto.Activity, error)
	// GetActivity gets activity by ID
	GetActivity(ctx context.Context, id string) (*dto.Activity, error)
	// QueryActivities queries activities by sort, range, filter
	QueryActivities(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Activity, error)
	/////////////
}
