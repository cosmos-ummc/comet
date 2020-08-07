package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Calling Report
type ReportDAO struct {
	client *mongo.Client
}

func InitReportDAO(client *mongo.Client) IReportDAO {
	return &ReportDAO{client: client}
}

// Create creates report
func (v *ReportDAO) Create(ctx context.Context, date string, report *dto.Report) error {
	report.Date = date
	collection := v.client.Database(constants.Mhpss).Collection(constants.Reports)

	r := &dto.Report{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(r); err != nil {
		// if not found, create new one
		if _, err := collection.InsertOne(ctx, report); err != nil {
			return err
		}
		return nil
	}

	// else update
	_, err := collection.UpdateOne(ctx, bson.D{{constants.Date, date}}, bson.D{
		{"$set", report},
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateByFields decrease decrementField by 1 and increase incrementField by 1
func (v *ReportDAO) UpdateByFields(ctx context.Context, date, decrementField, incrementField string) error {

	// derive field names based on patient type
	def := decrementField
	inf := incrementField

	collection := v.client.Database(constants.Mhpss).Collection(constants.Reports)
	doc := &dto.Report{}
	// check if report exist
	err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(doc)
	if err != nil {
		doc = &dto.Report{Date: date}
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
func (v *ReportDAO) Get(ctx context.Context, date string) (*dto.Report, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Reports)
	report := &dto.Report{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		report = &dto.Report{Date: date}
	}

	return report, nil
}

func (v *ReportDAO) BatchGet(ctx context.Context, dates []string) ([]*dto.Report, error) {
	var reports []*dto.Report

	for _, date := range dates {
		report, err := v.Get(ctx, date)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
