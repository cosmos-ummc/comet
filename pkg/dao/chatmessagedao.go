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

// ChatMessageDAO ...
type ChatMessageDAO struct {
	client *mongo.Client
}

// InitChatMessageDAO ...
func InitChatMessageDAO(client *mongo.Client) IChatMessageDAO {
	return &ChatMessageDAO{client: client}
}

// Create creates new chatMessage
func (v *ChatMessageDAO) Create(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatMessages)
	if _, err := collection.InsertOne(ctx, chatMessage); err != nil {
		return nil, err
	}
	return chatMessage, nil
}

// Get gets chatMessage by ID
func (v *ChatMessageDAO) Get(ctx context.Context, id string) (*dto.ChatMessage, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatMessages)
	filter := bson.D{{constants.ID, id}}

	chatMessage := &dto.ChatMessage{}
	if err := collection.FindOne(ctx, filter).Decode(&chatMessage); err != nil {
		return nil, err
	}

	return chatMessage, nil
}

// BatchGet gets chatMessages by slice of IDs
func (v *ChatMessageDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.ChatMessage, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatMessages)

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

	var chatMessages []*dto.ChatMessage
	for cursor.Next(ctx) {
		chatMessage := &dto.ChatMessage{}
		if err = cursor.Decode(&chatMessage); err != nil {
			return nil, err
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	return chatMessages, nil
}

// Query queries chatMessages by sort, range, filter
func (v *ChatMessageDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatMessage, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes chatMessage by ID
func (v *ChatMessageDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatMessages)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes chatMessages by IDs
func (v *ChatMessageDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates chatMessage
func (v *ChatMessageDAO) Update(ctx context.Context, chatMessage *dto.ChatMessage) (*dto.ChatMessage, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatMessages)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, chatMessage.ID}}, bson.D{
		{"$set", chatMessage},
	})
	if err != nil {
		return nil, err
	}
	return chatMessage, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *ChatMessageDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.ChatMessage, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatMessages)

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

	var chatMessages []*dto.ChatMessage
	for cursor.Next(ctx) {
		chatMessage := &dto.ChatMessage{}
		if err = cursor.Decode(&chatMessage); err != nil {
			return 0, nil, err
		}
		chatMessages = append(chatMessages, chatMessage)
	}

	count := int64(len(chatMessages))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, chatMessages, nil
}

// Low level filter parser, to be extended ...
func (v *ChatMessageDAO) parseFilter(filter map[string]interface{}) bson.D {
	// cannot be nil
	result := bson.D{}

	if filter != nil && len(filter) > 0 {
		for key, value := range filter {
			if key == "q" {
				result = append(result, bson.E{
					Key: "$or",
					Value: bson.A{
						bson.M{
							constants.Category: bson.M{
								"$regex":   fmt.Sprintf("%s.*", value),
								"$options": "i",
							},
						},
						bson.M{
							constants.Contents: bson.M{
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
