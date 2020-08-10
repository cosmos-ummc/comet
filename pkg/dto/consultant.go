package dto

// Consultant ...
type Consultant struct {
	ID          string   `json:"id" bson:"id"`
	UserID      string   `json:"userId" bson:"userId"`
	Name        string   `json:"name" bson:"name"`
	PhoneNumber string   `json:"phoneNumber" bson:"phoneNumber"`
	Email       string   `json:"email" bson:"email"`
	TakenSlots  []string `json:"takenSlots" bson:"takenSlots"`
}
