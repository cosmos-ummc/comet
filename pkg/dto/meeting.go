package dto

// Meeting ...
type Meeting struct {
	ID                    string `json:"id" bson:"id"`
	PatientID             string `json:"patientId" bson:"patientId"`
	PatientName           string `json:"patientName" bson:"patientName"`
	PatientPhoneNumber    string `json:"patientPhoneNumber" bson:"patientPhoneNumber"`
	ConsultantID          string `json:"consultantId" bson:"consultantId"`
	ConsultantName        string `json:"consultantName" bson:"consultantName"`
	ConsultantPhoneNumber string `json:"consultantPhoneNumber" bson:"consultantPhoneNumber"`
	Time                  string `json:"time" bson:"time"`
	Status                int64  `json:"status" bson:"status"`
}
