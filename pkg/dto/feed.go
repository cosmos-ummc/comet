package dto

// Feed ...
type Feed struct {
	ID          string `json:"id" bson:"id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	Link        string `json:"link" bson:"link"`
	ImgPath     string `json:"imgPath" bson:"imgPath"`
	Type        int64  `json:"type" bson:"type"`
}
