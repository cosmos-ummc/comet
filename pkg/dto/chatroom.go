package dto

// ChatRoom ...
type ChatRoom struct {
	ID             string   `json:"id" bson:"id"`
	ParticipantIDs []string `json:"participantIds" bson:"participantIds"`
	Blocked        bool     `json:"blocked" bson:"blocked"`
	Timestamp      int64    `json:"timestamp" bson:"timestamp"`
}
