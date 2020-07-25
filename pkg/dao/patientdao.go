package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/logger"
	"comet/pkg/utility"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strconv"
	"strings"
)

// PatientDAO ...
type PatientDAO struct {
	client *mongo.Client
}

// InitPatientDAO ...
func InitPatientDAO(client *mongo.Client) IPatientDAO {
	return &PatientDAO{client: client}
}

// Create creates new patient
func (v *PatientDAO) Create(ctx context.Context, patient *dto.Patient) (*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)
	if _, err := collection.InsertOne(ctx, patient); err != nil {
		return nil, err
	}
	return patient, nil
}

// Get gets patient by ID and type
func (v *PatientDAO) Get(ctx context.Context, id string, patientType int64) (*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)

	filter := bson.D{{constants.ID, id}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	patient := &dto.Patient{}
	if err := collection.FindOne(ctx, filter).Decode(&patient); err != nil {
		return nil, err
	}

	return patient, nil
}

// BatchGet gets patients by slice of IDs and type
func (v *PatientDAO) BatchGet(ctx context.Context, ids []string, patientType int64) ([]*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)

	filter := bson.D{{
		constants.ID,
		bson.D{{
			"$in",
			ids,
		}},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var patients []*dto.Patient
	for cursor.Next(ctx) {
		patient := &dto.Patient{}
		if err = cursor.Decode(&patient); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

// Query queries patients by sort, range, type and a filter to match any number of fields
func (v *PatientDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Patient, error) {

	f := v.parseFilter(filter)
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		f = append(f, bson.E{Key: constants.Type, Value: patientType})
	}

	return v.query(ctx, sort, itemsRange, f)
}

// GetByStatus gets patients by statuses and type
func (v *PatientDAO) GetByStatus(ctx context.Context, status []int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {

	filter := bson.D{{
		constants.Status,
		bson.D{{
			"$in",
			status}},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	return v.query(ctx, sort, itemsRange, filter)
}

// GetSwabPatients gets swab patients of the specified type >= 14 days (ONLY RETURN PATIENTS WITH STATUS 1, 2, 3)
func (v *PatientDAO) GetSwabPatients(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	statusBSON := bson.A{constants.Symptomatic, constants.Asymptomatic, constants.ConfirmedButNotAdmitted}
	filter := bson.D{{
		"$and",
		bson.A{
			bson.D{{
				constants.DaysSinceSwab,
				bson.D{{
					"$gte",
					14,
				}},
			}},
			bson.D{{
				constants.Status,
				bson.D{{
					"$in",
					statusBSON,
				}},
			}},
		},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	return v.query(ctx, sort, itemsRange, filter)
}

// GetDeclaredByTime gets declared patients of the specified type in given from timestamp
func (v *PatientDAO) GetDeclaredByTime(ctx context.Context, from int64, patientType int64) ([]*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)

	statusBSON := bson.A{constants.Symptomatic, constants.Asymptomatic, constants.ConfirmedButNotAdmitted}
	filter := bson.D{{
		"$and",
		bson.A{
			bson.D{{
				constants.LastDeclared,
				bson.D{{
					"$gte",
					from,
				}},
			}},
			bson.D{{
				constants.Status,
				bson.D{{
					"$in",
					statusBSON,
				}},
			}},
		},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var patients []*dto.Patient
	for cursor.Next(ctx) {
		patient := &dto.Patient{}
		if err = cursor.Decode(&patient); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

// GetUndeclaredByTime gets undeclared patients of the specified type given from timestamp
func (v *PatientDAO) GetUndeclaredByTime(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	statusBSON := bson.A{constants.Symptomatic, constants.Asymptomatic, constants.ConfirmedButNotAdmitted}
	filter := bson.D{{
		"$and",
		bson.A{
			bson.D{{
				constants.LastDeclared,
				bson.D{{
					"$lt",
					from,
				}},
			}},
			bson.D{{
				constants.Status,
				bson.D{{
					"$in",
					statusBSON,
				}},
			}},
		},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	return v.query(ctx, sort, itemsRange, filter)
}

// GetByConsentTime gets patients given from and to consent timestamp
func (v *PatientDAO) GetByConsentTime(ctx context.Context, from int64, to int64) ([]*dto.Patient, error) {
	statusBSON := bson.A{constants.Asymptomatic, constants.Symptomatic, constants.ConfirmedButNotAdmitted}
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)
	cursor, err := collection.Find(ctx, bson.D{{
		"$and",
		bson.A{
			bson.D{{
				constants.Consent,
				bson.D{{
					"$gte",
					from,
				}},
			}},
			bson.D{{
				constants.Consent,
				bson.D{{
					"$lte",
					to,
				}},
			}},
			bson.D{{
				constants.Status,
				bson.D{{
					"$in",
					statusBSON,
				}},
			}},
		},
	}})
	if err != nil {
		return nil, err
	}

	var patients []*dto.Patient
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		patient := &dto.Patient{}
		if err = cursor.Decode(&patient); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

// Update updates patient
func (v *PatientDAO) Update(ctx context.Context, patient *dto.Patient) (*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, patient.ID}}, bson.D{
		{"$set", patient},
	})
	if err != nil {
		return nil, err
	}
	return patient, nil
}

// Delete deletes patient by ID
func (v *PatientDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes patients by IDs
func (v *PatientDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
	var deletedIDs []string
	for _, id := range ids {
		err := v.Delete(ctx, id)
		if err != nil {
			return nil, err
		}
		deletedIDs = append(deletedIDs, id)
	}
	return deletedIDs, nil
}

// ClientGetUndeclaredByTime gets undeclared patients given from timestamp (ONLY RETURN PATIENTS WITH STATUS 1, 2, 3 AND WITH TELEGRAM_ID)
func (v *PatientDAO) ClientGetUndeclaredByTime(ctx context.Context, from int64) ([]*dto.Patient, error) {
	statusBSON := bson.A{constants.Asymptomatic, constants.Symptomatic, constants.ConfirmedButNotAdmitted}
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)
	cursor, err := collection.Find(ctx, bson.D{{
		"$and",
		bson.A{
			bson.D{{
				constants.LastDeclared,
				bson.D{{
					"$lt",
					from,
				}},
			}},
			bson.D{{
				constants.Status,
				bson.D{{
					"$in",
					statusBSON,
				}},
			}},
			bson.D{{
				constants.TelegramID,
				bson.D{{
					"$not",
					bson.D{{"$eq", ""}},
				}},
			}},
		},
	}})
	if err != nil {
		return nil, err
	}

	var patients []*dto.Patient
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		patient := &dto.Patient{}
		if err = cursor.Decode(&patient); err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}

	return patients, nil
}

// QueryNoCall queries patients who have declared but no yet call
func (v *PatientDAO) QueryNoCall(ctx context.Context, from int64, sort *dto.SortData, itemsRange *dto.RangeData, patientType int64) (int64, []*dto.Patient, error) {
	statusBSON := bson.A{constants.Symptomatic, constants.Asymptomatic, constants.ConfirmedButNotAdmitted}
	filter := bson.D{{
		"$and",
		bson.A{
			bson.D{{
				constants.LastDeclared,
				bson.D{{
					"$gte",
					from,
				}},
			}},
			bson.D{{
				constants.Status,
				bson.D{{
					"$in",
					statusBSON,
				}},
			}},
			bson.D{{
				constants.Declarations + "." + constants.SubmittedAt,
				bson.D{{
					"$gte",
					from,
				}},
			}},
			bson.D{{
				constants.Declarations + "." + constants.CallingStatus, constants.NoYetCall,
			}},
		},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.Type, Value: patientType})
	}

	return v.aggregate(ctx, sort, itemsRange, filter)
}

// aggregate is a generic mongodb find helper method (based on filters in declarations collection)
func (v *PatientDAO) aggregate(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)

	result := mongo.Pipeline{}

	// stage 1: lookup
	result = append(result, bson.D{{
		"$lookup", bson.M{
			"from":         constants.Declarations,
			"localField":   constants.ID,
			"foreignField": constants.PatientID,
			"as":           "declarations",
		},
	}})

	// stage 2: unwind
	result = append(result, bson.D{{
		"$unwind", bson.M{
			"path": "$declarations",
		},
	}})

	// stage 3: match
	result = append(result, bson.D{{
		"$match", filter,
	}})

	// stage 4: sort
	if sort != nil {
		order := 1
		if sort.Order == constants.DESC {
			order = -1
		}
		result = append(result, bson.D{{
			"$sort", bson.M{
				sort.Item: order,
			},
		}})
	}

	// stage 5: set range
	if itemsRange != nil {
		result = append(result, bson.D{{
			"$facet", bson.M{
				"patients": mongo.Pipeline{
					bson.D{{
						"$skip", int64(itemsRange.From),
					}},
					bson.D{{
						"$limit", int64(itemsRange.To + 1 - itemsRange.From),
					}},
				},
				"count": mongo.Pipeline{
					bson.D{{
						"$count", "count",
					}},
				},
			},
		}})
	}

	cursor, err := collection.Aggregate(ctx, result)
	if err != nil {
		logger.Log.Error(err.Error())
		return 0, nil, err
	}
	defer cursor.Close(ctx)

	// if facet is used, decode into PatientsQuery
	if itemsRange != nil {
		for cursor.Next(ctx) {
			result := &dto.PatientsQuery{}
			if err = cursor.Decode(&result); err != nil {
				logger.Log.Error(err.Error())
				return 0, nil, err
			}
			if len(result.Count) < 1 || result.Count[0] == nil {
				return 0, nil, nil
			}
			return result.Count[0].Count, result.Patients, nil
		}
	}

	// else, decode into patient array
	var patients []*dto.Patient
	for cursor.Next(ctx) {
		patient := &dto.Patient{}
		if err = cursor.Decode(&patient); err != nil {
			return 0, nil, err
		}
		patients = append(patients, patient)
	}
	count := int64(len(patients))

	return count, patients, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *PatientDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Patient, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Patients)

	findOptions := options.Find()
	// set range
	if itemsRange != nil {
		findOptions.SetSkip(int64(itemsRange.From))
		findOptions.SetLimit(int64(itemsRange.To + 1 - itemsRange.From))
	}

	// set sorter
	if sort != nil {
		order := 1
		if sort.Order == constants.DESC {
			order = -1
		}
		findOptions.SetSort(bson.D{{sort.Item, order}})
	}

	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return 0, nil, err
	}
	defer cursor.Close(ctx)

	var patients []*dto.Patient
	for cursor.Next(ctx) {
		patient := &dto.Patient{}
		if err = cursor.Decode(&patient); err != nil {
			return 0, nil, err
		}
		patients = append(patients, patient)
	}

	count := int64(len(patients))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, patients, nil
}

// Low level filter parser, to be extended ...
func (v *PatientDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
						bson.M{
							constants.ID: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.Name: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.PhoneNumber: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
					},
				})
			} else if key == constants.LastDeclareResult {
				if v, err := strconv.ParseBool(strings.ToLower(fmt.Sprint(value))); err != nil {
					result = append(result, bson.E{Key: key, Value: v})
				}

			} else if key == constants.Status ||
				key == constants.LastDeclared ||
				key == constants.SwabCount ||
				key == constants.Episode ||
				key == constants.DaysSinceExposure ||
				key == constants.DaysSinceSwab ||
				key == constants.FeverContDay ||
				key == constants.Localization ||
				key == constants.Consent ||
				key == constants.PrivacyPolicy ||
				key == constants.CallingStatus {
				result = append(result, bson.E{Key: key, Value: utility.SafeCastInt64(value)})

			} else if key != constants.Type { // prevent access-control by-passing and nasty bugs
				result = append(result, bson.E{Key: key, Value: fmt.Sprint(value)})
			}
		}
	}

	return result
}
