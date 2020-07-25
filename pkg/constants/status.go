package constants

const (
	DontHaveToCall = iota + 1
	PatientCalled
	UMMCCalled
	NoYetCall
)

const (
	Undeclared = iota + 1
	Declared
)

var CallingStatusMap = map[int64]string{
	DontHaveToCall: "dontHaveToCallCount",
	PatientCalled:  "patientCalledCount",
	UMMCCalled:     "ummcCalledCount",
	NoYetCall:      "noYetCallCount",
}

var DeclarationStatusMap = map[int64]string{
	Undeclared: "undeclaredCount",
	Declared:   "declaredCount",
}

var PatientStatusMap = map[int64]string{
	Symptomatic:             "symptomatic",
	Asymptomatic:            "asymptomatic",
	ConfirmedButNotAdmitted: "confirmedButNotAdmitted",
	ConfirmedAndAdmitted:    "confirmedAndAdmitted",
	Completed:               "completed",
	Quit:                    "quit",
	Recovered:               "recovered",
	PassedAway:              "passedAway",
}
