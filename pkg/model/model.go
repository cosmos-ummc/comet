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
	authDAO        dao.IAuthDAO
	chatMessageDAO dao.IChatMessageDAO
	chatRoomDAO    dao.IChatRoomDAO
	consultantDAO  dao.IConsultantDAO
	meetingDAO     dao.IMeetingDAO
	feedDAO        dao.IFeedDAO
	gameDAO        dao.IGameDAO
	meditationDAO  dao.IMeditationDAO
}

// InitModel ...
func InitModel(client *mongo.Client) IModel {
	return &Model{
		declarationDAO: dao.InitDeclarationDAO(client),
		patientDAO:     dao.InitPatientDAO(client),
		questionDAO:    dao.InitQuestionDAO(client),
		userDAO:        dao.InitUserDAO(client),
		authDAO:        dao.InitAuthDAO(client),
		chatMessageDAO: dao.InitChatMessageDAO(client),
		chatRoomDAO:    dao.InitChatRoomDAO(client),
		consultantDAO:  dao.InitConsultantDAO(client),
		meetingDAO:     dao.InitMeetingDAO(client),
		feedDAO:        dao.InitFeedDAO(client),
		gameDAO:        dao.InitGameDAO(client),
		meditationDAO:  dao.InitMeditationDAO(client),
	}
}
