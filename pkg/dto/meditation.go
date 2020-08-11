package dto

// Meditation ...
type Meditation struct {
	ID   string `json:"id" bson:"id"`
	Link string `json:"link" bson:"link"`
}
