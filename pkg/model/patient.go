package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/utility"
	"context"
	"github.com/twinj/uuid"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePatient creates new patient
func (m *Model) CreatePatient(ctx context.Context, patient *dto.Patient, user *dto.User) (*dto.Patient, error) {

	// check if patient exist
	_, err := m.patientDAO.Get(ctx, patient.ID, constants.AllPatients)

	// only can create patient if not found
	if err != nil && status.Code(err) == codes.Unknown {

		// query patients by phone number
		filter := map[string]interface{}{constants.PhoneNumber: patient.PhoneNumber}
		_, patients, err := m.patientDAO.Query(ctx, nil, nil, filter, constants.AllPatients)
		if err != nil {
			return nil, err
		}

		// return error if phone number already in use
		if len(patients) > 0 {
			return nil, constants.PhoneNumberAlreadyExistError
		}

		// calculate days since fever if needed
		if patient.FeverStartDate == "" {
			patient.FeverContDay = 0
		} else {
			t, err := utility.DateStringToTime(patient.FeverStartDate)
			if err != nil {
				return nil, err
			}
			patient.FeverContDay = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
		}

		// update days since exposure where needed
		if patient.ExposureDate == "" {
			patient.DaysSinceExposure = 0
		} else {
			t, err := utility.DateStringToTime(patient.ExposureDate)
			if err != nil {
				return nil, err
			}
			patient.DaysSinceExposure = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
		}

		// create patient
		result, err := m.patientDAO.Create(ctx, patient)
		if err != nil {
			return result, err
		}

		// add into declaration report if patient is symptomatic / asymptomatic / confirmed but not admitted
		if patient.Status == constants.Symptomatic || patient.Status == constants.Asymptomatic || patient.Status == constants.ConfirmedButNotAdmitted {
			date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
			err = m.declarationReportDAO.UpdateByFields(ctx, date,
				"",
				constants.DeclarationStatusMap[constants.Undeclared], patient.Type)
			if err != nil {
				return nil, err
			}
		}

		// add into status report
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		err = m.patientStatusReportDAO.UpdateByFields(ctx, date,
			"",
			constants.PatientStatusMap[patient.Status], patient.Type)
		if err != nil {
			return nil, err
		}

		// add into activity if user not nil
		if user != nil {
			_, err = m.createPatientActivity(ctx, nil, patient, user)
			if err != nil {
				return nil, err
			}
		}

		return result, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.PatientAlreadyExistError
}

// GetPatient gets patient by ID
func (m *Model) GetPatient(ctx context.Context, id string, patientType int64) (*dto.Patient, error) {
	return m.patientDAO.Get(ctx, id, patientType)
}

// QueryPatients queries patients by sort, range, filter
func (m *Model) QueryPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Patient, error) {
	return m.patientDAO.Query(ctx, sort, itemsRange, filter, patientType)
}

// BatchGetPatients get patients by slice of IDs
func (m *Model) BatchGetPatients(ctx context.Context, ids []string, patientType int64) ([]*dto.Patient, error) {
	return m.patientDAO.BatchGet(ctx, ids, patientType)
}

// GetPatientsByStatus gets patients by list of status
func (m *Model) GetPatientsByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	return m.patientDAO.GetByStatus(ctx, status, sort, itemsRange, patientType)
}

// GetNoCallPatients queries patients who have declared but no yet call
func (m *Model) GetNoCallPatients(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	return m.patientDAO.QueryNoCall(ctx, from, sort, itemsRange, patientType)
}

// GetSwabPatients get patients days since swab >= 14
func (m *Model) GetSwabPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	return m.patientDAO.GetSwabPatients(ctx, sort, itemsRange, patientType)
}

// GetDeclaredByTime gets declared patients in given from timestamp
func (m *Model) GetDeclaredPatientsByTime(ctx context.Context, from int64, patientType int64) ([]*dto.Patient, error) {
	return m.patientDAO.GetDeclaredByTime(ctx, from, patientType)
}

// GetUndeclaredPatientsByTime gets undeclared patients given from timestamp
func (m *Model) GetUndeclaredPatientsByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	return m.patientDAO.GetUndeclaredByTime(ctx, from, sort, itemsRange, patientType)
}

// UpdatePatient updates patient
func (m *Model) UpdatePatient(ctx context.Context, patient *dto.Patient, patientType int64, user *dto.User) (*dto.Patient, error) {
	// check if patient exist
	p, err := m.patientDAO.Get(ctx, patient.ID, patientType)
	if err != nil {
		return nil, err
	}

	// tech debt: get old patient
	oldPatient, err := m.patientDAO.Get(ctx, patient.ID, patientType)
	if err != nil {
		return nil, err
	}

	phoneNumberChanged := false
	nameChanged := false
	needAddUndeclared := false
	needRemoveUndeclared := false
	needRemoveDeclared := false
	needAddDeclared := false
	needAddCallingStatus := false
	needRemoveCallingStatus := false
	typeChanged := false
	oldType := p.Type
	oldStatus := p.Status
	if p.PhoneNumber != patient.PhoneNumber {
		phoneNumberChanged = true
	}
	if p.Name != patient.Name {
		nameChanged = true
	}
	if (p.Status != constants.Asymptomatic && p.Status != constants.Symptomatic && p.Status != constants.ConfirmedButNotAdmitted) &&
		(patient.Status == constants.Asymptomatic || patient.Status == constants.Symptomatic || patient.Status == constants.ConfirmedButNotAdmitted) {
		today, err := utility.DateStringToTime(utility.TimeToDateString(utility.MalaysiaTime(time.Now())))
		if err != nil {
			return nil, err
		}
		if p.LastDeclared < utility.TimeToMilli(today) {
			needAddUndeclared = true
		} else {
			needAddCallingStatus = true
			needAddDeclared = true
		}
	}
	if (p.Status == constants.Asymptomatic || p.Status == constants.Symptomatic || p.Status == constants.ConfirmedButNotAdmitted) &&
		(patient.Status != constants.Asymptomatic && patient.Status != constants.Symptomatic && patient.Status != constants.ConfirmedButNotAdmitted) {
		// send completed message is patient status has changed to completed
		if patient.Status == constants.Completed {
			err = utility.SendBotNotification(p.TelegramID, constants.CompletedMessage)
			if err != nil {
				logger.Log.Warn("failed to send completed message to telegramID: " + p.TelegramID)
			}
		}

		// if declared, remove declared
		today, err := utility.DateStringToTime(utility.TimeToDateString(utility.MalaysiaTime(time.Now())))
		if err != nil {
			return nil, err
		}
		if p.LastDeclared < utility.TimeToMilli(today) {
			needRemoveUndeclared = true
		} else {
			needRemoveCallingStatus = true
			needRemoveDeclared = true
		}

	}
	if oldType != patient.Type {
		// check for special case: patient type cannot be changed from PUI to contact tracing for PUI admin
		if patientType != constants.AllPatients && oldType == constants.PUI {
			return nil, constants.UnauthorizedAccessError
		}
		typeChanged = true
		p.TypeChangeDate = utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	}

	// patch patient
	p.Name = patient.Name
	p.Status = patient.Status
	p.PhoneNumber = patient.PhoneNumber
	p.Email = patient.Email
	p.FeverContDay = patient.FeverContDay
	p.ExposureDate = patient.ExposureDate
	p.SymptomDate = patient.SymptomDate
	p.IsolationAddress = patient.IsolationAddress
	p.AlternateContact = patient.AlternateContact
	p.RegistrationNum = patient.RegistrationNum
	p.ExposureSource = patient.ExposureSource
	p.Episode = patient.Episode
	p.Remarks = patient.Remarks
	p.Localization = patient.Localization
	p.FeverStartDate = patient.FeverStartDate
	p.Type = patient.Type
	p.HomeAddress = patient.HomeAddress
	p.IsSameAddress = patient.IsSameAddress
	if phoneNumberChanged {
		p.TelegramID = ""
	}

	// calculate days since fever if needed
	if patient.FeverStartDate == "" {
		p.FeverContDay = 0
	} else {
		t, err := utility.DateStringToTime(patient.FeverStartDate)
		if err != nil {
			return nil, err
		}
		p.FeverContDay = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
	}

	// update days since exposure where needed
	if patient.ExposureDate == "" {
		p.DaysSinceExposure = 0
	} else {
		t, err := utility.DateStringToTime(patient.ExposureDate)
		if err != nil {
			return nil, err
		}
		p.DaysSinceExposure = utility.DaysElapsed(t, utility.MalaysiaTime(time.Now())) + 1
	}

	// check phone number if phone number is changed
	if phoneNumberChanged {
		// query patients by phone number
		filter := map[string]interface{}{constants.PhoneNumber: patient.PhoneNumber}
		_, patients, err := m.patientDAO.Query(ctx, nil, nil, filter, constants.AllPatients)
		if err != nil {
			return nil, err
		}

		// return error if phone number already in use
		if len(patients) > 0 {
			return nil, constants.PhoneNumberAlreadyExistError
		}
	}

	// update patient
	_, err = m.patientDAO.Update(ctx, p)
	if err != nil {
		return nil, err
	}

	// add into activity if user is not nil
	if user != nil {
		_, err = m.createPatientActivity(ctx, oldPatient, p, user)
		if err != nil {
			return nil, err
		}
	}

	// update declarations and swabs if patient name or phone number changed
	if nameChanged || phoneNumberChanged || typeChanged {
		// update declarations
		_, declarations, err := m.QueryDeclarationsByPatientID(ctx, patient.ID, constants.AllPatients)
		for _, declaration := range declarations {
			if nameChanged {
				declaration.PatientName = patient.Name
			}
			if phoneNumberChanged {
				declaration.PatientPhoneNumber = patient.PhoneNumber
			}
			if typeChanged {
				declaration.PatientType = patient.Type
			}
			_, err = m.declarationDAO.Update(ctx, declaration)
			if err != nil {
				return nil, err
			}
		}

		// update swabs
		swabs, err := m.QuerySwabsByPatientID(ctx, patient.ID, patientType)
		for _, swab := range swabs {
			if nameChanged {
				swab.PatientName = patient.Name
			}
			if phoneNumberChanged {
				swab.PatientPhoneNumber = patient.PhoneNumber
			}
			if typeChanged {
				swab.PatientType = patient.Type
			}
			_, err = m.swabDAO.Update(ctx, swab)
			if err != nil {
				return nil, err
			}
		}
	}

	if needAddUndeclared && !typeChanged {
		// add into undeclared report (new type)
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		err = m.declarationReportDAO.UpdateByFields(ctx, date,
			"",
			constants.DeclarationStatusMap[constants.Undeclared], p.Type)
		if err != nil {
			return nil, err
		}
	}

	if needAddDeclared && !typeChanged {
		// add into declared report (new type)
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		err = m.declarationReportDAO.UpdateByFields(ctx, date,
			"",
			constants.DeclarationStatusMap[constants.Declared], p.Type)
		if err != nil {
			return nil, err
		}
	}

	if needRemoveUndeclared && !typeChanged {
		// add into undeclared report (old type)
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		err = m.declarationReportDAO.UpdateByFields(ctx, date,
			constants.DeclarationStatusMap[constants.Undeclared],
			"", p.Type)
		if err != nil {
			return nil, err
		}
	}

	if needRemoveDeclared && !typeChanged {
		// add into undeclared report
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		err = m.declarationReportDAO.UpdateByFields(ctx, date,
			constants.DeclarationStatusMap[constants.Declared], "", p.Type)
		if err != nil {
			return nil, err
		}
	}

	if needAddCallingStatus && !typeChanged {
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		filter := map[string]interface{}{constants.PatientID: p.ID}
		_, decl, err := m.QueryDeclarations(ctx, &dto.SortData{
			Item:  constants.SubmittedAt,
			Order: constants.DESC,
		}, &dto.RangeData{
			From: 0,
			To:   1,
		}, filter, constants.AllPatients)
		if err != nil {
			return nil, err
		}
		if len(decl) > 0 {
			err = m.callingReportDAO.UpdateByFields(ctx, date, "", constants.CallingStatusMap[decl[0].CallingStatus], p.Type)
			if err != nil {
				return nil, err
			}
		}
	}

	if needRemoveCallingStatus && !typeChanged {
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		filter := map[string]interface{}{constants.PatientID: p.ID}
		_, decl, err := m.QueryDeclarations(ctx, &dto.SortData{
			Item:  constants.SubmittedAt,
			Order: constants.DESC,
		}, &dto.RangeData{
			From: 0,
			To:   1,
		}, filter, constants.AllPatients)
		if err != nil {
			return nil, err
		}
		if len(decl) > 0 {
			// remove from the old patient type
			err = m.callingReportDAO.UpdateByFields(ctx, date, constants.CallingStatusMap[decl[0].CallingStatus], "", p.Type)
			if err != nil {
				return nil, err
			}
		}
	}

	if oldStatus != p.Status {
		date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
		// decrease patient status from old patient type
		err = m.patientStatusReportDAO.UpdateByFields(ctx, date,
			constants.PatientStatusMap[oldStatus],
			"", oldType)
		if err != nil {
			return nil, err
		}
		// increase patient status to new patient type
		err = m.patientStatusReportDAO.UpdateByFields(ctx, date,
			"",
			constants.PatientStatusMap[p.Status], p.Type)
		if err != nil {
			return nil, err
		}
	}

	// if patient type change, take some time to sync the report counter!
	if typeChanged {
		err = m.SyncPatientReport(ctx)
		if err != nil {
			logger.Log.Error("SyncPatientReport: " + err.Error())
			return nil, err
		}
	}

	return p, nil
}

// UpdatePatients update patients
func (m *Model) UpdatePatients(ctx context.Context, patient *dto.Patient, ids []string, patientType int64, user *dto.User) ([]string, error) {
	// TODO: Support Batch Updates
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	patient.ID = ids[0]
	p, err := m.UpdatePatient(ctx, patient, patientType, user)
	if err != nil {
		return nil, err
	}

	return []string{p.ID}, nil
}

// DeletePatient deletes patient by ID
func (m *Model) DeletePatient(ctx context.Context, id string, patientType int64) (*dto.Patient, error) {
	// check if patient exist
	p, err := m.patientDAO.Get(ctx, id, patientType)
	if err != nil {
		return nil, err
	}

	// delete all declarations by id
	_, declarations, err := m.QueryDeclarationsByPatientID(ctx, id, constants.AllPatients)
	if err != nil {
		return nil, err
	}
	for _, declaration := range declarations {
		_, err = m.DeleteDeclaration(ctx, declaration.ID, constants.AllPatients)
		if err != nil {
			return nil, err
		}
	}

	// delete all swabs by id
	swabs, err := m.QuerySwabsByPatientID(ctx, id, constants.AllPatients)
	if err != nil {
		return nil, err
	}
	for _, swab := range swabs {
		_, err = m.DeleteSwab(ctx, swab.ID, constants.AllPatients)
		if err != nil {
			return nil, err
		}
	}

	// delete patients
	err = m.patientDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	// update declaration report
	date := utility.TimeToDateString(utility.MalaysiaTime(time.Now()))
	err = m.declarationReportDAO.UpdateByFields(ctx, date,
		constants.DeclarationStatusMap[constants.Undeclared], "", p.Type)
	if err != nil {
		return nil, err
	}

	// remove from status report
	err = m.patientStatusReportDAO.UpdateByFields(ctx, date,
		constants.PatientStatusMap[p.Status], "", p.Type)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// DeletePatients delete patients by IDs
func (m *Model) DeletePatients(ctx context.Context, ids []string, patientType int64) ([]string, error) {

	var deletedIDs []string
	for _, id := range ids {
		p, err := m.DeletePatient(ctx, id, patientType)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, p.ID)
	}

	return deletedIDs, nil
}

// ClientUpdatePatient updates patient from chatbot client
func (m *Model) ClientUpdatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error) {
	// check if patient exist
	p, err := m.patientDAO.Get(ctx, patient.ID, constants.AllPatients)
	if err != nil {
		return nil, err
	}

	// patch patient
	p.TelegramID = patient.TelegramID
	p.PrivacyPolicy = patient.PrivacyPolicy
	p.Consent = patient.Consent
	p.Localization = patient.Localization

	// update patient
	_, err = m.patientDAO.Update(ctx, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ClientGetUndeclaredPatientsByTime gets undeclared patients given from timestamp (with telegramID)
func (m *Model) ClientGetUndeclaredPatientsByTime(ctx context.Context, from int64) ([]*dto.Patient, error) {
	return m.patientDAO.ClientGetUndeclaredByTime(ctx, from)
}

// ClientGetPatientsByConsentTime gets patients given from and to timestamp
func (m *Model) ClientGetPatientsByConsentTime(ctx context.Context, from int64, to int64) ([]*dto.Patient, error) {
	return m.patientDAO.GetByConsentTime(ctx, from, to)
}

// createPatientActivity creates new patient activity
func (m *Model) createPatientActivity(ctx context.Context, oldPatient, newPatient *dto.Patient, user *dto.User) (*dto.Activity, error) {
	activity := &dto.Activity{
		ID:         uuid.NewV4().String(),
		UserID:     user.ID,
		UserName:   user.DisplayName,
		OldPatient: oldPatient,
		NewPatient: newPatient,
		Time:       utility.TimeToMilli(utility.MalaysiaTime(time.Now())),
		TTL:        utility.MilliToTime(time.Now().Add(time.Hour*24*constants.ActivityTTLDays).Unix()*1000 - 1000),
	}

	// create activity
	s, err := m.activityDAO.Create(ctx, activity)
	if err != nil {
		return nil, err
	}

	// return result
	return s, nil
}
