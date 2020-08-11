package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ClientCreateDeclaration creates new declaration or updates existing declarations
func (m *Model) ClientCreateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error) {

	// check if patient exist, put patient name
	p, err := m.GetPatient(ctx, declaration.PatientID)
	if err != nil {
		return nil, constants.PatientNotFoundError
	}
	declaration.PatientName = p.Name
	declaration.PatientPhoneNumber = p.PhoneNumber

	// TODO: compute declaration result
	declaration.Status = constants.DeclarationMild
	declaration.Score = 100
	p.LastDassTime = declaration.SubmittedAt
	p.LastDassResult = declaration.Score

	// create declaration
	resp, err := m.declarationDAO.Create(ctx, declaration)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// CreateDeclaration creates new declaration
func (m *Model) CreateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error) {

	// check if declaration exist
	_, err := m.declarationDAO.Get(ctx, declaration.ID)

	// only can create declaration if not found
	if err != nil && status.Code(err) == codes.Unknown {

		// check if patient exist, put patient name
		p, err := m.GetPatient(ctx, declaration.PatientID)
		if err != nil {
			return nil, constants.PatientNotFoundError
		}
		declaration.PatientName = p.Name
		declaration.PatientPhoneNumber = p.PhoneNumber

		// create declaration
		resp, err := m.declarationDAO.Create(ctx, declaration)
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
func (m *Model) GetDeclaration(ctx context.Context, declarationID string) (*dto.Declaration, error) {
	return m.declarationDAO.Get(ctx, declarationID)
}

// UpdateDeclaration updates declaration
func (m *Model) UpdateDeclaration(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error) {
	// check if declaration exist
	d, err := m.declarationDAO.Get(ctx, declaration.ID)
	if err != nil {
		return nil, err
	}

	// patch declaration
	d.Score = declaration.Score
	d.Status = declaration.Status
	d.DoctorRemarks = declaration.DoctorRemarks

	// update declaration
	_, err = m.declarationDAO.Update(ctx, d)
	if err != nil {
		return nil, err
	}

	return d, err
}

// UpdateDeclarations update declarations
func (m *Model) UpdateDeclarations(ctx context.Context, declaration *dto.Declaration, ids []string) ([]string, error) {
	// TODO: Support Batch Updates
	if len(ids) > 1 {
		return nil, constants.OperationUnsupportedError
	}
	declaration.ID = ids[0]
	d, err := m.UpdateDeclaration(ctx, declaration)
	if err != nil {
		return nil, err
	}
	return []string{d.ID}, nil
}

// DeleteDeclaration deletes declaration
func (m *Model) DeleteDeclaration(ctx context.Context, declarationID string) (*dto.Declaration, error) {
	d, err := m.GetDeclaration(ctx, declarationID)
	if err != nil {
		return nil, err
	}

	err = m.declarationDAO.Delete(ctx, declarationID)
	if err != nil {
		return nil, err
	}

	return d, nil
}

// BatchGetDeclarations get declarations
func (m *Model) BatchGetDeclarations(ctx context.Context, declarationIDs []string) ([]*dto.Declaration, error) {
	return m.declarationDAO.BatchGet(ctx, declarationIDs)
}

// DeleteDeclarations delete declarations
func (m *Model) DeleteDeclarations(ctx context.Context, declarationIDs []string) ([]string, error) {

	var deletedIDs []string
	for _, id := range declarationIDs {
		d, err := m.DeleteDeclaration(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, d.ID)
	}

	return deletedIDs, nil
}

// QueryDeclarations query declarations
func (m *Model) QueryDeclarations(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Declaration, error) {
	return m.declarationDAO.Query(ctx, sort, itemsRange, filter)
}

// QueryDeclarationsByPatientID ...
func (m *Model) QueryDeclarationsByPatientID(ctx context.Context, id string) (int64, []*dto.Declaration, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	total, declarations, err := m.declarationDAO.Query(ctx, nil, nil, filter)
	return total, declarations, err
}
