package dto

// ChatMessage ...
type ChatMessage struct {
	ID        string `json:"id" bson:"id"`
	RoomID    string `json:"roomId" bson:"roomId"`
	SenderID  string `json:"senderId" bson:"senderId"`
	Content   string `json:"content" bson:"content"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}
