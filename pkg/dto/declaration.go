package dto

// Declaration ...
type Declaration struct {
	ID                 string `json:"id" bson:"id"`
	PatientID          string `json:"patientId" bson:"patientId"`
	PatientName        string `json:"patientName" bson:"patientName"`
	PatientType        int64  `json:"-" bson:"patientType"`
	Cough              int64  `json:"cough" bson:"cough"`
	Throat             int64  `json:"throat" bson:"throat"`
	Fever              int64  `json:"fever" bson:"fever"`
	Breathe            int64  `json:"breathe" bson:"breathe"`
	Chest              int64  `json:"chest" bson:"chest"`
	Blue               int64  `json:"blue" bson:"blue"`
	Drowsy             int64  `json:"drowsy" bson:"drowsy"`
	HasSymptom         bool   `json:"hasSymptom" bson:"hasSymptom"`
	SubmittedAt        int64  `json:"submittedAt" bson:"submittedAt"`
	CallingStatus      int64  `json:"callingStatus" bson:"callingStatus"`
	Date               string `json:"date" bson:"date"`
	PatientPhoneNumber string `json:"patientPhoneNumber" bson:"patientPhoneNumber"`
	DoctorRemarks      string `json:"doctorRemarks" bson:"doctorRemarks"`
	Loss               int64  `json:"loss" bson:"loss"`
}
