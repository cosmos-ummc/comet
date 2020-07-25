package dto

// Swab ...
type Swab struct {
	ID                  string `json:"id" bson:"id"`
	PatientID           string `json:"patientId" bson:"patientId"`
	PatientName         string `json:"patientName" bson:"patientName"`
	PatientPhoneNumber  string `json:"patientPhoneNumber" bson:"patientPhoneNumber"`
	PatientType         int64  `json:"-" bson:"patientType"`
	Status              int64  `json:"status" bson:"status"`
	Date                string `json:"date" bson:"date"`
	Location            string `json:"location" bson:"location"`
	IsOtherSwabLocation bool   `json:"isOtherSwabLocation" bson:"isOtherSwabLocation"`
}
