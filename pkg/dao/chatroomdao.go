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

// ChatRoomDAO ...
type ChatRoomDAO struct {
	client *mongo.Client
}

// InitChatRoomDAO ...
func InitChatRoomDAO(client *mongo.Client) IChatRoomDAO {
	return &ChatRoomDAO{client: client}
}

// Create creates new chatRoom
func (v *ChatRoomDAO) Create(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatRooms)
	if _, err := collection.InsertOne(ctx, chatRoom); err != nil {
		return nil, err
	}
	return chatRoom, nil
}

// Get gets chatRoom by ID
func (v *ChatRoomDAO) Get(ctx context.Context, id string) (*dto.ChatRoom, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatRooms)
	filter := bson.D{{constants.ID, id}}

	chatRoom := &dto.ChatRoom{}
	if err := collection.FindOne(ctx, filter).Decode(&chatRoom); err != nil {
		return nil, err
	}

	return chatRoom, nil
}

// BatchGet gets chatRooms by slice of IDs
func (v *ChatRoomDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.ChatRoom, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatRooms)

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

	var chatRooms []*dto.ChatRoom
	for cursor.Next(ctx) {
		chatRoom := &dto.ChatRoom{}
		if err = cursor.Decode(&chatRoom); err != nil {
			return nil, err
		}
		chatRooms = append(chatRooms, chatRoom)
	}

	return chatRooms, nil
}

// Query queries chatRooms by sort, range, filter
func (v *ChatRoomDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.ChatRoom, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes chatRoom by ID
func (v *ChatRoomDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatRooms)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes chatRooms by IDs
func (v *ChatRoomDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates chatRoom
func (v *ChatRoomDAO) Update(ctx context.Context, chatRoom *dto.ChatRoom) (*dto.ChatRoom, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatRooms)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, chatRoom.ID}}, bson.D{
		{"$set", chatRoom},
	})
	if err != nil {
		return nil, err
	}
	return chatRoom, nil
}

// Query By User
func (v *ChatRoomDAO) QueryByUsers(ctx context.Context, users []string) ([]*dto.ChatRoom, error) {
	filter := bson.D{{
		constants.ParticipantIDs,
		users,
	}}
	_, rooms, err := v.query(ctx, &dto.SortData{
		Item:  constants.Timestamp,
		Order: constants.DESC,
	}, nil, filter)
	return rooms, err
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *ChatRoomDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.ChatRoom, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.ChatRooms)

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

	var chatRooms []*dto.ChatRoom
	for cursor.Next(ctx) {
		chatRoom := &dto.ChatRoom{}
		if err = cursor.Decode(&chatRoom); err != nil {
			return 0, nil, err
		}
		chatRooms = append(chatRooms, chatRoom)
	}

	count := int64(len(chatRooms))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, chatRooms, nil
}

// Low level filter parser, to be extended ...
func (v *ChatRoomDAO) parseFilter(filter map[string]interface{}) bson.D {
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
					},
				})
			} else if key != constants.Type { // prevent access-control by-passing and nasty bugs
				result = append(result, bson.E{Key: key, Value: fmt.Sprint(value)})
			}
		}
	}

	return result
}
