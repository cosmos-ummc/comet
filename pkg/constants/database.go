package constants

// Database
const (
	Mhpss = "mhpss"
)

// Collections
const (
	Declarations = "declarations"
	Patients     = "patients"
	Users        = "users"
	AuthTokens   = "authtokens"
	Reports      = "reports"
	Questions    = "questions"
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
	LastDeclared      = "lastDeclared"
	LastDeclareResult = "lastDeclareResult"
	Consent           = "consent"
	PrivacyPolicy     = "privacyPolicy"

	// Users
	Email    = "email"
	Role     = "role"
	Disabled = "disabled"
	Password = "password"

	Date = "date"

	// Questions
	Contents = "contents"
	Category = "category"

	// Declarations
	SubmittedAt   = "submittedAt"
	DoctorRemarks = "doctorRemarks"

	// AuthObjects
	Authorized  = "authorized"
	AccessUuid  = "accessUuid"
	UserId      = "userId"
	Exp         = "exp"
	RefreshUuid = "refreshUuid"
	Type        = "type"
	DisplayName = "displayName"
)

// Keywords
const (
	ASC  = "ASC"
	DESC = "DESC"
)
