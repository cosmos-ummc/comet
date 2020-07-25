package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Calling Report
type CallingReportDAO struct {
	client *mongo.Client
}

func InitCallingReportDAO(client *mongo.Client) ICallingReport {
	return &CallingReportDAO{client: client}
}

// Create creates calling report
func (v *CallingReportDAO) Create(ctx context.Context, date string, callingReport *dto.CallingReport) error {
	callingReport.Date = date
	collection := v.client.Database(constants.Cosmos).Collection(constants.CallingReports)

	report := &dto.CallingReport{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		// if not found, create new one
		if _, err := collection.InsertOne(ctx, callingReport); err != nil {
			return err
		}
		return nil
	}

	// else update
	_, err := collection.UpdateOne(ctx, bson.D{{constants.Date, date}}, bson.D{
		{"$set", callingReport},
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateByFields decrease decrementField by 1 and increase incrementField by 1 (based on patient type)
func (v *CallingReportDAO) UpdateByFields(ctx context.Context, date, decrementField, incrementField string, patientType int64) error {

	// derive field names based on patient type
	def := constants.PatientFieldMap[patientType] + decrementField
	inf := constants.PatientFieldMap[patientType] + incrementField

	collection := v.client.Database(constants.Cosmos).Collection(constants.CallingReports)
	doc := &dto.CallingReport{}
	// check if report exist
	err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(doc)
	if err != nil {
		doc = &dto.CallingReport{Date: date}
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

// Get gets report by date
func (v *CallingReportDAO) Get(ctx context.Context, date string) (*dto.CallingReport, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.CallingReports)
	report := &dto.CallingReport{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		report = &dto.CallingReport{Date: date}
	}

	return report, nil
}

func (v *CallingReportDAO) BatchGet(ctx context.Context, dates []string) ([]*dto.CallingReport, error) {
	var reports []*dto.CallingReport

	for _, date := range dates {
		report, err := v.Get(ctx, date)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
