package dto

// Declaration ...
type Declaration struct {
	ID            string      `json:"id" bson:"id"`
	PatientID     string      `json:"patientId" bson:"patientId"`
	Result        []*Question `json:"result" bson:"result"`
	Category      string      `json:"category" bson:"category"`
	Score         int64       `json:"score" bson:"score"`
	Depression    int64       `json:"depression" bson:"depression"`
	Anxiety       int64       `json:"anxiety" bson:"anxiety"`
	Stress        int64       `json:"stress" bson:"stress"`
	Status        int64       `json:"status" bson:"status"`
	SubmittedAt   int64       `json:"submittedAt" bson:"submittedAt"`
	DoctorRemarks string      `json:"doctorRemarks" bson:"doctorRemarks"`
}
