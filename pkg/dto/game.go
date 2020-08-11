package dto

// Game ...
type Game struct {
	ID      string `json:"id" bson:"id"`
	Link    string `json:"link" bson:"link"`
	ImgPath string `json:"imgPath" bson:"imgPath"`
	Type    string `json:"type" bson:"type"`
}
