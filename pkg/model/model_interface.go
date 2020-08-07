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
	CreateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// GetDeclaration gets declaration
	GetDeclaration(ctx context.Context, declarationID string) (*dto.Declaration, error)
	// UpdateDeclaration updates declaration
	UpdateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error)
	// DeleteDeclaration deletes declaration
	DeleteDeclaration(ctx context.Context, declarationID string) (*dto.Declaration, error)
	// BatchGetDeclarations get declarations
	BatchGetDeclarations(ctx context.Context, declarationID []string) ([]*dto.Declaration, error)
	// UpdateDeclarations update declarations
	UpdateDeclarations(ctx context.Context, declaration *dto.Declaration, ids []string) ([]string, error)
	// DeleteDeclarations delete declarations
	DeleteDeclarations(ctx context.Context, declarationID []string) ([]string, error)
	// QueryDeclarations query declarations
	QueryDeclarations(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Declaration, error)
	// QueryDeclarationsByTime ...
	QueryDeclarationsByTime(ctx context.Context, from int64) (int64, []*dto.Declaration, error)
	// QueryDeclarationsByCallingStatusAndTime ...
	QueryDeclarationsByCallingStatusAndTime(ctx context.Context, callingStatus []int64, from int64) (int64, []*dto.Declaration, error)
	// QueryDeclarationsByPatientID ...
	QueryDeclarationsByPatientID(ctx context.Context, id string) (int64, []*dto.Declaration, error)
	// GetStableDeclarations ...
	GetStableDeclarations(ctx context.Context, from int64) (int64, []*dto.Declaration, error)
	/////////////

	///////////// Patient models
	// CreatePatient creates new patient
	CreatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// GetPatient gets patient of specified type by ID
	GetPatient(ctx context.Context, id string) (*dto.Patient, error)
	// QueryPatients queries patients of specified type by sort, range, filter
	QueryPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Patient, error)
	// BatchGetPatients get patients of specified type by slice of IDs
	BatchGetPatients(ctx context.Context, ids []string) ([]*dto.Patient, error)
	// GetPatientsByStatus gets patients of specified type by list of status
	GetPatientsByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error)
	// GetNoCallPatients queries patients who have declared but no yet call
	GetNoCallPatients(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error)
	// GetDeclaredByTime gets declared patients of specified type in given from timestamp
	GetDeclaredPatientsByTime(ctx context.Context, from int64) ([]*dto.Patient, error)
	// GetUndeclaredPatientsByTime gets undeclared patients of specified type given from timestamp
	GetUndeclaredPatientsByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error)
	// UpdatePatient updates patient
	UpdatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// UpdatePatients update patients
	UpdatePatients(ctx context.Context, patient *dto.Patient, ids []string) ([]string, error)
	// DeletePatient deletes patient by ID
	DeletePatient(ctx context.Context, id string) (*dto.Patient, error)
	// DeletePatients delete patients by IDs
	DeletePatients(ctx context.Context, ids []string) ([]string, error)
	// ClientGetPatientsByConsentTime gets patients given from and to timestamp
	ClientGetPatientsByConsentTime(ctx context.Context, from int64, to int64) ([]*dto.Patient, error)
	// ClientUpdatePatient updates patient from chatbot client
	ClientUpdatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error)
	// ClientGetUndeclaredPatientsByTime gets undeclared patients (with telegramID) given from timestamp
	ClientGetUndeclaredPatientsByTime(ctx context.Context, from int64) ([]*dto.Patient, error)
	/////////////

	///////////// Question models
	// CreateQuestion creates new question
	CreateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error)
	// UpdateQuestion updates question
	UpdateQuestion(ctx context.Context, question *dto.Question) (*dto.Question, error)
	// UpdateQuestions update questions
	UpdateQuestions(ctx context.Context, question *dto.Question, ids []string) ([]string, error)
	// GetQuestion gets question by ID
	GetQuestion(ctx context.Context, id string) (*dto.Question, error)
	// BatchGetQuestions get questions by slice of IDs
	BatchGetQuestions(ctx context.Context, ids []string) ([]*dto.Question, error)
	// QueryQuestions queries questions by sort, range, filter
	QueryQuestions(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Question, error)
	// DeleteQuestion deletes question by ID
	DeleteQuestion(ctx context.Context, id string) (*dto.Question, error)
	// DeleteQuestions delete questions by IDs
	DeleteQuestions(ctx context.Context, ids []string) ([]string, error)
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
	// GetReport gets report
	GetReport(ctx context.Context, dateString string) (*dto.Report, error)
	// GetReports get reports given dates inclusive
	GetReports(ctx context.Context, from, to string) ([]*dto.Report, error)
	// GenerateReport generates all reports based on latest data
	GenerateReport(ctx context.Context, date string) error
	// SyncDays syncs days
	SyncDays(ctx context.Context) error
	// SyncPatientReport sync patient record when changed type
	SyncPatientReport(ctx context.Context) error
	/////////////
}
