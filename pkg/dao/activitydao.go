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

// ActivityDAO ...
type ActivityDAO struct {
	client *mongo.Client
}

// InitActivityDAO ...
func InitActivityDAO(client *mongo.Client) IActivityDAO {
	return &ActivityDAO{client: client}
}

// Create creates new activity
func (v *ActivityDAO) Create(ctx context.Context, activity *dto.Activity) (*dto.Activity, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Activities)
	if _, err := collection.InsertOne(ctx, activity); err != nil {
		return nil, err
	}
	return activity, nil
}

// Get gets activity by ID
func (v *ActivityDAO) Get(ctx context.Context, id string) (*dto.Activity, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Activities)
	filter := bson.D{{constants.ID, id}}

	activity := &dto.Activity{}
	if err := collection.FindOne(ctx, filter).Decode(&activity); err != nil {
		return nil, err
	}

	return activity, nil
}

// Query queries activities by sort, range, filter
func (v *ActivityDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Activity, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *ActivityDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Activity, error) {
	collection := v.client.Database(constants.Cosmos).Collection(constants.Activities)

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

	var activities []*dto.Activity
	for cursor.Next(ctx) {
		activity := &dto.Activity{}
		if err = cursor.Decode(&activity); err != nil {
			return 0, nil, err
		}
		activities = append(activities, activity)
	}

	count := int64(len(activities))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, activities, nil
}

// Low level filter parser, to be extended ...
func (v *ActivityDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
						bson.M{
							constants.UserName: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.NewPatient + "." + constants.ID: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.NewSwab + "." + constants.ID: bson.M{
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
