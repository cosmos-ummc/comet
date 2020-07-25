package model

import (
	"comet/pkg/dao"

	"go.mongodb.org/mongo-driver/mongo"
)

// Model ...
type Model struct {
	declarationDAO         dao.IDeclarationDAO
	patientDAO             dao.IPatientDAO
	swabDAO                dao.ISwabDAO
	userDAO                dao.IUserDAO
	declarationReportDAO   dao.IDeclarationReport
	callingReportDAO       dao.ICallingReport
	patientStatusReportDAO dao.IPatientStatusReport
	authDAO                dao.IAuthDAO
	activityDAO            dao.IActivityDAO
}

// InitModel ...
func InitModel(client *mongo.Client) IModel {
	return &Model{
		declarationDAO:         dao.InitDeclarationDAO(client),
		patientDAO:             dao.InitPatientDAO(client),
		swabDAO:                dao.InitSwabDAO(client),
		userDAO:                dao.InitUserDAO(client),
		declarationReportDAO:   dao.InitDeclarationReportDAO(client),
		callingReportDAO:       dao.InitCallingReportDAO(client),
		patientStatusReportDAO: dao.InitPatientStatusReportDAO(client),
		authDAO:                dao.InitAuthDAO(client),
		activityDAO:            dao.InitActivityDAO(client),
	}
}
