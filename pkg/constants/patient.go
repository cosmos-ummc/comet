package constants

const (
	PUI = iota + 1
	PUS
	HCW
	Patient
	Others
)

const (
	NoMentalIssue = iota
	Depression
	Anxiety
	Stress
	PTSD
)

// Registration Status
const (
	NotFound = iota
	Complete
	Incomplete
)

// Swab Result
const (
	SwabNone = iota
	SwabPending
	SwabPositive
	SwabNegative
)
