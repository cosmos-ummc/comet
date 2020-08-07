package constants

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
