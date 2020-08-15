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

// UserDAO ...
type UserDAO struct {
	client *mongo.Client
}

// InitUserDAO ...
func InitUserDAO(client *mongo.Client) IUserDAO {
	return &UserDAO{client: client}
}

// Create creates new user
func (v *UserDAO) Create(ctx context.Context, user *dto.User) (*dto.User, error) {

	// hash password
	var err error
	user.Password, err = utility.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// create user
	collection := v.client.Database(constants.Mhpss).Collection(constants.Users)
	if _, err := collection.InsertOne(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Get gets user by ID
func (v *UserDAO) Get(ctx context.Context, id string) (*dto.User, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Users)
	user := &dto.User{}
	if err := collection.FindOne(ctx, bson.D{{constants.ID, id}}).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

// BatchGet gets users by slice of IDs
func (v *UserDAO) BatchGet(ctx context.Context, ids []string) ([]*dto.User, error) {
	var users []*dto.User
	for _, id := range ids {
		user, err := v.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Query queries users by sort, range, filter
func (v *UserDAO) Query(ctx context.Context, sort *dto.SortData, itemsRange *dto.RangeData, filter *dto.FilterData, superuserOnly bool) (int64, []*dto.User, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Users)

	var cursor *mongo.Cursor
	var err error
	var count int64
	var users []*dto.User

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

	// set filter
	if filter != nil {
		// special case: if filter item is primary key we can directly call Get
		if filter.Item == "id" {
			d, err := v.Get(ctx, filter.Value)
			if err != nil {
				return 0, nil, err
			}
			return 1, []*dto.User{d}, nil
		}

		// else: do query filter
		if filter.Item == "q" {
			query := bson.M{
				"$or": bson.A{
					bson.M{
						constants.Name: bson.M{
							"$regex":   fmt.Sprintf("%s.*", filter.Value),
							"$options": "i",
						},
					},
					bson.M{
						constants.PhoneNumber: bson.M{
							"$regex":   fmt.Sprintf("%s.*", filter.Value),
							"$options": "i",
						},
					},
					bson.M{
						constants.Email: bson.M{
							"$regex":   fmt.Sprintf("%s.*", filter.Value),
							"$options": "i",
						},
					},
				},
			}

			// if superuser only
			if superuserOnly {
				newQuery := bson.D{{
					"$and",
					bson.A{
						bson.D{{
							constants.Role,
							bson.D{{
								"$in",
								[]string{constants.Consultant, constants.Admin, constants.Superuser},
							}},
						}},
						query,
					},
				}}
				cursor, err = collection.Find(ctx, newQuery, findOptions)
				if err != nil {
					return 0, nil, err
				}
				count, err = collection.CountDocuments(ctx, newQuery)
				if err != nil {
					return 0, nil, err
				}
			} else {
				cursor, err = collection.Find(ctx, query, findOptions)
				if err != nil {
					return 0, nil, err
				}
				count, err = collection.CountDocuments(ctx, query)
				if err != nil {
					return 0, nil, err
				}
			}
		} else {
			// if superuser only
			if superuserOnly {
				newQuery := bson.D{{
					"$and",
					bson.A{
						bson.D{{
							constants.Role, constants.Superuser,
						}},
						bson.D{
							{filter.Item, filter.Value},
						},
					},
				}}
				cursor, err = collection.Find(ctx, newQuery, findOptions)
				if err != nil {
					return 0, nil, err
				}
				count, err = collection.CountDocuments(ctx, newQuery)
				if err != nil {
					return 0, nil, err
				}
			} else {
				cursor, err = collection.Find(ctx, bson.D{
					{filter.Item, filter.Value},
				}, findOptions)
				if err != nil {
					return 0, nil, err
				}
				count, err = collection.CountDocuments(ctx, bson.D{
					{filter.Item, filter.Value},
				})
				if err != nil {
					return 0, nil, err
				}
			}
		}
	} else {
		// if superuser only
		if superuserOnly {
			newQuery := bson.D{{
				"$and",
				bson.A{
					bson.D{{
						constants.Role, constants.Superuser,
					}},
					bson.D{
						{},
					},
				},
			}}
			cursor, err = collection.Find(ctx, newQuery, findOptions)
			if err != nil {
				return 0, nil, err
			}
			count, err = collection.CountDocuments(ctx, newQuery)
			if err != nil {
				return 0, nil, err
			}
		} else {
			cursor, err = collection.Find(ctx, bson.D{{}}, findOptions)
			if err != nil {
				return 0, nil, err
			}
			count, err = collection.CountDocuments(ctx, bson.D{{}})
			if err != nil {
				return 0, nil, err
			}
		}
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		user := &dto.User{}
		if err = cursor.Decode(&user); err != nil {
			return 0, nil, err
		}
		users = append(users, user)
	}

	return count, users, nil
}

// Delete deletes user by ID
func (v *UserDAO) Delete(ctx context.Context, id string) error {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Users)
	if _, err := collection.DeleteOne(ctx, bson.D{{constants.ID, id}}); err != nil {
		return err
	}
	return nil
}

// BatchDelete deletes users by IDs
func (v *UserDAO) BatchDelete(ctx context.Context, ids []string) ([]string, error) {
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

// Update updates user
func (v *UserDAO) Update(ctx context.Context, user *dto.User) (*dto.User, error) {
	collection := v.client.Database(constants.Mhpss).Collection(constants.Users)
	_, err := collection.UpdateOne(ctx, bson.D{{constants.ID, user.ID}}, bson.D{
		{"$set", user},
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
