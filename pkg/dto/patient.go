package dto

// Patient ...
type Patient struct {
	ID                 string `json:"id" bson:"id"`
	UserID             string `json:"userId" bson:"userId"`
	Name               string `json:"name" bson:"name"`
	TelegramID         string `json:"telegramId" bson:"telegramId"`
	PhoneNumber        string `json:"phoneNumber" bson:"phoneNumber"`
	Email              string `json:"email" bson:"email"`
	StressStatus       int64  `json:"stressStatus" bson:"stressStatus"`
	AnxietyStatus      int64  `json:"anxietyStatus" bson:"anxietyStatus"`
	DepressionStatus   int64  `json:"depressionStatus" bson:"depressionStatus"`
	PtsdStatus         int64  `json:"ptsdStatus" bson:"ptsdStatus"`
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
	SwabDate           string `json:"swabDate" bson:"swabDate"`
	SwabResult         int64  `json:"swabResult" bson:"swabResult"`
}
