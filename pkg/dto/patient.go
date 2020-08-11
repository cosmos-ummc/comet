package dto

// Patient ...
type Patient struct {
	ID                 string `json:"id" bson:"id"`
	UserID             string `json:"userId" bson:"userId"`
	Name               string `json:"name" bson:"name"`
	TelegramID         string `json:"telegramId" bson:"telegramId"`
	PhoneNumber        string `json:"phoneNumber" bson:"phoneNumber"`
	Email              string `json:"email" bson:"email"`
	Status             int64  `json:"status" bson:"status"`
	LastDassTime       int64  `json:"lastDassTime" bson:"lastDassTime"`
	LastIesrTime       int64  `json:"lastIesrTime" bson:"lastIesrTime"`
	LastDassResult     int64  `json:"lastDassResult" bson:"lastDassResult"`
	LastIesrResult     int64  `json:"lastIesrResult" bson:"lastIesrResult"`
	Remarks            string `json:"remarks" bson:"remarks"`
	Consent            int64  `json:"consent" bson:"consent"`
	PrivacyPolicy      int64  `json:"privacyPolicy" bson:"privacyPolicy"`
	HomeAddress        string `json:"homeAddress" bson:"homeAddress"`
	IsolationAddress   string `json:"isolationAddress" bson:"isolationAddress"`
	RegistrationStatus int64  `json:"registrationStatus" bson:"registrationStatus"`
	DaySinceMonitoring int64  `json:"daySinceMonitoring" bson:"daySinceMonitoring"`
	HasCompleted       bool   `json:"hasCompleted" bson:"hasCompleted"`
	MentalStatus       int64  `json:"mentalStatus" bson:"mentalStatus"`
	Type               int64  `json:"type" bson:"type"`
}
