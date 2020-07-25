package dto

import "time"

// Activity ...
type Activity struct {
	ID         string    `json:"id" bson:"id"`
	UserID     string    `json:"userId" bson:"userId"`
	UserName   string    `json:"userName" bson:"userName"`
	OldPatient *Patient  `json:"oldPatient" bson:"oldPatient"`
	NewPatient *Patient  `json:"newPatient" bson:"newPatient"`
	OldSwab    *Swab     `json:"oldSwab" bson:"oldSwab"`
	NewSwab    *Swab     `json:"newSwab" bson:"newSwab"`
	Time       int64     `json:"time" bson:"time"`
	TTL        time.Time `json:"ttl" bson:"ttl"`
}
