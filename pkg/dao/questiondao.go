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

// QuestionDAO ...
type QuestionDAO struct {
	client *mongo.Client
}

// InitQuestionDAO ...
func InitQuestionDAO(client *mongo.Client) IQuestionDAO {
	return &QuestionDAO{client: client}
}

// Create creates new question
func (v *QuestionDAO) Create(ctx context.Context, question *dto.Question) (*dto.Question, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Questions)
	if _, err := collection.InsertOne(ctx, question); err != nil {
		return nil, err
	}
	return question, nil
}

// Get gets question by ID
func (v *QuestionDAO) Get(ctx context.Context, id string) (*dto.Question, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Questions)
	filter := bson.D{{constants.ID, id}}

	question := &dto.Question{}
	if err := collection.FindOne(ctx, filter).Decode(&question); err != nil {
		return nil, err
	}

	return question, nil
}

// BatchGet gets questions by slice of IDs
func (v *QuestionDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.Question, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Questions)

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

	var questions []*dto.Question
	for cursor.Next(ctx) {
		question := &dto.Question{}
		if err = cursor.Decode(&question); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}

// Query queries questions by sort, range, filter
func (v *QuestionDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter map[string]interface{}) (int64, []*dto.Question, error) {
	f := v.parseFilter(filter)
	return v.query(ctx, sort, itemsRange, f)
}

// Delete deletes question by ID
func (v *QuestionDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Questions)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes questions by IDs
func (v *QuestionDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates question
func (v *QuestionDAO) Update(ctx context.Context, question *dto.Question) (*dto.Question, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Questions)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, question.ID}}, bson.D{
		{"$set", question},
	})
	if err != nil {
		return nil, err
	}
	return question, nil
}

// query is a generic mongodb find helper method
// IMPORTANT SHIT: this query uses FIND. It will never return err codes.Unknown! Only FINDONE will return codes.Unknown
// DO NOT check for codes.Unknown to see if there's result. It will never hit! Use length instead please.
func (v *QuestionDAO) query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter bson.D) (int64, []*dto.Question, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Questions)

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

	var questions []*dto.Question
	for cursor.Next(ctx) {
		question := &dto.Question{}
		if err = cursor.Decode(&question); err != nil {
			return 0, nil, err
		}
		questions = append(questions, question)
	}

	count := int64(len(questions))
	if itemsRange != nil { // count only if client query with range, else default to length of query results
		if count, err = collection.CountDocuments(ctx, filter); err != nil {
			return 0, nil, err
		}
	}

	return count, questions, nil
}

// Low level filter parser, to be extended ...
func (v *QuestionDAO) parseFilter(filter map[string]interface{}) bson.D {
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
			} else if key == constants.Status {
				result = append(result, bson.E{Key: key, Value: utility.SafeCastInt64(value)})
			} else if key != constants.PatientType {
				result = append(result, bson.E{Key: key, Value: fmt.Sprint(value)})
			}
		}
	}

	return result
}
