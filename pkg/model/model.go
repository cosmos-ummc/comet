package model

import (
	"comet/pkg/dao"

	"go.mongodb.org/mongo-driver/mongo"
)

// Model ...
type Model struct {
	declarationDAO dao.IDeclarationDAO
	patientDAO     dao.IPatientDAO
	questionDAO    dao.IQuestionDAO
	userDAO        dao.IUserDAO
	reportDAO      dao.IReportDAO
	authDAO        dao.IAuthDAO
	chatMessageDAO dao.IChatMessageDAO
	chatRoomDAO    dao.IChatRoomDAO
	consultantDAO  dao.IConsultantDAO
	meetingDAO     dao.IMeetingDAO
	feedDAO        dao.IFeedDAO
}

// InitModel ...
func InitModel(client *mongo.Client) IModel {
	return &Model{
		declarationDAO: dao.InitDeclarationDAO(client),
		patientDAO:     dao.InitPatientDAO(client),
		questionDAO:    dao.InitQuestionDAO(client),
		userDAO:        dao.InitUserDAO(client),
		reportDAO:      dao.InitReportDAO(client),
		authDAO:        dao.InitAuthDAO(client),
		chatMessageDAO: dao.InitChatMessageDAO(client),
		chatRoomDAO:    dao.InitChatRoomDAO(client),
		consultantDAO:  dao.InitConsultantDAO(client),
		meetingDAO:     dao.InitMeetingDAO(client),
		feedDAO:        dao.InitFeedDAO(client),
	}
}
