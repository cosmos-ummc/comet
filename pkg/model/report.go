package model

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/utility"
	"context"
)

// GetReport gets report
func (m *Model) GetReport(ctx context.Context, id string) (*dto.Report, error) {
	report := &dto.Report{}

	if id == "1" {
		// get general report
		report.DepressionCounts = []int64{0, 0}
		report.StressCounts = []int64{0, 0}
		report.AnxietyCounts = []int64{0, 0}
		report.PtsdCounts = []int64{0, 0}

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
					report.DepressionCounts[0] += 1
				}
				if declarations[len(declarations)-1].DepressionStatus == constants.DeclarationSevere || declarations[len(declarations)-1].DepressionStatus == constants.DeclarationExtremelySevere {
					report.DepressionCounts[1] += 1
				}

				// Stress
				if declarations[0].StressStatus == constants.DeclarationSevere || declarations[0].StressStatus == constants.DeclarationExtremelySevere {
					report.StressCounts[0] += 1
				}
				if declarations[len(declarations)-1].StressStatus == constants.DeclarationSevere || declarations[len(declarations)-1].StressStatus == constants.DeclarationExtremelySevere {
					report.StressCounts[1] += 1
				}

				// Anxiety
				if declarations[0].AnxietyStatus == constants.DeclarationSevere || declarations[0].AnxietyStatus == constants.DeclarationExtremelySevere {
					report.AnxietyCounts[0] += 1
				}
				if declarations[len(declarations)-1].AnxietyStatus == constants.DeclarationSevere || declarations[len(declarations)-1].AnxietyStatus == constants.DeclarationExtremelySevere {
					report.AnxietyCounts[1] += 1
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
				// PTSD
				if declarations[0].PtsdStatus == constants.DeclarationSevere || declarations[0].PtsdStatus == constants.DeclarationExtremelySevere {
					report.PtsdCounts[0] += 1
				}
				if declarations[len(declarations)-1].PtsdStatus == constants.DeclarationSevere || declarations[len(declarations)-1].PtsdStatus == constants.DeclarationExtremelySevere {
					report.PtsdCounts[1] += 1
				}
			}
		}
	} else {
		// personalized report
		report.DepressionCounts = []int64{}
		report.StressCounts = []int64{}
		report.AnxietyCounts = []int64{}
		report.PtsdCounts = []int64{}
		report.DepressionStatuses = []int64{}
		report.StressStatuses = []int64{}
		report.AnxietyStatuses = []int64{}
		report.PtsdStatuses = []int64{}

		// get declarations (DASS)
		_, declarations, err := m.QueryDeclarations(ctx, &dto.SortData{
			Item:  constants.SubmittedAt,
			Order: constants.ASC,
		}, nil, map[string]interface{}{
			constants.PatientID: id,
			constants.Category:  constants.DASS,
		})
		if err != nil {
			return nil, err
		}
		for _, declaration := range declarations {
			report.DepressionCounts = append(report.DepressionCounts, declaration.Depression)
			report.DepressionStatuses = append(report.DepressionStatuses, utility.DepressionScoreToStatus(declaration.Depression))
			report.StressCounts = append(report.StressCounts, declaration.Stress)
			report.StressStatuses = append(report.StressStatuses, utility.StressScoreToStatus(declaration.Stress))
			report.AnxietyCounts = append(report.AnxietyCounts, declaration.Anxiety)
			report.AnxietyStatuses = append(report.AnxietyStatuses, utility.AnxietyScoreToStatus(declaration.Anxiety))
		}

		// get declarations (IES-R)
		_, declarations, err = m.QueryDeclarations(ctx, &dto.SortData{
			Item:  constants.SubmittedAt,
			Order: constants.ASC,
		}, nil, map[string]interface{}{
			constants.PatientID: id,
			constants.Category:  constants.IESR,
		})
		if err != nil {
			return nil, err
		}
		for _, declaration := range declarations {
			report.PtsdCounts = append(report.PtsdCounts, declaration.Score)
			report.PtsdStatuses = append(report.PtsdStatuses, utility.PtsdScoreToStatus(declaration.Score))
		}

		// get declarations (Daily)
		_, declarations, err = m.QueryDeclarations(ctx, &dto.SortData{
			Item:  constants.SubmittedAt,
			Order: constants.ASC,
		}, nil, map[string]interface{}{
			constants.PatientID: id,
			constants.Category:  constants.Daily,
		})
		if err != nil {
			return nil, err
		}
		for _, declaration := range declarations {
			report.DailyCounts = append(report.PtsdCounts, declaration.Score)
			// Todo: passing marks for daily report
			report.DailyStatuses = append(report.PtsdStatuses, utility.PtsdScoreToStatus(declaration.Score))
		}
	}

	return report, nil
}
