package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/utility"
	"context"
	"time"
)

// GetDeclarationReport gets declaration report
func (m *Model) GetDeclarationReport(ctx context.Context, dateString string, patientType int64) (*dto.DeclarationReport, error) {

	report, err := m.declarationReportDAO.Get(ctx, dateString)
	if err != nil {
		return nil, err
	}
	if patientType == constants.PUI {
		return &dto.DeclarationReport{
			Date:            report.Date,
			UndeclaredCount: report.PuiUndeclaredCount,
			DeclaredCount:   report.PuiDeclaredCount,
		}, nil
	} else if patientType == constants.ContactTracing {
		return &dto.DeclarationReport{
			Date:            report.Date,
			UndeclaredCount: report.ContactUndeclaredCount,
			DeclaredCount:   report.ContactDeclaredCount,
		}, nil
	}
	return &dto.DeclarationReport{
		Date:            report.Date,
		UndeclaredCount: report.PuiUndeclaredCount + report.ContactUndeclaredCount,
		DeclaredCount:   report.PuiDeclaredCount + report.ContactDeclaredCount,
	}, nil
}

// GetDeclarationReports get declaration reports given from and to date (inclusive)
func (m *Model) GetDeclarationReports(ctx context.Context, from, to string, patientType int64) ([]*dto.DeclarationReport, error) {
	reports, err := m.declarationReportDAO.BatchGet(ctx, utility.GetDatesByRange(from, to))
	if err != nil {
		return nil, err
	}
	var rslt []*dto.DeclarationReport
	for _, report := range reports {
		if patientType == constants.PUI {
			rslt = append(rslt, &dto.DeclarationReport{
				Date:            report.Date,
				UndeclaredCount: report.PuiUndeclaredCount,
				DeclaredCount:   report.PuiDeclaredCount,
			})
		} else if patientType == constants.ContactTracing {
			rslt = append(rslt, &dto.DeclarationReport{
				Date:            report.Date,
				UndeclaredCount: report.ContactUndeclaredCount,
				DeclaredCount:   report.ContactDeclaredCount,
			})
		} else {
			rslt = append(rslt, &dto.DeclarationReport{
				Date:            report.Date,
				UndeclaredCount: report.PuiUndeclaredCount + report.ContactUndeclaredCount,
				DeclaredCount:   report.PuiDeclaredCount + report.ContactDeclaredCount,
			})
		}
	}
	return rslt, nil
}

// GetCallingReport gets calling report
func (m *Model) GetCallingReport(ctx context.Context, date string, patientType int64) (*dto.CallingReport, error) {
	report, err := m.callingReportDAO.Get(ctx, date)
	if err != nil {
		return nil, err
	}
	if patientType == constants.PUI {
		return &dto.CallingReport{
			Date:           report.Date,
			DontHaveToCall: report.PuiDontHaveToCall,
			PatientCalled:  report.PuiPatientCalled,
			UMMCCalled:     report.PuiUMMCCalled,
			NoYetCall:      report.PuiNoYetCall,
		}, nil
	} else if patientType == constants.ContactTracing {
		return &dto.CallingReport{
			Date:           report.Date,
			DontHaveToCall: report.ContactDontHaveToCall,
			PatientCalled:  report.ContactPatientCalled,
			UMMCCalled:     report.ContactUMMCCalled,
			NoYetCall:      report.ContactNoYetCall,
		}, nil
	}
	return &dto.CallingReport{
		Date:           report.Date,
		DontHaveToCall: report.PuiDontHaveToCall + report.ContactDontHaveToCall,
		PatientCalled:  report.PuiPatientCalled + report.ContactPatientCalled,
		UMMCCalled:     report.PuiUMMCCalled + report.ContactUMMCCalled,
		NoYetCall:      report.PuiNoYetCall + report.ContactNoYetCall,
	}, nil
}

// GetCallingReports get calling reports given from and to date (inclusive)
func (m *Model) GetCallingReports(ctx context.Context, from, to string, patientType int64) ([]*dto.CallingReport, error) {
	reports, err := m.callingReportDAO.BatchGet(ctx, utility.GetDatesByRange(from, to))
	if err != nil {
		return nil, err
	}
	var rslt []*dto.CallingReport
	for _, report := range reports {
		if patientType == constants.PUI {
			rslt = append(rslt, &dto.CallingReport{
				Date:           report.Date,
				DontHaveToCall: report.PuiDontHaveToCall,
				PatientCalled:  report.PuiPatientCalled,
				UMMCCalled:     report.PuiUMMCCalled,
				NoYetCall:      report.PuiNoYetCall,
			})
		} else if patientType == constants.ContactTracing {
			rslt = append(rslt, &dto.CallingReport{
				Date:           report.Date,
				DontHaveToCall: report.ContactDontHaveToCall,
				PatientCalled:  report.ContactPatientCalled,
				UMMCCalled:     report.ContactUMMCCalled,
				NoYetCall:      report.ContactNoYetCall,
			})
		} else {
			rslt = append(rslt, &dto.CallingReport{
				Date:           report.Date,
				DontHaveToCall: report.PuiDontHaveToCall + report.ContactDontHaveToCall,
				PatientCalled:  report.PuiPatientCalled + report.ContactPatientCalled,
				UMMCCalled:     report.PuiUMMCCalled + report.ContactUMMCCalled,
				NoYetCall:      report.PuiNoYetCall + report.ContactNoYetCall,
			})
		}
	}
	return rslt, nil
}

// GetPatientReport gets patient status report
func (m *Model) GetPatientStatusReport(ctx context.Context, dateString string, patientType int64) (*dto.PatientStatusReport, error) {
	report, err := m.patientStatusReportDAO.Get(ctx, dateString)
	if err != nil {
		return nil, err
	}
	if patientType == constants.PUI {
		return &dto.PatientStatusReport{
			Date:                    report.Date,
			Symptomatic:             report.PuiSymptomatic,
			Asymptomatic:            report.PuiAsymptomatic,
			ConfirmedButNotAdmitted: report.PuiConfirmedButNotAdmitted,
			ConfirmedAndAdmitted:    report.PuiConfirmedAndAdmitted,
			Completed:               report.PuiCompleted,
			Quit:                    report.PuiQuit,
			Recovered:               report.PuiRecovered,
			PassedAway:              report.PuiPassedAway,
		}, nil
	} else if patientType == constants.ContactTracing {
		return &dto.PatientStatusReport{
			Date:                    report.Date,
			Symptomatic:             report.ContactSymptomatic,
			Asymptomatic:            report.ContactAsymptomatic,
			ConfirmedButNotAdmitted: report.ContactConfirmedButNotAdmitted,
			ConfirmedAndAdmitted:    report.ContactConfirmedAndAdmitted,
			Completed:               report.ContactCompleted,
			Quit:                    report.ContactQuit,
			Recovered:               report.ContactRecovered,
			PassedAway:              report.ContactPassedAway,
		}, nil
	}
	return &dto.PatientStatusReport{
		Date:                    report.Date,
		Symptomatic:             report.PuiSymptomatic + report.ContactSymptomatic,
		Asymptomatic:            report.PuiAsymptomatic + report.ContactAsymptomatic,
		ConfirmedButNotAdmitted: report.PuiConfirmedButNotAdmitted + report.ContactConfirmedButNotAdmitted,
		ConfirmedAndAdmitted:    report.PuiConfirmedAndAdmitted + report.ContactConfirmedAndAdmitted,
		Completed:               report.PuiCompleted + report.ContactCompleted,
		Quit:                    report.PuiQuit + report.ContactQuit,
		Recovered:               report.PuiRecovered + report.ContactRecovered,
		PassedAway:              report.PuiPassedAway + report.ContactPassedAway,
	}, nil
}

// GetPatientStatusReports get patient status reports given from and to date (inclusive)
func (m *Model) GetPatientStatusReports(ctx context.Context, from, to string, patientType int64) ([]*dto.PatientStatusReport, error) {
	reports, err := m.patientStatusReportDAO.BatchGet(ctx, utility.GetDatesByRange(from, to))
	if err != nil {
		return nil, err
	}
	var rslt []*dto.PatientStatusReport
	for _, report := range reports {
		if patientType == constants.PUI {
			rslt = append(rslt, &dto.PatientStatusReport{
				Date:                    report.Date,
				Symptomatic:             report.PuiSymptomatic,
				Asymptomatic:            report.PuiAsymptomatic,
				ConfirmedButNotAdmitted: report.PuiConfirmedButNotAdmitted,
				ConfirmedAndAdmitted:    report.PuiConfirmedAndAdmitted,
				Completed:               report.PuiCompleted,
				Quit:                    report.PuiQuit,
				Recovered:               report.PuiRecovered,
				PassedAway:              report.PuiPassedAway,
			})
		} else if patientType == constants.ContactTracing {
			rslt = append(rslt, &dto.PatientStatusReport{
				Date:                    report.Date,
				Symptomatic:             report.ContactSymptomatic,
				Asymptomatic:            report.ContactAsymptomatic,
				ConfirmedButNotAdmitted: report.ContactConfirmedButNotAdmitted,
				ConfirmedAndAdmitted:    report.ContactConfirmedAndAdmitted,
				Completed:               report.ContactCompleted,
				Quit:                    report.ContactQuit,
				Recovered:               report.ContactRecovered,
				PassedAway:              report.ContactPassedAway,
			})
		} else {
			rslt = append(rslt, &dto.PatientStatusReport{
				Date:                    report.Date,
				Symptomatic:             report.PuiSymptomatic + report.ContactSymptomatic,
				Asymptomatic:            report.PuiAsymptomatic + report.ContactAsymptomatic,
				ConfirmedButNotAdmitted: report.PuiConfirmedButNotAdmitted + report.ContactConfirmedButNotAdmitted,
				ConfirmedAndAdmitted:    report.PuiConfirmedAndAdmitted + report.ContactConfirmedAndAdmitted,
				Completed:               report.PuiCompleted + report.ContactCompleted,
				Quit:                    report.PuiQuit + report.ContactQuit,
				Recovered:               report.PuiRecovered + report.ContactRecovered,
				PassedAway:              report.PuiPassedAway + report.ContactPassedAway,
			})
		}
	}
	return rslt, nil
}

// GenerateReport force generates report given a new date (based on latest data)
func (m *Model) GenerateReport(ctx context.Context, date string) error {
	t, err := utility.DateStringToTime(date)
	if err != nil {
		return err
	}

	// sync days
	err = m.SyncDays(ctx)
	if err != nil {
		return err
	}

	// for calling status report, get all declared patients and check if they no yet call
	// get today's declarations
	_, declarations, err := m.QueryDeclarationsByTime(ctx, utility.TimeToMilli(t), constants.AllPatients)
	if err != nil {
		return err
	}

	puiNoYetCallPatientNum := int64(0)
	puiDontHaveToCallPatientNum := int64(0)
	puiPatientCalledNum := int64(0)
	puiUmmcCalledNum := int64(0)
	contactNoYetCallPatientNum := int64(0)
	contactDontHaveToCallPatientNum := int64(0)
	contactPatientCalledNum := int64(0)
	contactUmmcCalledNum := int64(0)

	for _, declaration := range declarations {
		p, err := m.GetPatient(ctx, declaration.PatientID, constants.AllPatients)
		if err != nil {
			return err
		}
		// skip patients that are other status
		if p.Status != constants.Symptomatic && p.Status != constants.Asymptomatic && p.Status != constants.ConfirmedButNotAdmitted {
			continue
		}

		if declaration.CallingStatus == constants.NoYetCall {
			if p.Type == constants.PUI {
				puiNoYetCallPatientNum += 1
			} else if p.Type == constants.ContactTracing {
				contactNoYetCallPatientNum += 1
			}
		}
		if declaration.CallingStatus == constants.DontHaveToCall {
			if p.Type == constants.PUI {
				puiDontHaveToCallPatientNum += 1
			} else if p.Type == constants.ContactTracing {
				contactDontHaveToCallPatientNum += 1
			}
		}
		if declaration.CallingStatus == constants.PatientCalled {
			if p.Type == constants.PUI {
				puiPatientCalledNum += 1
			} else if p.Type == constants.ContactTracing {
				contactPatientCalledNum += 1
			}
		}
		if declaration.CallingStatus == constants.UMMCCalled {
			if p.Type == constants.PUI {
				puiUmmcCalledNum += 1
			} else if p.Type == constants.ContactTracing {
				contactUmmcCalledNum += 1
			}
		}
	}
	c := &dto.CallingReport{
		PuiDontHaveToCall:     puiDontHaveToCallPatientNum,
		PuiPatientCalled:      puiPatientCalledNum,
		PuiUMMCCalled:         puiUmmcCalledNum,
		PuiNoYetCall:          puiNoYetCallPatientNum,
		ContactDontHaveToCall: contactDontHaveToCallPatientNum,
		ContactPatientCalled:  contactPatientCalledNum,
		ContactUMMCCalled:     contactUmmcCalledNum,
		ContactNoYetCall:      contactNoYetCallPatientNum,
	}
	err = m.callingReportDAO.Create(ctx, date, c)
	if err != nil {
		return err
	}

	// for patient status report, query all patients and update the the total, then update today.
	_, patients, err := m.patientDAO.Query(ctx, nil, nil, nil, constants.AllPatients)

	puiSymptomaticCount := int64(0)
	puiAsymptomaticCount := int64(0)
	puiConfirmedButNotAdmittedCount := int64(0)
	puiConfirmedAndAdmittedCount := int64(0)
	puiCompletedCount := int64(0)
	puiQuitCount := int64(0)
	puiRecoveredCount := int64(0)
	puiPassedAwayCount := int64(0)
	contactSymptomaticCount := int64(0)
	contactAsymptomaticCount := int64(0)
	contactConfirmedButNotAdmittedCount := int64(0)
	contactConfirmedAndAdmittedCount := int64(0)
	contactCompletedCount := int64(0)
	contactQuitCount := int64(0)
	contactRecoveredCount := int64(0)
	contactPassedAwayCount := int64(0)

	for _, patient := range patients {
		if patient.Status == constants.Symptomatic {
			if patient.Type == constants.PUI {
				puiSymptomaticCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactSymptomaticCount += 1
			}
		}
		if patient.Status == constants.Asymptomatic {
			if patient.Type == constants.PUI {
				puiAsymptomaticCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactAsymptomaticCount += 1
			}
		}
		if patient.Status == constants.ConfirmedButNotAdmitted {
			if patient.Type == constants.PUI {
				puiConfirmedButNotAdmittedCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactConfirmedButNotAdmittedCount += 1
			}
		}
		if patient.Status == constants.ConfirmedAndAdmitted {
			if patient.Type == constants.PUI {
				puiConfirmedAndAdmittedCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactConfirmedButNotAdmittedCount += 1
			}
		}
		if patient.Status == constants.Completed {
			if patient.Type == constants.PUI {
				puiCompletedCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactCompletedCount += 1
			}
		}
		if patient.Status == constants.Quit {
			if patient.Type == constants.PUI {
				puiQuitCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactQuitCount += 1
			}
		}
		if patient.Status == constants.Recovered {
			if patient.Type == constants.PUI {
				puiRecoveredCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactRecoveredCount += 1
			}
		}
		if patient.Status == constants.PassedAway {
			if patient.Type == constants.PUI {
				puiPassedAwayCount += 1
			} else if patient.Type == constants.ContactTracing {
				contactPassedAwayCount += 1
			}
		}
	}

	s := &dto.PatientStatusReport{
		PuiSymptomatic:                 puiSymptomaticCount,
		PuiAsymptomatic:                puiAsymptomaticCount,
		PuiConfirmedButNotAdmitted:     puiConfirmedButNotAdmittedCount,
		PuiConfirmedAndAdmitted:        puiConfirmedAndAdmittedCount,
		PuiCompleted:                   puiCompletedCount,
		PuiQuit:                        puiQuitCount,
		PuiRecovered:                   puiRecoveredCount,
		PuiPassedAway:                  puiPassedAwayCount,
		ContactSymptomatic:             contactSymptomaticCount,
		ContactAsymptomatic:            contactAsymptomaticCount,
		ContactConfirmedButNotAdmitted: contactConfirmedButNotAdmittedCount,
		ContactConfirmedAndAdmitted:    contactConfirmedAndAdmittedCount,
		ContactCompleted:               contactCompletedCount,
		ContactQuit:                    contactQuitCount,
		ContactRecovered:               contactRecoveredCount,
		ContactPassedAway:              contactPassedAwayCount,
	}

	// update today's count
	err = m.patientStatusReportDAO.Create(ctx, date, s)
	if err != nil {
		return err
	}

	// for declaration report
	// query undeclared patients (12 am) and set them to undeclared
	dateInTime, err := utility.DateStringToTime(date)
	if err != nil {
		return err
	}

	puiUndeclaredPatientNum, _, err := m.GetUndeclaredPatientsByTime(ctx, utility.TimeToMilli(dateInTime), nil, nil, constants.PUI)
	if err != nil {
		return err
	}
	puiDeclaredPatientNum := puiSymptomaticCount + puiAsymptomaticCount + puiConfirmedButNotAdmittedCount - puiUndeclaredPatientNum

	if puiDeclaredPatientNum < 0 {
		puiDeclaredPatientNum = 0
	}
	if puiUndeclaredPatientNum < 0 {
		puiUndeclaredPatientNum = 0
	}

	contactUndeclaredPatientNum, _, err := m.GetUndeclaredPatientsByTime(ctx, utility.TimeToMilli(dateInTime), nil, nil, constants.ContactTracing)
	if err != nil {
		return err
	}
	contactDeclaredPatientNum := contactSymptomaticCount + contactAsymptomaticCount + contactConfirmedButNotAdmittedCount - contactUndeclaredPatientNum

	if contactDeclaredPatientNum < 0 {
		contactDeclaredPatientNum = 0
	}
	if contactUndeclaredPatientNum < 0 {
		contactUndeclaredPatientNum = 0
	}

	// create declaration reports
	d := &dto.DeclarationReport{
		PuiUndeclaredCount:     puiUndeclaredPatientNum,
		PuiDeclaredCount:       puiDeclaredPatientNum,
		ContactUndeclaredCount: contactUndeclaredPatientNum,
		ContactDeclaredCount:   contactDeclaredPatientNum,
	}
	err = m.declarationReportDAO.Create(ctx, date, d)
	if err != nil {
		return err
	}

	return nil
}

// SyncPatientReport sync patient record when changed type
func (m *Model) SyncPatientReport(ctx context.Context) error {
	date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	// get today 12 am timestamp
	t, err := utility.DateStringToTime(date)
	if err != nil {
		return err
	}

	// for calling status report, get all declared patients and check if they no yet call
	_, declarations, err := m.QueryDeclarationsByTime(ctx, utility.TimeToMilli(t), constants.AllPatients)
	if err != nil {
		return err
	}

	puiNoYetCallPatientNum := int64(0)
	puiDontHaveToCallPatientNum := int64(0)
	puiPatientCalledNum := int64(0)
	puiUmmcCalledNum := int64(0)
	contactNoYetCallPatientNum := int64(0)
	contactDontHaveToCallPatientNum := int64(0)
	contactPatientCalledNum := int64(0)
	contactUmmcCalledNum := int64(0)

	for _, declaration := range declarations {
		p, err := m.GetPatient(ctx, declaration.PatientID, constants.AllPatients)
		if err != nil {
			return err
		}
		// skip patients that are other status
		if p.Status != constants.Symptomatic && p.Status != constants.Asymptomatic && p.Status != constants.ConfirmedButNotAdmitted {
			continue
		}

		if declaration.CallingStatus == constants.NoYetCall {
			if p.Type == constants.PUI {
				puiNoYetCallPatientNum += 1
			} else if p.Type == constants.ContactTracing {
				contactNoYetCallPatientNum += 1
			}
		}
		if declaration.CallingStatus == constants.DontHaveToCall {
			if p.Type == constants.PUI {
				puiDontHaveToCallPatientNum += 1
			} else if p.Type == constants.ContactTracing {
				contactDontHaveToCallPatientNum += 1
			}
		}
		if declaration.CallingStatus == constants.PatientCalled {
			if p.Type == constants.PUI {
				puiPatientCalledNum += 1
			} else if p.Type == constants.ContactTracing {
				contactPatientCalledNum += 1
			}
		}
		if declaration.CallingStatus == constants.UMMCCalled {
			if p.Type == constants.PUI {
				puiUmmcCalledNum += 1
			} else if p.Type == constants.ContactTracing {
				contactUmmcCalledNum += 1
			}
		}
	}
	c := &dto.CallingReport{
		PuiDontHaveToCall:     puiDontHaveToCallPatientNum,
		PuiPatientCalled:      puiPatientCalledNum,
		PuiUMMCCalled:         puiUmmcCalledNum,
		PuiNoYetCall:          puiNoYetCallPatientNum,
		ContactDontHaveToCall: contactDontHaveToCallPatientNum,
		ContactPatientCalled:  contactPatientCalledNum,
		ContactUMMCCalled:     contactUmmcCalledNum,
		ContactNoYetCall:      contactNoYetCallPatientNum,
	}
	err = m.callingReportDAO.Create(ctx, date, c)
	if err != nil {
		return err
	}

	// query undeclared patients (12 am) and set them to undeclared
	puiUndeclaredPatientNum, _, err := m.GetUndeclaredPatientsByTime(ctx, utility.TimeToMilli(t), nil, nil, constants.PUI)
	if err != nil {
		return err
	}
	puiDeclaredPatients, err := m.GetDeclaredPatientsByTime(ctx, utility.TimeToMilli(t), constants.PUI)
	if err != nil {
		return err
	}
	puiDeclaredPatientNum := int64(len(puiDeclaredPatients))

	if puiDeclaredPatientNum < 0 {
		puiDeclaredPatientNum = 0
	}
	if puiUndeclaredPatientNum < 0 {
		puiUndeclaredPatientNum = 0
	}

	contactUndeclaredPatientNum, _, err := m.GetUndeclaredPatientsByTime(ctx, utility.TimeToMilli(t), nil, nil, constants.ContactTracing)
	if err != nil {
		return err
	}
	contactDeclaredPatients, err := m.GetDeclaredPatientsByTime(ctx, utility.TimeToMilli(t), constants.ContactTracing)
	if err != nil {
		return err
	}
	contactDeclaredPatientNum := int64(len(contactDeclaredPatients))

	if contactDeclaredPatientNum < 0 {
		contactDeclaredPatientNum = 0
	}
	if contactUndeclaredPatientNum < 0 {
		contactUndeclaredPatientNum = 0
	}

	// create declaration reports
	d := &dto.DeclarationReport{
		PuiUndeclaredCount:     puiUndeclaredPatientNum,
		PuiDeclaredCount:       puiDeclaredPatientNum,
		ContactUndeclaredCount: contactUndeclaredPatientNum,
		ContactDeclaredCount:   contactDeclaredPatientNum,
	}
	err = m.declarationReportDAO.Create(ctx, date, d)
	if err != nil {
		return err
	}

	return nil
}

// SyncDays ...
func (m *Model) SyncDays(ctx context.Context) error {
	_, patients, err := m.patientDAO.Query(ctx, nil, nil, nil, constants.AllPatients)
	if err != nil {
		return err
	}

	for _, p := range patients {
		// calculate days since swab if needed
		if p.SwabDate == "" {
			p.DaysSinceSwab = 0
		} else {
			t, err := utility.DateStringToTime(p.SwabDate)
			if err != nil {
				return err
			}
			p.DaysSinceSwab = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
		}

		// calculate days since exposure if needed
		if p.ExposureDate == "" {
			p.DaysSinceExposure = 0
		} else {
			t, err := utility.DateStringToTime(p.ExposureDate)
			if err != nil {
				return err
			}
			p.DaysSinceExposure = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
		}

		// calculate days since fever if needed
		if p.FeverStartDate == "" {
			p.FeverContDay = 0
		} else {
			t, err := utility.DateStringToTime(p.FeverStartDate)
			if err != nil {
				return err
			}
			p.FeverContDay = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
		}

		// update patient
		_, err = m.UpdatePatient(ctx, p, constants.AllPatients, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
