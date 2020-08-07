package dto

type Report struct {
	Date   string           `json:"date" bson:"date"`
	Counts map[string]int64 `json:"counts" bson:"counts"`
}
