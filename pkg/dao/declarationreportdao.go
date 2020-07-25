package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeclarationReportDAO ...
type DeclarationReportDAO struct {
	client *mongo.Client
}

func InitDeclarationReportDAO(client *mongo.Client) IDeclarationReport {
	return &DeclarationReportDAO{client: client}
}

// Create creates declaration report
func (v *DeclarationReportDAO) Create(ctx context.Context, date string, declarationReport *dto.DeclarationReport) error {
	declarationReport.Date = date
	collection := v.client.Database(constants.Cosmos).Collection(constants.DeclarationReports)

	report := &dto.DeclarationReport{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		// if not found, create new one
		if _, err := collection.InsertOne(ctx, declarationReport); err != nil {
			return err
		}
		return nil
	}

	// else update
	_, err := collection.UpdateOne(ctx, bson.D{{constants.Date, date}}, bson.D{
		{"$set", declarationReport},
	})
	if err != nil {
		return err
	}
	return nil
}

// UpdateByFields decrease decrementField by 1 and increase incrementField by 1 (based on patient type)
func (v *DeclarationReportDAO) UpdateByFields(ctx context.Context, date, decrementField, incrementField string, patientType int64) error {

	// derive field names based on patient type
	def := constants.PatientFieldMap[patientType] + decrementField
	inf := constants.PatientFieldMap[patientType] + incrementField

	collection := v.client.Database(constants.Cosmos).Collection(constants.DeclarationReports)
	doc := &dto.DeclarationReport{}
	// check if report exist
	err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(doc)
	if err != nil {
		doc = &dto.DeclarationReport{Date: date}
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

// Get gets declaration report by date
func (v *DeclarationReportDAO) Get(ctx context.Context, date string) (*dto.DeclarationReport, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.DeclarationReports)
	report := &dto.DeclarationReport{}
	if err := collection.FindOne(ctx, bson.D{{constants.Date, date}}).Decode(report); err != nil {
		report = &dto.DeclarationReport{Date: date}
	}
	return report, nil
}

func (v *DeclarationReportDAO) BatchGet(ctx context.Context, dates []string) ([]*dto.DeclarationReport, error) {
	var reports []*dto.DeclarationReport

	for _, date := range dates {
		report, err := v.Get(ctx, date)
		if err != nil {
			return nil, err
		}
		reports = append(reports, report)
	}

	return reports, nil
}
