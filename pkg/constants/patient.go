package constants

const (
	// WARNING! AllPatient is NOT a valid patient type value. It is used solely for access control
	AllPatients = iota

	PUI
	ContactTracing
)

// PatientFieldMap maps patient report fields
var PatientFieldMap = map[int64]string{
	AllPatients:    "",
	PUI:            "pui",
	ContactTracing: "contact",
}

const (
	English = iota + 1
	Malay
	Chinese
)

const (
	Symptomatic = iota + 1
	Asymptomatic
	ConfirmedButNotAdmitted
	ConfirmedAndAdmitted
	Completed
	Recovered
	Quit
	PassedAway
)

// Registration Status
const (
	NotFound = iota
	Complete
	Incomplete
)
