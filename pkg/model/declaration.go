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

	m.computeResult(ctx, declaration, p)

	// create declaration
	resp, err := m.declarationDAO.Create(ctx, declaration)
	if err != nil {
		return nil, err
	}

	// update patient
	_, err = m.patientDAO.Update(ctx, p)
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

// QueryDeclarationsByCategories query declarations by categories
func (m *Model) QueryDeclarationsByCategories(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, id string, categories []string) (int64, []*dto.Declaration, error) {
	return m.declarationDAO.QueryByCategories(ctx, sort, itemsRange, id, categories)
}

// QueryDeclarationsByPatientID ...
func (m *Model) QueryDeclarationsByPatientID(ctx context.Context, id string) (int64, []*dto.Declaration, error) {
	filter := map[string]interface{}{constants.PatientID: id}
	total, declarations, err := m.declarationDAO.Query(ctx, nil, nil, filter)
	return total, declarations, err
}

func (m *Model) computeResult(ctx context.Context, declaration *dto.Declaration, patient *dto.Patient) {
	// skip if no result
	if len(declaration.Result) == 0 {
		return
	}

	// get declaration category
	q, err := m.questionDAO.Get(ctx, declaration.Result[0].ID)
	if err != nil {
		return
	}
	declaration.Category = q.Category

	if declaration.Category == constants.DASS {
		for _, result := range declaration.Result {
			q, err := m.questionDAO.Get(ctx, result.ID)
			if err != nil {
				continue
			}
			switch q.Type {
			case "stress":
				declaration.Stress += result.Score
			case "anxiety":
				declaration.Anxiety += result.Score
			case "depression":
				declaration.Depression += result.Score
			}
			result.Category = q.Category
			result.Type = q.Type
			result.Content = q.Content
			declaration.Score += result.Score
		}
		declaration.Score *= 2
		declaration.Stress *= 2
		declaration.Anxiety *= 2
		declaration.Depression *= 2
		patient.LastDassTime = declaration.SubmittedAt
		patient.LastDassResult = declaration.Score

		// Depression
		if declaration.Depression >= 28 {
			declaration.DepressionStatus = constants.DeclarationExtremelySevere
		} else if declaration.Depression >= 21 {
			declaration.DepressionStatus = constants.DeclarationSevere
		} else if declaration.Depression >= 14 {
			declaration.DepressionStatus = constants.DeclarationModerate
		} else if declaration.Depression >= 10 {
			declaration.DepressionStatus = constants.DeclarationMild
		} else {
			declaration.DepressionStatus = constants.DeclarationNormal
		}

		// Anxiety
		if declaration.Anxiety >= 20 {
			declaration.AnxietyStatus = constants.DeclarationExtremelySevere
		} else if declaration.Anxiety >= 15 {
			declaration.AnxietyStatus = constants.DeclarationSevere
		} else if declaration.Anxiety >= 10 {
			declaration.AnxietyStatus = constants.DeclarationModerate
		} else if declaration.Anxiety >= 8 {
			declaration.AnxietyStatus = constants.DeclarationMild
		} else {
			declaration.AnxietyStatus = constants.DeclarationNormal
		}

		// Stress
		if declaration.Stress >= 34 {
			declaration.StressStatus = constants.DeclarationExtremelySevere
		} else if declaration.Stress >= 26 {
			declaration.StressStatus = constants.DeclarationSevere
		} else if declaration.Stress >= 19 {
			declaration.StressStatus = constants.DeclarationModerate
		} else if declaration.Stress >= 15 {
			declaration.StressStatus = constants.DeclarationMild
		} else {
			declaration.StressStatus = constants.DeclarationNormal
		}

		// Patient mental status
		if declaration.Stress > declaration.Depression && declaration.Stress > declaration.Anxiety {
			patient.MentalStatus = constants.Stress
		} else if declaration.Anxiety > declaration.Depression && declaration.Anxiety > declaration.Stress {
			patient.MentalStatus = constants.Anxiety
		} else {
			patient.MentalStatus = constants.Depression
		}

		patient.DepressionStatus = declaration.DepressionStatus
		patient.AnxietyStatus = declaration.AnxietyStatus
		patient.StressStatus = declaration.StressStatus
	} else {
		score := int64(0)
		for _, result := range declaration.Result {
			q, err := m.questionDAO.Get(ctx, result.ID)
			if err != nil {
				continue
			}
			result.Content = q.Content
			result.Category = q.Category
			result.Type = q.Type
			score += result.Score
		}
		declaration.Score = score

		if declaration.Category == constants.IESR {
			patient.LastIesrTime = declaration.SubmittedAt
			patient.LastIesrResult = score

			if score >= 37 {
				declaration.PtsdStatus = constants.DeclarationSevere
				patient.MentalStatus = constants.PTSD
			} else if score >= 24 {
				declaration.PtsdStatus = constants.DeclarationModerate
				patient.MentalStatus = constants.PTSD
			} else {
				declaration.PtsdStatus = constants.DeclarationNormal
			}
			patient.PtsdStatus = declaration.PtsdStatus
		}
	}
}
