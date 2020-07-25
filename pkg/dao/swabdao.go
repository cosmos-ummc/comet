package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"comet/pkg/utility"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SwabDAO ...
type SwabDAO struct {
	client *mongo.Client
}

// InitSwabDAO ...
func InitSwabDAO(client *mongo.Client) ISwabDAO {
	return &SwabDAO{client: client}
}

// Create creates new swab
func (v *SwabDAO) Create(ctx context.Context, swab *dto.Swab) (*dto.Swab, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Swabs)
	if _, err := collection.InsertOne(ctx, swab); err != nil {
		return nil, err
	}
	return swab, nil
}

// Get gets swab by ID
func (v *SwabDAO) Get(ctx context.Context, id string, patientType int64) (*dto.Swab, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Swabs)
	filter := bson.D{{constants.ID, id}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.PatientType, Value: patientType})
	}

	swab := &dto.Swab{}
	if err := collection.FindOne(ctx, filter).Decode(&swab); err != nil {
		return nil, err
	}

	return swab, nil
}

// BatchGet gets swabs by slice of IDs
func (v *SwabDAO) BatchGet(ctx context.Context, ids []string, patientType int64) ([]*dto.Swab, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Swabs)

	filter := bson.D{{
		constants.ID,
		bson.D{{
			"$in",
			ids,
		}},
	}}
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		filter = append(filter, bson.E{Key: constants.PatientType, Value: patientType})
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var swabs []*dto.Swab
	for cursor.Next(ctx) {
		swab := &dto.Swab{}
		if err = cursor.Decode(&swab); err != nil {
			return nil, err
		}
		swabs = append(swabs, swab)
	}

	return swabs, nil
}

// Query queries swabs by sort, range, filter
func (v *SwabDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}, patientType int64) (int64, []*dto.Swab, error) {
	f := v.parseFilter(filter)
	if patientType == constants.PUI || patientType == constants.ContactTracing {
		f = append(f, bson.E{Key: constants.PatientType, Value: patientType})
	}

	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes swab by ID
func (v *SwabDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Swabs)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes swabs by IDs
func (v *SwabDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates swab
func (v *SwabDAO) Update(ctx context.Context, swab *dto.Swab) (*dto.Swab, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Swabs)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, swab.ID}}, bson.D{
		{"$set", swab},
	})
	if err != nil {
		return nil, err
	}
	return swab, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *SwabDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Swab, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Swabs)

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

	var swabs []*dto.Swab
	for cursor.Next(ctx) {
		swab := &dto.Swab{}
		if err = cursor.Decode(&swab); err != nil {
			return 0, nil, err
		}
		swabs = append(swabs, swab)
	}

	count := int64(len(swabs))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, swabs, nil
}

// Low level filter parser, to be extended ...
func (v *SwabDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
						bson.M{
							constants.Location: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.PatientID: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.PatientName: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.PatientPhoneNumber: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
					},
				})
			} else if key == constants.Status {
				result = append(result, bson.E{Key: key, Value: utility.SafeCastInt64(value)})
			} else if key != constants.PatientType {
				result = append(result, bson.E{Key: key, Value: fmt.Sprint(value)})
			}
		}
	}

	return result
}
