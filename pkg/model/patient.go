package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
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
	nameChanged := false
	if p.PhoneNumber != patient.PhoneNumber {
		phoneNumberChanged = true
	}
	if p.Name != patient.Name {
		nameChanged = true
	}

	// patch patient
	p.Name = patient.Name
	p.Status = patient.Status
	p.PhoneNumber = patient.PhoneNumber
	p.Email = patient.Email
	p.IsolationAddress = patient.IsolationAddress
	p.Remarks = patient.Remarks
	p.HomeAddress = patient.HomeAddress

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

	// update declarations and swabs if patient name or phone number changed
	if nameChanged || phoneNumberChanged {
		// update declarations
		_, declarations, err := m.QueryDeclarationsByPatientID(ctx, patient.ID)
		for _, declaration := range declarations {
			if nameChanged {
				declaration.PatientName = patient.Name
			}
			if phoneNumberChanged {
				declaration.PatientPhoneNumber = patient.PhoneNumber
			}
			_, err = m.declarationDAO.Update(ctx, declaration)
			if err != nil {
				return nil, err
			}
		}
	}

	return p, nil
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
