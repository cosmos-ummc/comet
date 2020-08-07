package dto

// Question ...
type Question struct {
	ID       string `json:"id" bson:"id"`
	Category string `json:"category" bson:"category"`
	Type     string `json:"type" bson:"type"`
	Content  string `json:"content" bson:"content"`
	Score    int64  `json:"score" bson:"score"`
}
