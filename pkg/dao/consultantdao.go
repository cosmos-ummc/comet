package dao

import (
	"comet/pkg/constants"
	"comet/pkg/dto"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConsultantDAO ...
type ConsultantDAO struct {
	client *mongo.Client
}

// InitConsultantDAO ...
func InitConsultantDAO(client *mongo.Client) IConsultantDAO {
	return &ConsultantDAO{client: client}
}

// Create creates new consultant
func (v *ConsultantDAO) Create(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Consultants)
	if _, err := collection.InsertOne(ctx, consultant); err != nil {
		return nil, err
	}
	return consultant, nil
}

// Get gets consultant by ID
func (v *ConsultantDAO) Get(ctx context.Context, id string) (*dto.Consultant, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Consultants)
	filter := bson.D{{constants.ID, id}}

	consultant := &dto.Consultant{}
	if err := collection.FindOne(ctx, filter).Decode(&consultant); err != nil {
		return nil, err
	}

	return consultant, nil
}

// BatchGet gets consultants by slice of IDs
func (v *ConsultantDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Consultant, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Consultants)

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

	var consultants []*dto.Consultant
	for cursor.Next(ctx) {
		consultant := &dto.Consultant{}
		if err = cursor.Decode(&consultant); err != nil {
			return nil, err
		}
		consultants = append(consultants, consultant)
	}

	return consultants, nil
}

// Query queries consultants by sort, range, filter
func (v *ConsultantDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Consultant, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes consultant by ID
func (v *ConsultantDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Consultants)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes consultants by IDs
func (v *ConsultantDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates consultant
func (v *ConsultantDAO) Update(ctx context.Context, consultant *dto.Consultant) (*dto.Consultant, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Consultants)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, consultant.ID}}, bson.D{
		{"$set", consultant},
	})
	if err != nil {
		return nil, err
	}
	return consultant, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *ConsultantDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Consultant, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Consultants)

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

	var consultants []*dto.Consultant
	for cursor.Next(ctx) {
		consultant := &dto.Consultant{}
		if err = cursor.Decode(&consultant); err != nil {
			return 0, nil, err
		}
		consultants = append(consultants, consultant)
	}

	count := int64(len(consultants))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, consultants, nil
}

// Low level filter parser, to be extended ...
func (v *ConsultantDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
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
						bson.M{
							constants.Email: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
					},
				})
			}
		}
	}

	return result
}
