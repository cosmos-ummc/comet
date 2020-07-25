package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// PatientStatusReportDAO ...
type PatientStatusReportDAO struct {
	client *mongo.Client
}

func InitPatientStatusReportDAO(client *mongo.Client) IPatientStatusReport {
	return &PatientStatusReportDAO{client: client}
}

// Create creates PatientStatus report
func (v *PatientStatusReportDAO) Create(ctx context.Context, date string, patientStatusReport *dto.PatientStatusReport) error {
	patientStatusReport.Date = date
	collection := v.client.Database(constants.Cosmos).Collection(constants.PatientStatusReports)

	report := &dto.PatientStatusReport{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		// if not found, create new one
		if _, err := collection.InsertOne(ctx, patientStatusReport); err != nil {
			return err
		}
		return nil
	}

	// else update
	_, err := collection.UpdateOne(ctx, bson.D{{constants.Date, date}}, bson.D{
		{"$set", patientStatusReport},
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateByFields decrease decrementField by 1 and increase incrementField by 1 (based on patient type)
func (v *PatientStatusReportDAO) UpdateByFields(ctx context.Context, date, decrementField, incrementField string, patientType int64) error {

	// derive field names based on patient type
	def := constants.PatientFieldMap[patientType] + decrementField
	inf := constants.PatientFieldMap[patientType] + incrementField

	collection := v.client.Database(constants.Cosmos).Collection(constants.PatientStatusReports)
	doc := &dto.PatientStatusReport{}
	err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(doc)
	if err != nil {
		doc = &dto.PatientStatusReport{Date: date}
		// create one if not exist
		err = v.Create(ctx, date, doc)
		if err != nil {
			return err
		}
	}
	if incrementField != "" {
		// increment counter
		_, err = collection.UpdateOne(ctx, bson.D{{constants.Date, date}}, bson.D{
			{"$inc", bson.D{
				{inf, 1},
			}},
		})
		if err != nil {
			return err
		}
	}
	if decrementField != "" {
		// decrement counter
		_, err = collection.UpdateOne(ctx, bson.D{{constants.Date, date}}, bson.D{
			{"$inc", bson.D{
				{def, -1},
			}},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

// Get gets PatientStatus report by date
func (v *PatientStatusReportDAO) Get(ctx context.Context, date string) (*dto.PatientStatusReport, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.PatientStatusReports)
	report := &dto.PatientStatusReport{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		report = &dto.PatientStatusReport{Date: date}
	}
	return report, nil
}

func (v *PatientStatusReportDAO) BatchGet(ctx context.Context, dates []string) ([]*dto.PatientStatusReport, error) {
	var reports []*dto.PatientStatusReport

	for _, date := range dates {
		report, err := v.Get(ctx, date)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
