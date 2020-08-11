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

// GameDAO ...
type GameDAO struct {
	client *mongo.Client
}

// InitGameDAO ...
func InitGameDAO(client *mongo.Client) IGameDAO {
	return &GameDAO{client: client}
}

// Create creates new game
func (v *GameDAO) Create(ctx context.Context, game *dto.Game) (*dto.Game, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Games)
	if _, err := collection.InsertOne(ctx, game); err != nil {
		return nil, err
	}
	return game, nil
}

// Get gets game by ID
func (v *GameDAO) Get(ctx context.Context, id string) (*dto.Game, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Games)
	filter := bson.D{{constants.ID, id}}

	game := &dto.Game{}
	if err := collection.FindOne(ctx, filter).Decode(&game); err != nil {
		return nil, err
	}

	return game, nil
}

// BatchGet gets games by slice of IDs
func (v *GameDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Game, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Games)

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

	var games []*dto.Game
	for cursor.Next(ctx) {
		game := &dto.Game{}
		if err = cursor.Decode(&game); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	return games, nil
}

// Query queries games by sort, range, filter
func (v *GameDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Game, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes game by ID
func (v *GameDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Games)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes games by IDs
func (v *GameDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates game
func (v *GameDAO) Update(ctx context.Context, game *dto.Game) (*dto.Game, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Games)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, game.ID}}, bson.D{
		{"$set", game},
	})
	if err != nil {
		return nil, err
	}
	return game, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *GameDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Game, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Games)

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

	var games []*dto.Game
	for cursor.Next(ctx) {
		game := &dto.Game{}
		if err = cursor.Decode(&game); err != nil {
			return 0, nil, err
		}
		games = append(games, game)
	}

	count := int64(len(games))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, games, nil
}

// Low level filter parser, to be extended ...
func (v *GameDAO) parseFilter(filter map[string]interface{}) bson.D {
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
