package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
)

// GetReport gets report
func (m *Model) GetReport(ctx context.Context, id string) (*dto.Report, error) {
	report := &dto.Report{}
	if id == "1" {
		// get general report

		// get all patients
		_, patients, err := m.QueryPatients(ctx, nil, nil, nil)
		if err != nil {
			return nil, err
		}

		for _, patient := range patients {
			// Depression
			switch patient.DepressionStatus {
			case constants.DeclarationNormal:
				report.DepressionNormal += 1
			case constants.DeclarationMild:
				report.DepressionMild += 1
			case constants.DeclarationModerate:
				report.DepressionModerate += 1
			case constants.DeclarationSevere:
				report.DepressionSevere += 1
			case constants.DeclarationExtremelySevere:
				report.DepressionExtreme += 1
			}

			// Anxiety
			switch patient.AnxietyStatus {
			case constants.DeclarationNormal:
				report.AnxietyNormal += 1
			case constants.DeclarationMild:
				report.AnxietyMild += 1
			case constants.DeclarationModerate:
				report.AnxietyModerate += 1
			case constants.DeclarationSevere:
				report.AnxietySevere += 1
			case constants.DeclarationExtremelySevere:
				report.AnxietyExtreme += 1
			}

			// Stress
			switch patient.StressStatus {
			case constants.DeclarationNormal:
				report.StressNormal += 1
			case constants.DeclarationMild:
				report.StressMild += 1
			case constants.DeclarationModerate:
				report.StressModerate += 1
			case constants.DeclarationSevere:
				report.StressSevere += 1
			case constants.DeclarationExtremelySevere:
				report.StressExtreme += 1
			}

			// PTSD
			switch patient.PtsdStatus {
			case constants.DeclarationNormal:
				report.PtsdNormal += 1
			case constants.DeclarationSevere:
				report.PtsdSevere += 1
			}

			// get declarations (DASS)
			total, declarations, err := m.QueryDeclarations(ctx, &dto.SortData{
				Item:  constants.SubmittedAt,
				Order: constants.ASC,
			}, nil, map[string]interface{}{
				constants.PatientID: patient.ID,
				constants.Category:  constants.DASS,
			})
			if err != nil {
				return nil, err
			}
			if total >= 2 {
				// Depression
				if declarations[0].DepressionStatus == constants.DeclarationSevere || declarations[0].DepressionStatus == constants.DeclarationExtremelySevere {
					report.DepressionCount1 += 1
				}
				if declarations[1].DepressionStatus == constants.DeclarationSevere || declarations[1].DepressionStatus == constants.DeclarationExtremelySevere {
					report.DepressionCount2 += 1
				}

				// Stress
				if declarations[0].StressStatus == constants.DeclarationSevere || declarations[0].StressStatus == constants.DeclarationExtremelySevere {
					report.StressCount1 += 1
				}
				if declarations[1].StressStatus == constants.DeclarationSevere || declarations[1].StressStatus == constants.DeclarationExtremelySevere {
					report.StressCount2 += 1
				}

				// Anxiety
				if declarations[0].AnxietyStatus == constants.DeclarationSevere || declarations[0].AnxietyStatus == constants.DeclarationExtremelySevere {
					report.AnxietyCount1 += 1
				}
				if declarations[1].AnxietyStatus == constants.DeclarationSevere || declarations[1].AnxietyStatus == constants.DeclarationExtremelySevere {
					report.AnxietyCount2 += 1
				}
			}

			// get declarations (IES-R)
			total, declarations, err = m.QueryDeclarations(ctx, &dto.SortData{
				Item:  constants.SubmittedAt,
				Order: constants.ASC,
			}, nil, map[string]interface{}{
				constants.PatientID: patient.ID,
				constants.Category:  constants.IESR,
			})
			if err != nil {
				return nil, err
			}
			if total >= 2 {
				// Anxiety
				if declarations[0].PtsdStatus == constants.DeclarationSevere || declarations[0].PtsdStatus == constants.DeclarationExtremelySevere {
					report.PtsdCount1 += 1
				}
				if declarations[1].PtsdStatus == constants.DeclarationSevere || declarations[1].PtsdStatus == constants.DeclarationExtremelySevere {
					report.PtsdCount2 += 1
				}
			}
		}
	}

	return report, nil
}
