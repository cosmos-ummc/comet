package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/utility"
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ClientCreateDeclaration creates new declaration or updates existing declarations
func (m *Model) ClientCreateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error) {

	// process some data
	now := utility.MalaysiaTime(time.Now())
	declaration.SubmittedAt = utility.TimeToMilli(now)
	declaration.Date = utility.TimeToDateString(now)

	// check if patient exist, put patient name
	p, err := m.GetPatient(ctx, declaration.PatientID, constants.AllPatients)
	if err != nil {
		return nil, constants.PatientNotFoundError
	}
	declaration.PatientName = p.Name
	declaration.PatientPhoneNumber = p.PhoneNumber
	declaration.PatientType = p.Type

	// if no fever and recorded date is today, set fever start date to 0
	if declaration.Fever == constants.NoFever && p.FeverStartDate == utility.TimeToDateString(utility.MalaysiaTime(time.Now())) {
		p.FeverStartDate = ""
		p.FeverContDay = 0
	}
	// if have fever and no recorded date, set fever start date to today
	if declaration.Fever != constants.NoFever && p.FeverStartDate == "" {
		p.FeverStartDate = utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	}
	// calculate days since fever
	if p.FeverStartDate != "" {
		t, err := utility.DateStringToTime(p.FeverStartDate)
		if err != nil {
			return nil, err
		}
		p.FeverContDay = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
	}
	// if fever for 3 days continuous, is hasSymptom
	if declaration.Fever != constants.NoFever && p.FeverContDay >= 3 {
		declaration.HasSymptom = true
	}

	// if patient status is invalid, return error
	if p.Status != constants.Symptomatic && p.Status != constants.Asymptomatic && p.Status != constants.ConfirmedButNotAdmitted {
		return nil, constants.InvalidPatientStatusError
	}

	// update patient
	oldLastDeclared := p.LastDeclared
	p.LastDeclared = declaration.SubmittedAt
	p.LastDeclareResult = declaration.HasSymptom

	// check if declaration exists
	d, err := m.declarationDAO.Get(ctx, declaration.ID, constants.AllPatients)
	var resp *dto.Declaration
	if err != nil {
		if declaration.HasSymptom {
			declaration.CallingStatus = constants.NoYetCall
		} else {
			declaration.CallingStatus = constants.DontHaveToCall
		}
		// create declaration
		resp, err = m.declarationDAO.Create(ctx, declaration)
		if err != nil {
			return nil, err
		}

		// add into calling report
		err = m.callingReportDAO.UpdateByFields(ctx, declaration.Date, "", constants.CallingStatusMap[declaration.CallingStatus], p.Type)
		if err != nil {
			return nil, err
		}
	} else {
		// if today already got declaration, update declaration
		oldCallingStatus := d.CallingStatus

		// patch declaration
		d.Fever = declaration.Fever
		d.Drowsy = declaration.Drowsy
		d.Blue = declaration.Blue
		d.Chest = declaration.Chest
		d.Breathe = declaration.Breathe
		d.Throat = declaration.Throat
		d.Cough = declaration.Cough
		d.HasSymptom = declaration.HasSymptom
		d.SubmittedAt = declaration.SubmittedAt
		d.Loss = declaration.Loss
		if d.HasSymptom && oldCallingStatus == constants.DontHaveToCall {
			d.CallingStatus = constants.NoYetCall
		} else if !d.HasSymptom && oldCallingStatus != constants.DontHaveToCall {
			d.CallingStatus = constants.DontHaveToCall
		}

		// update declaration
		resp, err = m.declarationDAO.Update(ctx, d)
		if err != nil {
			return nil, err
		}

		// update calling report if needed
		if oldCallingStatus != d.CallingStatus {
			err = m.callingReportDAO.UpdateByFields(ctx, declaration.Date, constants.CallingStatusMap[oldCallingStatus], constants.CallingStatusMap[d.CallingStatus], p.Type)
			if err != nil {
				return nil, err
			}
		}
	}

	// update declaration if patient oldLastDeclared is less than today
	todayTime, err := utility.DateStringToTime(today)
	if err != nil {
		return nil, err
	}
	if oldLastDeclared == 0 || oldLastDeclared < utility.TimeToMilli(todayTime) {
		// move undeclared to declared
		err = m.declarationReportDAO.UpdateByFields(ctx, today,
			constants.DeclarationStatusMap[constants.Undeclared],
			constants.DeclarationStatusMap[constants.Declared], p.Type)
		if err != nil {
			return nil, err
		}
	}

	// update patient lastDeclareResult and lastDeclared
	_, err = m.patientDAO.Update(ctx, p)
	if err != nil {
		logger.Log.Error(fmt.Sprintf("failed to update patient lastDeclared, id: %s | patientId: %s", declaration.ID, declaration.PatientID))
	}

	return resp, err
}

// CreateDeclaration creates new declaration
func (m *Model) CreateDeclaration(ctx context.Context, declaration *dto.Declaration, patientType int64) (*dto.Declaration, error) {

	// process some data
	declaration.HasSymptom = utility.ValidateDeclarationSymptom(declaration)
	declaration.ID = declaration.PatientID + "_" + utility.TimeToDateString(utility.MilliToTime(declaration.SubmittedAt))
	if declaration.SubmittedAt != 0 {
		declaration.Date = utility.TimeToDateString(utility.MilliToTime(declaration.SubmittedAt))
	} else {
		t, err := utility.DateStringToTime(declaration.Date)
		if err == nil {
			declaration.SubmittedAt = utility.TimeToMilli(t)
		}
	}
	if declaration.HasSymptom {
		declaration.CallingStatus = constants.NoYetCall
	} else {
		declaration.CallingStatus = constants.DontHaveToCall
	}

	// check if declaration exist
	_, err := m.declarationDAO.Get(ctx, declaration.ID, constants.AllPatients)

	// only can create declaration if not found
	if err != nil && status.Code(err) == codes.Unknown {

		// check if patient exist, put patient name
		p, err := m.GetPatient(ctx, declaration.PatientID, patientType)
		if err != nil {
			return nil, constants.PatientNotFoundError
		}
		declaration.PatientName = p.Name
		declaration.PatientPhoneNumber = p.PhoneNumber
		declaration.PatientType = p.Type
		// if fever for 3 days continuous, is hasSymptom
		if p.FeverContDay >= 3 {
			declaration.HasSymptom = true
		}

		// create declaration
		resp, err := m.declarationDAO.Create(ctx, declaration)
		if err != nil {
			return nil, err
		}

		// update patient status, lastDeclareResult and lastDeclared if declaration is latest
		_, date := utility.DeclarationIDToPatientIDAndDate(declaration.ID)
		if date == utility.TimeToDateString(utility.MalaysiaTime(time.Now())) {
			p.LastDeclared = declaration.SubmittedAt
			p.LastDeclareResult = declaration.HasSymptom
			_, err = m.patientDAO.Update(ctx, p)
			if err != nil {
				logger.Log.Error(fmt.Sprintf("failed to update patient lastDeclared, id: %s | patientId: %s", declaration.ID, declaration.PatientID))
			}
		}

		// add into calling report
		err = m.callingReportDAO.UpdateByFields(ctx, declaration.Date, "", constants.CallingStatusMap[declaration.CallingStatus], p.Type)
		if err != nil {
			return nil, err
		}

		// move undeclared to declared
		err = m.declarationReportDAO.UpdateByFields(ctx, date,
			constants.DeclarationStatusMap[constants.Undeclared],
			constants.DeclarationStatusMap[constants.Declared], p.Type)
		if err != nil {
			return nil, err
		}

		return resp, err
	}

	if err != nil {
		return nil, err
	}

	return nil, status.Error(codes.AlreadyExists, "Declaration already exist!")
}

// GetDeclaration gets declaration
func (m *Model) GetDeclaration(ctx context.Context, declarationID string, patientType int64) (*dto.Declaration, error) {
	return m.declarationDAO.Get(ctx, declarationID, patientType)
}

// UpdateDeclaration updates declaration
func (m *Model) UpdateDeclaration(ctx context.Context, declaration *dto.Declaration, patientType int64) (*dto.Declaration, error) {

	// check if patient exist
	p, err := m.GetPatient(ctx, declaration.PatientID, constants.AllPatients)
	if err != nil {
		return nil, constants.PatientNotFoundError
	}

	// check if declaration exist
	d, err := m.declarationDAO.Get(ctx, declaration.ID, patientType)
	if err != nil {
		return nil, err
	}
	oldCallingStatus := d.CallingStatus

	// patch declaration
	d.Fever = declaration.Fever
	d.Drowsy = declaration.Drowsy
	d.Blue = declaration.Blue
	d.Chest = declaration.Chest
	d.Breathe = declaration.Breathe
	d.Throat = declaration.Throat
	d.Cough = declaration.Cough
	d.CallingStatus = declaration.CallingStatus
	d.DoctorRemarks = declaration.DoctorRemarks
	d.HasSymptom = declaration.HasSymptom
	d.Loss = declaration.Loss

	// validate calling status with symptom
	if d.HasSymptom && d.CallingStatus == constants.DontHaveToCall {
		d.CallingStatus = constants.NoYetCall
	} else if !d.HasSymptom && d.CallingStatus != constants.DontHaveToCall {
		d.CallingStatus = constants.DontHaveToCall
	}

	// update declaration
	_, err = m.declarationDAO.Update(ctx, d)
	if err != nil {
		return nil, err
	}

	// If latest declaration is updated, update patient too.
	if d.Date == utility.TimeToDateString(utility.MalaysiaTime(time.Now())) {
		// update patient lastDeclareResult and lastDeclared
		p.LastDeclared = d.SubmittedAt
		p.LastDeclareResult = d.HasSymptom
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			logger.Log.Error(fmt.Sprintf("failed to update patient lastDeclared, id: %s | patientId: %s", declaration.ID, declaration.PatientID))
		}

		// update calling status
		if oldCallingStatus != d.CallingStatus {
			// add into calling report if calling status changed
			err = m.callingReportDAO.UpdateByFields(ctx, declaration.Date,
				constants.CallingStatusMap[oldCallingStatus],
				constants.CallingStatusMap[d.CallingStatus], p.Type)
			if err != nil {
				return nil, err
			}
		}
	}

	return d, err
}

// UpdateDeclarations update declarations
func (m *Model) UpdateDeclarations(ctx context.Context, declaration *dto.Declaration, ids []string, patientType int64) ([]string, error) {
	// TODO: Support Batch Updates
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	declaration.ID = ids[0]
	d, err := m.UpdateDeclaration(ctx, declaration, patientType)
	if err != nil {
		return nil, err
	}
	return []string{d.ID}, nil
}

// DeleteDeclaration deletes declaration
func (m *Model) DeleteDeclaration(ctx context.Context, declarationID string, patientType int64) (*dto.Declaration, error) {
	d, err := m.GetDeclaration(ctx, declarationID, patientType)
	if err != nil {
		return nil, err
	}

	// get patient
	p, err := m.GetPatient(ctx, d.PatientID, constants.AllPatients)
	if err != nil {
		return nil, err
	}

	err = m.declarationDAO.Delete(ctx, declarationID)
	if err != nil {
		return nil, err
	}

	// If latest declaration is deleted, update report counters
	if d.Date == utility.TimeToDateString(utility.MalaysiaTime(time.Now())) {
		// get the previous declarations
		filter := map[string]interface{}{constants.PatientID: d.PatientID}
		_, prevDeclarations, err := m.declarationDAO.Query(ctx, &dto.SortData{
			Item:  constants.SubmittedAt,
			Order: constants.DESC,
		}, nil, filter, constants.AllPatients)
		if err != nil {
			return nil, err
		}

		// update patient lastDeclared
		if len(prevDeclarations) == 0 {
			p.LastDeclared = 0
			p.LastDeclareResult = false
		} else {
			p.LastDeclared = prevDeclarations[0].SubmittedAt
			p.LastDeclareResult = prevDeclarations[0].HasSymptom
		}
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			return nil, err
		}

		// remove from calling report
		err = m.callingReportDAO.UpdateByFields(ctx, d.Date, constants.CallingStatusMap[d.CallingStatus], "", p.Type)
		if err != nil {
			return nil, err
		}

		// remove declared from declaration report, add into undeclared
		err = m.declarationReportDAO.UpdateByFields(ctx, d.Date,
			constants.DeclarationStatusMap[constants.Declared],
			constants.DeclarationStatusMap[constants.Undeclared], p.Type)
		if err != nil {
			return nil, err
		}
	}

	return d, nil
}

// BatchGetDeclarations get declarations
func (m *Model) BatchGetDeclarations(ctx context.Context, declarationIDs []string, patientType int64) ([]*dto.Declaration, error) {
	return m.declarationDAO.BatchGet(ctx, declarationIDs, patientType)
}

// DeleteDeclarations delete declarations
func (m *Model) DeleteDeclarations(ctx context.Context, declarationIDs []string, patientType int64) ([]string, error) {

	var deletedIDs []string
	for _, id := range declarationIDs {
		d, err := m.DeleteDeclaration(ctx, id, patientType)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, d.ID)
	}

	return deletedIDs, nil
}

// QueryDeclarations query declarations
func (m *Model) QueryDeclarations(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Declaration, error) {
	return m.declarationDAO.Query(ctx, sort, itemsRange, filter, patientType)
}

// QueryDeclarations query declarations by time
func (m *Model) QueryDeclarationsByTime(ctx context.Context, from int64, patientType int64) (int64, []*dto.Declaration, error) {
	return m.declarationDAO.QueryByTime(ctx, from, patientType)
}

// QueryDeclarationsByCallingStatusAndTime ...
func (m *Model) QueryDeclarationsByCallingStatusAndTime(ctx context.Context, callingStatus []int64, from int64, patientType int64) (int64, []*dto.Declaration, error) {
	return m.declarationDAO.QueryByCallingStatusAndTime(ctx, callingStatus, from, patientType)
}

// QueryDeclarationsByPatientID ...
func (m *Model) QueryDeclarationsByPatientID(ctx context.Context, id string, patientType int64) (int64, []*dto.Declaration, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	total, declarations, err := m.declarationDAO.Query(ctx, nil, nil, filter, patientType)
	return total, declarations, err
}

// Get stable declarations
func (m *Model) GetStableDeclarations(ctx context.Context, from int64, patientType int64) (int64, []*dto.Declaration, error) {
	return m.declarationDAO.QueryStableDeclarations(ctx, from, patientType)
}
