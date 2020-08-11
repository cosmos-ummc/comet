package dto

// Game ...
type Game struct {
	ID         string `json:"id" bson:"id"`
	LinkAdr    string `json:"linkAdr" bson:"linkAdr"`
	LinkIos    string `json:"linkIos" bson:"linkIos"`
	ImgPathAdr string `json:"imgPathAdr" bson:"imgPathAdr"`
	ImgPathIos string `json:"imgPathIos" bson:"imgPathIos"`
}
