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

// DeclarationDAO ...
type DeclarationDAO struct {
	client *mongo.Client
}

// InitDeclarationDAO ...
func InitDeclarationDAO(client *mongo.Client) IDeclarationDAO {
	return &DeclarationDAO{client: client}
}

// Create creates new declaration
func (v *DeclarationDAO) Create(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Declarations)
	if _, err := collection.InsertOne(ctx, declaration); err != nil {
		return nil, err
	}
	return declaration, nil
}

// Get gets declaration by ID
func (v *DeclarationDAO) Get(ctx context.Context, id string) (*dto.Declaration, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Declarations)
	filter := bson.D{{constants.ID, id}}
	declaration := &dto.Declaration{}
	if err := collection.FindOne(ctx, filter).Decode(&declaration); err != nil {
		return nil, err
	}
	return declaration, nil
}

// Update updates declaration
func (v *DeclarationDAO) Update(ctx context.Context, declaration *dto.Declaration) (*dto.Declaration, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Declarations)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, declaration.ID}}, bson.D{
		{"$set", declaration},
	})
	if err != nil {
		return nil, err
	}
	return declaration, nil
}

// Delete deletes declaration
func (v *DeclarationDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Declarations)
	_, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}})
	if err != nil {
		return err
	}
	return nil
}

func (v *DeclarationDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Declaration, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

func (v *DeclarationDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Declaration, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Declarations)

	filter := bson.D{{
		constants.ID,
		bson.D{{
			"$in",
			ids,
		}},
	}}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var declarations []*dto.Declaration
	for cursor.Next(ctx) {
		declaration := &dto.Declaration{}
		if err = cursor.Decode(&declaration); err != nil {
			return nil, err
		}
		declarations = append(declarations, declaration)
	}

	return declarations, nil
}

func (v *DeclarationDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *DeclarationDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Declaration, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Declarations)

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

	var declarations []*dto.Declaration
	for cursor.Next(ctx) {
		declaration := &dto.Declaration{}
		if err = cursor.Decode(&declaration); err != nil {
			return 0, nil, err
		}
		declarations = append(declarations, declaration)
	}

	count := int64(len(declarations))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, declarations, nil
}

// Low level filter parser, to be extended ...
func (v *DeclarationDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
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
			} else if key == constants.ID ||
				key == constants.PatientID ||
				key == constants.PatientName ||
				key == constants.Remarks ||
				key == constants.Date ||
				key == constants.PatientPhoneNumber ||
				key == constants.DoctorRemarks ||
				key == constants.Category {
				result = append(result, bson.E{Key: key, Value: fmt.Sprint(value)})
			} else if key != constants.PatientType {
				result = append(result, bson.E{Key: key, Value: utility.SafeCastInt64(value)})
			}
		}
	}

	return result
}
