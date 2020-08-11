package constants

const (
	AccessTokenTTLMinutes     = 15
	RefreshTokenTTLDays       = 7
	PasswordResetTokenTTLDays = 1
	ActivityTTLDays           = 7
)

const (
	Superuser   = "superuser"
	Consultant  = "consultant"
	PatientUser = "patient"
	ChatBot     = "chatbot"
)

var SuperUserOnly = []string{Superuser}
var AllCanAccess = []string{Superuser, Consultant, PatientUser}
var ChatBotOnly = []string{ChatBot}
