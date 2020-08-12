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

// TipDAO ...
type TipDAO struct {
	client *mongo.Client
}

// InitTipDAO ...
func InitTipDAO(client *mongo.Client) ITipDAO {
	return &TipDAO{client: client}
}

// Create creates new tip
func (v *TipDAO) Create(ctx context.Context, tip *dto.Tip) (*dto.Tip, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Tips)
	if _, err := collection.InsertOne(ctx, tip); err != nil {
		return nil, err
	}
	return tip, nil
}

// Get gets tip by ID
func (v *TipDAO) Get(ctx context.Context, id string) (*dto.Tip, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Tips)
	filter := bson.D{{constants.ID, id}}

	tip := &dto.Tip{}
	if err := collection.FindOne(ctx, filter).Decode(&tip); err != nil {
		return nil, err
	}

	return tip, nil
}

// BatchGet gets tips by slice of IDs
func (v *TipDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Tip, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Tips)

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

	var tips []*dto.Tip
	for cursor.Next(ctx) {
		tip := &dto.Tip{}
		if err = cursor.Decode(&tip); err != nil {
			return nil, err
		}
		tips = append(tips, tip)
	}

	return tips, nil
}

// Query queries tips by sort, range, filter
func (v *TipDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Tip, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes tip by ID
func (v *TipDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Tips)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes tips by IDs
func (v *TipDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates tip
func (v *TipDAO) Update(ctx context.Context, tip *dto.Tip) (*dto.Tip, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Tips)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, tip.ID}}, bson.D{
		{"$set", tip},
	})
	if err != nil {
		return nil, err
	}
	return tip, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *TipDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Tip, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Tips)

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

	var tips []*dto.Tip
	for cursor.Next(ctx) {
		tip := &dto.Tip{}
		if err = cursor.Decode(&tip); err != nil {
			return 0, nil, err
		}
		tips = append(tips, tip)
	}

	count := int64(len(tips))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, tips, nil
}

// Low level filter parser, to be extended ...
func (v *TipDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
						bson.M{
							constants.Title: bson.M{
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
