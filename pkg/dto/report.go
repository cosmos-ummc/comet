package dto

type PatientStatusReport struct {
	Date string `json:"date" bson:"date"`
	// report to return to user
	Symptomatic             int64 `json:"symptomatic" bson:"-"`
	Asymptomatic            int64 `json:"asymptomatic" bson:"-"`
	ConfirmedButNotAdmitted int64 `json:"confirmedButNotAdmitted" bson:"-"`
	ConfirmedAndAdmitted    int64 `json:"confirmedAndAdmitted" bson:"-"`
	Completed               int64 `json:"completed" bson:"-"`
	Quit                    int64 `json:"quit" bson:"-"`
	Recovered               int64 `json:"recovered" bson:"-"`
	PassedAway              int64 `json:"passedAway" bson:"-"`
	// PUI patients
	PuiSymptomatic             int64 `json:"-" bson:"puisymptomatic"`
	PuiAsymptomatic            int64 `json:"-" bson:"puiasymptomatic"`
	PuiConfirmedButNotAdmitted int64 `json:"-" bson:"puiconfirmedButNotAdmitted"`
	PuiConfirmedAndAdmitted    int64 `json:"-" bson:"puiconfirmedAndAdmitted"`
	PuiCompleted               int64 `json:"-" bson:"puicompleted"`
	PuiQuit                    int64 `json:"-" bson:"puiquit"`
	PuiRecovered               int64 `json:"-" bson:"puirecovered"`
	PuiPassedAway              int64 `json:"-" bson:"puipassedAway"`
	// Contact tracing patients
	ContactSymptomatic             int64 `json:"-" bson:"contactsymptomatic"`
	ContactAsymptomatic            int64 `json:"-" bson:"contactasymptomatic"`
	ContactConfirmedButNotAdmitted int64 `json:"-" bson:"contactconfirmedButNotAdmitted"`
	ContactConfirmedAndAdmitted    int64 `json:"-" bson:"contactconfirmedAndAdmitted"`
	ContactCompleted               int64 `json:"-" bson:"contactcompleted"`
	ContactQuit                    int64 `json:"-" bson:"contactquit"`
	ContactRecovered               int64 `json:"-" bson:"contactrecovered"`
	ContactPassedAway              int64 `json:"-" bson:"contactpassedAway"`
}

type DeclarationReport struct {
	Date string `json:"date" bson:"date"`
	// report to return to user
	UndeclaredCount int64 `json:"undeclaredCount" bson:"-"`
	DeclaredCount   int64 `json:"declaredCount" bson:"-"`
	// PUI patients
	PuiUndeclaredCount int64 `json:"-" bson:"puiundeclaredCount"`
	PuiDeclaredCount   int64 `json:"-" bson:"puideclaredCount"`
	// Contact tracing patients
	ContactUndeclaredCount int64 `json:"-" bson:"contactundeclaredCount"`
	ContactDeclaredCount   int64 `json:"-" bson:"contactdeclaredCount"`
}

type CallingReport struct {
	Date string `json:"date" bson:"date"`
	// report to return to user
	DontHaveToCall int64 `json:"dontHaveToCallCount" bson:"-"`
	PatientCalled  int64 `json:"patientCalledCount" bson:"-"`
	UMMCCalled     int64 `json:"ummcCalledCount" bson:"-"`
	NoYetCall      int64 `json:"noYetCallCount" bson:"-"`
	// PUI patients
	PuiDontHaveToCall int64 `json:"-" bson:"puidontHaveToCallCount"`
	PuiPatientCalled  int64 `json:"-" bson:"puipatientCalledCount"`
	PuiUMMCCalled     int64 `json:"-" bson:"puiummcCalledCount"`
	PuiNoYetCall      int64 `json:"-" bson:"puinoYetCallCount"`
	// Contact tracing patients
	ContactDontHaveToCall int64 `json:"-" bson:"contactdontHaveToCallCount"`
	ContactPatientCalled  int64 `json:"-" bson:"contactpatientCalledCount"`
	ContactUMMCCalled     int64 `json:"-" bson:"contactummcCalledCount"`
	ContactNoYetCall      int64 `json:"-" bson:"contactnoYetCallCount"`
}
