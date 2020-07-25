package constants

const (
	AccessTokenTTLMinutes     = 15
	RefreshTokenTTLDays       = 7
	PasswordResetTokenTTLDays = 1
	ActivityTTLDays           = 7
)

const (
	Superuser    = "superuser"
	PuiAdmin     = "puiadmin"
	ContactAdmin = "contactadmin"
	PuiIk        = "puiik"
	ContactIk    = "contactik"
	ChatBot      = "chatbot"
)

var SuperUserOnly = []string{Superuser}
var AllCanAccess = []string{Superuser, PuiAdmin, ContactAdmin, PuiIk, ContactIk}
var ChatBotOnly = []string{ChatBot}

// UserPatientMap is used for access-control. It maps user roles to their respective patient types
// Role that is mapped to AllPatients can access to both PUI and ContactTracing patients
var UserPatientMap = map[string]int64{
	ChatBot:      AllPatients,
	Superuser:    AllPatients,
	PuiAdmin:     PUI,
	ContactAdmin: ContactTracing,
	PuiIk:        PUI,
	ContactIk:    ContactTracing,
}
