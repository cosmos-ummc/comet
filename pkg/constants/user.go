package constants

const (
	AccessTokenTTLMinutes     = 15
	RefreshTokenTTLDays       = 7
	PasswordResetTokenTTLDays = 1
	ActivityTTLDays           = 7
)

const (
	Superuser  = "superuser"
	Consultant = "consultant"
	ChatBot    = "chatbot"
)

var SuperUserOnly = []string{Superuser}
var AllCanAccess = []string{Superuser, Consultant}
var ChatBotOnly = []string{ChatBot}
