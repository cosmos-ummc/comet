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

// MeetingDAO ...
type MeetingDAO struct {
	client *mongo.Client
}

// InitMeetingDAO ...
func InitMeetingDAO(client *mongo.Client) IMeetingDAO {
	return &MeetingDAO{client: client}
}

// Create creates new meeting
func (v *MeetingDAO) Create(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meetings)
	if _, err := collection.InsertOne(ctx, meeting); err != nil {
		return nil, err
	}
	return meeting, nil
}

// Get gets meeting by ID
func (v *MeetingDAO) Get(ctx context.Context, id string) (*dto.Meeting, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meetings)
	filter := bson.D{{constants.ID, id}}

	meeting := &dto.Meeting{}
	if err := collection.FindOne(ctx, filter).Decode(&meeting); err != nil {
		return nil, err
	}

	return meeting, nil
}

// BatchGet gets meetings by slice of IDs
func (v *MeetingDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Meeting, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meetings)

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

	var meetings []*dto.Meeting
	for cursor.Next(ctx) {
		meeting := &dto.Meeting{}
		if err = cursor.Decode(&meeting); err != nil {
			return nil, err
		}
		meetings = append(meetings, meeting)
	}

	return meetings, nil
}

// Query queries meetings by sort, range, filter
func (v *MeetingDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Meeting, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes meeting by ID
func (v *MeetingDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meetings)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes meetings by IDs
func (v *MeetingDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates meeting
func (v *MeetingDAO) Update(ctx context.Context, meeting *dto.Meeting) (*dto.Meeting, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meetings)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, meeting.ID}}, bson.D{
		{"$set", meeting},
	})
	if err != nil {
		return nil, err
	}
	return meeting, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *MeetingDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Meeting, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Meetings)

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

	var meetings []*dto.Meeting
	for cursor.Next(ctx) {
		meeting := &dto.Meeting{}
		if err = cursor.Decode(&meeting); err != nil {
			return 0, nil, err
		}
		meetings = append(meetings, meeting)
	}

	count := int64(len(meetings))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, meetings, nil
}

// Low level filter parser, to be extended ...
func (v *MeetingDAO) parseFilter(filter map[string]interface{}) bson.D {
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
					},
				})
			}
		}
	}

	return result
}
