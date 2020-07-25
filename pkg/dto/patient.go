package dto

// Patient ...
type Patient struct {
	ID                 string `json:"id" bson:"id"`
	Name               string `json:"name" bson:"name"`
	TelegramID         string `json:"telegramId" bson:"telegramId"`
	PhoneNumber        string `json:"phoneNumber" bson:"phoneNumber"`
	Email              string `json:"email" bson:"email"`
	Status             int64  `json:"status" bson:"status"`
	LastDeclared       int64  `json:"lastDeclared" bson:"lastDeclared"`
	SwabCount          int64  `json:"swabCount" bson:"swabCount"`
	Episode            int64  `json:"episode" bson:"episode"`
	Type               int64  `json:"type" bson:"type"`
	TypeChangeDate     string `json:"typeChangeDate" bson:"typeChangeDate"`
	LastDeclareResult  bool   `json:"lastDeclareResult" bson:"lastDeclareResult"`
	ExposureDate       string `json:"exposureDate" bson:"exposureDate"`
	ExposureSource     string `json:"exposureSource" bson:"exposureSource"`
	DaysSinceExposure  int64  `json:"daysSinceExposure" bson:"daysSinceExposure"`
	RegistrationNum    string `json:"registrationNum" bson:"registrationNum"`
	AlternateContact   string `json:"alternateContact" bson:"alternateContact"`
	IsolationAddress   string `json:"isolationAddress" bson:"isolationAddress"`
	SymptomDate        string `json:"symptomDate" bson:"symptomDate"`
	SwabDate           string `json:"swabDate" bson:"swabDate"`
	FeverContDay       int64  `json:"feverContDay" bson:"feverContDay"`
	Remarks            string `json:"remarks" bson:"remarks"`
	Localization       int64  `json:"localization" bson:"localization"`
	Consent            int64  `json:"consent" bson:"consent"`
	PrivacyPolicy      int64  `json:"privacyPolicy" bson:"privacyPolicy"`
	CallingStatus      int64  `json:"callingStatus" bson:"-"`
	FeverStartDate     string `json:"feverStartDate" bson:"feverStartDate"`
	DaysSinceSwab      int64  `json:"daysSinceSwab" bson:"daysSinceSwab"`
	HomeAddress        string `json:"homeAddress" bson:"homeAddress"`
	IsSameAddress      bool   `json:"isSameAddress" bson:"isSameAddress"`
	RegistrationStatus int64  `json:"registrationStatus" bson:"-"`
}
