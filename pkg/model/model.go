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
	}
}
