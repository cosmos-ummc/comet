package constants

// Database
const (
	Cosmos = "cosmos"
)

// Collections
const (
	Declarations         = "declarations"
	Patients             = "patients"
	Activities           = "activities"
	Swabs                = "swabs"
	Users                = "users"
	PatientStatusReports = "patientstatusreports"
	CallingReports       = "callingreports"
	DeclarationReports   = "declarationreports"
	AuthTokens           = "authtokens"
)

// Documents
const (
	Total = "total"
)

// Fields
const (
	// Token
	Token   = "token"
	TTL     = "ttl"
	Access  = "access"
	Refresh = "refresh"
	Reset   = "reset"
	// Common
	ID                 = "id"
	PatientID          = "patientId"
	PatientName        = "patientName"
	PatientPhoneNumber = "patientPhoneNumber"
	PatientType        = "patientType"
	Remarks            = "remarks"
	PhoneNumber        = "phoneNumber"
	Status             = "status"

	// Patients
	Name              = "name"
	TelegramID        = "telegramId"
	DaysSinceExposure = "daysSinceExposure"
	LastDeclared      = "lastDeclared"
	SwabCount         = "swabCount"
	Episode           = "episode"
	Type              = "type"
	LastDeclareResult = "lastDeclareResult"
	ExposureDate      = "exposureDate"
	RegistrationNum   = "registrationNum"
	AlternateContact  = "alternateContact"
	IsolationAddress  = "isolationAddress"
	SymptomDate       = "symptomDate"
	SwabDate          = "swabDate"
	FeverContDay      = "feverContDay"
	Localization      = "localization"
	ExposureSource    = "exposureSource"
	Consent           = "consent"
	PrivacyPolicy     = "privacyPolicy"
	FeverStartDate    = "feverStartDate"
	DaysSinceSwab     = "daysSinceSwab"

	// Users
	DisplayName = "displayName"
	Email       = "email"
	Role        = "role"
	Disabled    = "disabled"
	Password    = "password"

	// Activities
	UserName   = "userName"
	OldPatient = "oldPatient"
	NewPatient = "newPatient"
	OldSwab    = "oldSwab"
	NewSwab    = "newSwab"

	// Swabs
	Count    = "count"
	Date     = "date"
	Location = "location"

	// Declarations
	HasSymptom    = "hasSymptom"
	SubmittedAt   = "submittedAt"
	Cough         = "cough"
	Throat        = "throat"
	Fever         = "fever"
	Breathe       = "breathe"
	Chest         = "chest"
	Blue          = "blue"
	Drowsy        = "drowsy"
	CallingStatus = "callingStatus"
	DoctorRemarks = "doctorRemarks"

	Authorized  = "authorized"
	AccessUuid  = "accessUuid"
	UserId      = "userId"
	Exp         = "exp"
	RefreshUuid = "refreshUuid"
)

// Keywords
const (
	ASC  = "ASC"
	DESC = "DESC"
)
