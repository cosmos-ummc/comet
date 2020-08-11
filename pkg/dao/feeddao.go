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

// FeedDAO ...
type FeedDAO struct {
	client *mongo.Client
}

// InitFeedDAO ...
func InitFeedDAO(client *mongo.Client) IFeedDAO {
	return &FeedDAO{client: client}
}

// Create creates new feed
func (v *FeedDAO) Create(ctx context.Context, feed *dto.Feed) (*dto.Feed, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Feeds)
	if _, err := collection.InsertOne(ctx, feed); err != nil {
		return nil, err
	}
	return feed, nil
}

// Get gets feed by ID
func (v *FeedDAO) Get(ctx context.Context, id string) (*dto.Feed, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Feeds)
	filter := bson.D{{constants.ID, id}}

	feed := &dto.Feed{}
	if err := collection.FindOne(ctx, filter).Decode(&feed); err != nil {
		return nil, err
	}

	return feed, nil
}

// BatchGet gets feeds by slice of IDs
func (v *FeedDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Feed, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Feeds)

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

	var feeds []*dto.Feed
	for cursor.Next(ctx) {
		feed := &dto.Feed{}
		if err = cursor.Decode(&feed); err != nil {
			return nil, err
		}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}

// Query queries feeds by sort, range, filter
func (v *FeedDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Feed, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes feed by ID
func (v *FeedDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Feeds)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes feeds by IDs
func (v *FeedDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates feed
func (v *FeedDAO) Update(ctx context.Context, feed *dto.Feed) (*dto.Feed, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Feeds)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, feed.ID}}, bson.D{
		{"$set", feed},
	})
	if err != nil {
		return nil, err
	}
	return feed, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *FeedDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Feed, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Feeds)

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

	var feeds []*dto.Feed
	for cursor.Next(ctx) {
		feed := &dto.Feed{}
		if err = cursor.Decode(&feed); err != nil {
			return 0, nil, err
		}
		feeds = append(feeds, feed)
	}

	count := int64(len(feeds))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, feeds, nil
}

// Low level filter parser, to be extended ...
func (v *FeedDAO) parseFilter(filter map[string]interface{}) bson.D {
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
