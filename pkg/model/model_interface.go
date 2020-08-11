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
	// QueryDeclarationsByPatientID ...
	QueryDeclarationsByPatientID(ctx context.Context, id string) (int64, []*dto.Declaration, error)
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

	///////////// Feed models
	// CreateFeed creates new feed
	CreateFeed(ctx context.Context, feed *dto.Feed) (*dto.Feed, error)
	// UpdateFeed updates feed
	UpdateFeed(ctx context.Context, feed *dto.Feed) (*dto.Feed, error)
	// UpdateFeeds update feeds
	UpdateFeeds(ctx context.Context, feed *dto.Feed, ids []string) ([]string, error)
	// GetFeed gets feed by ID
	GetFeed(ctx context.Context, id string) (*dto.Feed, error)
	// BatchGetFeeds get feeds by slice of IDs
	BatchGetFeeds(ctx context.Context, ids []string) ([]*dto.Feed, error)
	// QueryFeeds queries feeds by sort, range, filter
	QueryFeeds(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Feed, error)
	// DeleteFeed deletes feed by ID
	DeleteFeed(ctx context.Context, id string) (*dto.Feed, error)
	// DeleteFeeds delete feeds by IDs
	DeleteFeeds(ctx context.Context, ids []string) ([]string, error)
	/////////////

	///////////// Game models
	// CreateGame creates new game
	CreateGame(ctx context.Context, game *dto.Game) (*dto.Game, error)
	// UpdateGame updates game
	UpdateGame(ctx context.Context, game *dto.Game) (*dto.Game, error)
	// UpdateGames update games
	UpdateGames(ctx context.Context, game *dto.Game, ids []string) ([]string, error)
	// GetGame gets game by ID
	GetGame(ctx context.Context, id string) (*dto.Game, error)
	// BatchGetGames get games by slice of IDs
	BatchGetGames(ctx context.Context, ids []string) ([]*dto.Game, error)
	// QueryGames queries games by sort, range, filter
	QueryGames(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Game, error)
	// DeleteGame deletes game by ID
	DeleteGame(ctx context.Context, id string) (*dto.Game, error)
	// DeleteGames delete games by IDs
	DeleteGames(ctx context.Context, ids []string) ([]string, error)
	/////////////

	///////////// Meditation models
	// CreateMeditation creates new meditation
	CreateMeditation(ctx context.Context, meditation *dto.Meditation) (*dto.Meditation, error)
	// UpdateMeditation updates meditation
	UpdateMeditation(ctx context.Context, meditation *dto.Meditation) (*dto.Meditation, error)
	// UpdateMeditations update meditations
	UpdateMeditations(ctx context.Context, meditation *dto.Meditation, ids []string) ([]string, error)
	// GetMeditation gets meditation by ID
	GetMeditation(ctx context.Context, id string) (*dto.Meditation, error)
	// BatchGetMeditations get meditations by slice of IDs
	BatchGetMeditations(ctx context.Context, ids []string) ([]*dto.Meditation, error)
	// QueryMeditations queries meditations by sort, range, filter
	QueryMeditations(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meditation, error)
	// DeleteMeditation deletes meditation by ID
	DeleteMeditation(ctx context.Context, id string) (*dto.Meditation, error)
	// DeleteMeditations delete meditations by IDs
	DeleteMeditations(ctx context.Context, ids []string) ([]string, error)
	/////////////

	///////////// ChatMessage models
	// CreateChatMessage creates new chatMessage
	CreateChatMessage(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error)
	// UpdateChatMessage updates chatMessage
	UpdateChatMessage(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error)
	// UpdateChatMessages update chatMessages
	UpdateChatMessages(ctx context.Context, chatMessage *dto.ChatMessage, ids []string) ([]string, error)
	// GetChatMessage gets chatMessage by ID
	GetChatMessage(ctx context.Context, id string) (*dto.ChatMessage, error)
	// BatchGetChatMessages get chatMessages by slice of IDs
	BatchGetChatMessages(ctx context.Context, ids []string) ([]*dto.ChatMessage, error)
	// QueryChatMessages queries chatMessages by sort, range, filter
	QueryChatMessages(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatMessage, error)
	// DeleteChatMessage deletes chatMessage by ID
	DeleteChatMessage(ctx context.Context, id string) (*dto.ChatMessage, error)
	// DeleteChatMessages delete chatMessages by IDs
	DeleteChatMessages(ctx context.Context, ids []string) ([]string, error)
	/////////////

	///////////// ChatRoom models
	// CreateChatRoom creates new chatRoom
	CreateChatRoom(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error)
	// UpdateChatRoom updates chatRoom
	UpdateChatRoom(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error)
	// UpdateChatRooms update chatRooms
	UpdateChatRooms(ctx context.Context, chatRoom *dto.ChatRoom, ids []string) ([]string, error)
	// GetChatRoom gets chatRoom by ID
	GetChatRoom(ctx context.Context, id string) (*dto.ChatRoom, error)
	// BatchGetChatRooms get chatRooms by slice of IDs
	BatchGetChatRooms(ctx context.Context, ids []string) ([]*dto.ChatRoom, error)
	// QueryChatRooms queries chatRooms by sort, range, filter
	QueryChatRooms(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatRoom, error)
	// DeleteChatRoom deletes chatRoom by ID
	DeleteChatRoom(ctx context.Context, id string) (*dto.ChatRoom, error)
	// DeleteChatRooms delete chatRooms by IDs
	DeleteChatRooms(ctx context.Context, ids []string) ([]string, error)
	/////////////

	///////////// Consultant models
	// CreateConsultant creates new consultant
	CreateConsultant(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error)
	// UpdateConsultant updates consultant
	UpdateConsultant(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error)
	// UpdateConsultants update consultants
	UpdateConsultants(ctx context.Context, consultant *dto.Consultant, ids []string) ([]string, error)
	// GetConsultant gets consultant by ID
	GetConsultant(ctx context.Context, id string) (*dto.Consultant, error)
	// BatchGetConsultants get consultants by slice of IDs
	BatchGetConsultants(ctx context.Context, ids []string) ([]*dto.Consultant, error)
	// QueryConsultants queries consultants by sort, range, filter
	QueryConsultants(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Consultant, error)
	// DeleteConsultant deletes consultant by ID
	DeleteConsultant(ctx context.Context, id string) (*dto.Consultant, error)
	// DeleteConsultants delete consultants by IDs
	DeleteConsultants(ctx context.Context, ids []string) ([]string, error)
	/////////////

	///////////// Meeting models
	// CreateMeeting creates new meeting
	CreateMeeting(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error)
	// UpdateMeeting updates meeting
	UpdateMeeting(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error)
	// UpdateMeetings update meetings
	UpdateMeetings(ctx context.Context, meeting *dto.Meeting, ids []string) ([]string, error)
	// GetMeeting gets meeting by ID
	GetMeeting(ctx context.Context, id string) (*dto.Meeting, error)
	// BatchGetMeetings get meetings by slice of IDs
	BatchGetMeetings(ctx context.Context, ids []string) ([]*dto.Meeting, error)
	// QueryMeetings queries meetings by sort, range, filter
	QueryMeetings(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meeting, error)
	// DeleteMeeting deletes meeting by ID
	DeleteMeeting(ctx context.Context, id string) (*dto.Meeting, error)
	// DeleteMeetings delete meetings by IDs
	DeleteMeetings(ctx context.Context, ids []string) ([]string, error)
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
	QueryUsers(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter *dto.FilterData, superuserOnly bool) (int64, []*dto.User, error)
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
