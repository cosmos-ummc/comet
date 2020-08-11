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

// MeditationDAO ...
type MeditationDAO struct {
	client *mongo.Client
}

// InitMeditationDAO ...
func InitMeditationDAO(client *mongo.Client) IMeditationDAO {
	return &MeditationDAO{client: client}
}

// Create creates new meditation
func (v *MeditationDAO) Create(ctx context.Context, meditation *dto.Meditation) (*dto.Meditation, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meditations)
	if _, err := collection.InsertOne(ctx, meditation); err != nil {
		return nil, err
	}
	return meditation, nil
}

// Get gets meditation by ID
func (v *MeditationDAO) Get(ctx context.Context, id string) (*dto.Meditation, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meditations)
	filter := bson.D{{constants.ID, id}}

	meditation := &dto.Meditation{}
	if err := collection.FindOne(ctx, filter).Decode(&meditation); err != nil {
		return nil, err
	}

	return meditation, nil
}

// BatchGet gets meditations by slice of IDs
func (v *MeditationDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Meditation, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meditations)

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

	var meditations []*dto.Meditation
	for cursor.Next(ctx) {
		meditation := &dto.Meditation{}
		if err = cursor.Decode(&meditation); err != nil {
			return nil, err
		}
		meditations = append(meditations, meditation)
	}

	return meditations, nil
}

// Query queries meditations by sort, range, filter
func (v *MeditationDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meditation, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes meditation by ID
func (v *MeditationDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meditations)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes meditations by IDs
func (v *MeditationDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates meditation
func (v *MeditationDAO) Update(ctx context.Context, meditation *dto.Meditation) (*dto.Meditation, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meditations)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, meditation.ID}}, bson.D{
		{"$set", meditation},
	})
	if err != nil {
		return nil, err
	}
	return meditation, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *MeditationDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Meditation, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meditations)

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

	var meditations []*dto.Meditation
	for cursor.Next(ctx) {
		meditation := &dto.Meditation{}
		if err = cursor.Decode(&meditation); err != nil {
			return 0, nil, err
		}
		meditations = append(meditations, meditation)
	}

	count := int64(len(meditations))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, meditations, nil
}

// Low level filter parser, to be extended ...
func (v *MeditationDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
						bson.M{
							constants.Link: bson.M{
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
