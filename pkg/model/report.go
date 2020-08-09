package model

import (
	"comet/pkg/dto"
	"comet/pkg/utility"
	"context"
	"time"
)

// GetReport gets report
func (m *Model) GetReport(ctx context.Context, dateString string) (*dto.Report, error) {
	report, err := m.reportDAO.Get(ctx, dateString)
	if err != nil {
		return nil, err
	}
	return report, nil
}

// GetReports get reports given from and to date (inclusive)
func (m *Model) GetReports(ctx context.Context, from, to string) ([]*dto.Report, error) {
	reports, err := m.reportDAO.BatchGet(ctx, utility.GetDatesByRange(from, to))
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// GenerateReport force generates report given a new date (based on latest data)
func (m *Model) GenerateReport(ctx context.Context, date string) error {
	_, err := utility.DateStringToTime(date)
	if err != nil {
		return err
	}

	// sync days
	err = m.SyncDays(ctx)
	if err != nil {
		return err
	}

	return nil
}

// SyncPatientReport sync patient record when changed type
func (m *Model) SyncPatientReport(ctx context.Context) error {
	date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	// get today 12 am timestamp
	_, err := utility.DateStringToTime(date)
	if err != nil {
		return err
	}

	return nil
}

// SyncDays ...
func (m *Model) SyncDays(ctx context.Context) error {
	_, patients, err := m.patientDAO.Query(ctx, nil, nil, nil)
	if err != nil {
		return err
	}

	for _, p := range patients {
		if p.Consent == 0 {
			p.Consent = 0
		} else {
			t := utility.MilliToTime(p.Consent)
			p.DaySinceMonitoring = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
		}

		// update patient
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			return err
		}
	}
	return nil
}
