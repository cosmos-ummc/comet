package dto

// Meeting ...
type Meeting struct {
	ID           string `json:"id" bson:"id"`
	PatientID    string `json:"patientId" bson:"patientId"`
	ConsultantID string `json:"consultantId" bson:"consultantId"`
	Time         int64  `json:"time" bson:"time"`
	Status       int64  `json:"status" bson:"status"`
}
