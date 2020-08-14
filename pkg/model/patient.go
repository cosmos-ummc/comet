package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/utility"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreatePatient creates new patient
func (m *Model) CreatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error) {

	// check if patient exist
	_, err := m.patientDAO.Get(ctx, patient.ID)

	// only can create patient if not found
	if err != nil && status.Code(err) == codes.Unknown {

		// query patients by phone number
		filter := map[string]interface{}{constants.PhoneNumber: patient.PhoneNumber}
		_, patients, err := m.patientDAO.Query(ctx, nil, nil, filter)
		if err != nil {
			return nil, err
		}

		// return error if phone number already in use
		if len(patients) > 0 {
			return nil, constants.PhoneNumberAlreadyExistError
		}

		// create patient
		result, err := m.patientDAO.Create(ctx, patient)
		if err != nil {
			return result, err
		}

		return result, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, constants.PatientAlreadyExistError
}

// GetPatient gets patient by ID
func (m *Model) GetPatient(ctx context.Context, id string) (*dto.Patient, error) {
	return m.patientDAO.Get(ctx, id)
}

// QueryPatients queries patients by sort, range, filter
func (m *Model) QueryPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Patient, error) {
	return m.patientDAO.Query(ctx, sort, itemsRange, filter)
}

// BatchGetPatients get patients by slice of IDs
func (m *Model) BatchGetPatients(ctx context.Context, ids []string) ([]*dto.Patient, error) {
	return m.patientDAO.BatchGet(ctx, ids)
}

// GetPatientsByStatus gets patients by list of status
func (m *Model) GetPatientsByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error) {
	return m.patientDAO.GetByStatus(ctx, status, sort, itemsRange)
}

// GetUndeclaredPatientsByTime gets undeclared patients given from timestamp
func (m *Model) GetUndeclaredPatientsByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData) (int64, []*dto.Patient, error) {
	return m.patientDAO.GetUndeclaredByTime(ctx, from, sort, itemsRange)
}

// UpdatePatient updates patient
func (m *Model) UpdatePatient(ctx context.Context, patient *dto.Patient) (*dto.Patient, error) {
	// check if patient exist
	p, err := m.patientDAO.Get(ctx, patient.ID)
	if err != nil {
		return nil, err
	}

	phoneNumberChanged := false
	if p.PhoneNumber != patient.PhoneNumber {
		phoneNumberChanged = true
	}

	// patch patient
	oldHasCompleted := p.HasCompleted
	p.Name = patient.Name
	p.PhoneNumber = patient.PhoneNumber
	p.Email = patient.Email
	p.IsolationAddress = patient.IsolationAddress
	p.Remarks = patient.Remarks
	p.HomeAddress = patient.HomeAddress
	p.DaySinceMonitoring = patient.DaySinceMonitoring
	p.HasCompleted = patient.HasCompleted
	p.MentalStatus = patient.MentalStatus
	p.Type = patient.Type
	p.SwabDate = patient.SwabDate
	p.SwabResult = patient.SwabResult
	p.TutorialStage = patient.TutorialStage
	p.TelegramID = patient.TelegramID

	if p.HasCompleted && !oldHasCompleted {
		_ = utility.SendBotNotification(p.TelegramID, constants.CompletedMessage)
	}

	if phoneNumberChanged {
		p.TelegramID = ""
	}

	// check phone number if phone number is changed
	if phoneNumberChanged {
		// query patients by phone number
		filter := map[string]interface{}{constants.PhoneNumber: patient.PhoneNumber}
		_, patients, err := m.patientDAO.Query(ctx, nil, nil, filter)
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

	return p, nil
}

func (m *Model) RemindPatients(ctx context.Context) error {
	_, patients, err := m.QueryPatients(ctx, nil, nil, nil)
	if err != nil {
		return err
	}
	for _, p := range patients {
		if p.TelegramID != "" && !p.HasCompleted {
			if p.DaySinceMonitoring % 7 == 0 {
				_ = utility.SendBotNotification(p.TelegramID, constants.ReminderMessage)
			} else {
				_ = utility.SendBotNotification(p.TelegramID, constants.ReminderMessageDaily)
			}
		}
	}
	return nil
}

// UpdatePatients update patients
func (m *Model) UpdatePatients(ctx context.Context, patient *dto.Patient, ids []string) ([]string, error) {
	// TODO: Support Batch Updates
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	patient.ID = ids[0]
	p, err := m.UpdatePatient(ctx, patient)
	if err != nil {
		return nil, err
	}

	return []string{p.ID}, nil
}

// DeletePatient deletes patient by ID
func (m *Model) DeletePatient(ctx context.Context, id string) (*dto.Patient, error) {
	// check if patient exist
	p, err := m.patientDAO.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete all declarations by id
	_, declarations, err := m.QueryDeclarationsByPatientID(ctx, id)
	if err != nil {
		return nil, err
	}
	for _, declaration := range declarations {
		_, err = m.DeleteDeclaration(ctx, declaration.ID)
		if err != nil {
			return nil, err
		}
	}

	// delete patients
	err = m.patientDAO.Delete(ctx, id)
	if err != nil {
		return nil, err
	}

	// delete user
	_, err = m.DeleteUser(ctx, p.UserID)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// DeletePatients delete patients by IDs
func (m *Model) DeletePatients(ctx context.Context, ids []string) ([]string, error) {

	var deletedIDs []string
	for _, id := range ids {
		p, err := m.DeletePatient(ctx, id)
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
	p, err := m.patientDAO.Get(ctx, patient.ID)
	if err != nil {
		return nil, err
	}

	// patch patient
	p.TelegramID = patient.TelegramID
	p.PrivacyPolicy = patient.PrivacyPolicy
	p.Consent = patient.Consent
	p.DaySinceMonitoring = 1

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

// VerifyPatientComplete verifies if patient has completed monitoring
func (m *Model) VerifyPatientComplete(ctx context.Context, id string, force bool) (bool, error) {
	p, err := m.patientDAO.Get(ctx, id)
	if err != nil {
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			return false, err
		}

		return false, err
	}

	// check if patient is PUI or PUS
	if p.Type != constants.PUI && p.Type != constants.PUS && !force {
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	// check if patient has negative swab
	if p.SwabResult != constants.SwabNegative && !force {
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	// check if days since monitoring > 14
	if p.DaySinceMonitoring <= 14 && !force {
		_, err = m.patientDAO.Update(ctx, p)
		if err != nil {
			return false, err
		}

		return false, nil
	}

	// update patient completed
	p.HasCompleted = true
	_, err = m.patientDAO.Update(ctx, p)
	if err != nil {
		return false, err
	}

	// TODO: Send congratulation message

	return true, nil
}
