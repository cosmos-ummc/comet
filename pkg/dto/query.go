package dto

// SortData ...
type SortData struct {
	Item  string
	Order string
}

// RangeData ...
type RangeData struct {
	From int
	To   int
}

// FilterData ...
type FilterData struct {
	Item  string
	Value string
}

// PatientsQuery ...
type PatientsQuery struct {
	Patients []*Patient `json:"patients" bson:"patients"`
	Count    []*Count   `json:"count" bson:"count"`
}

// Count ...
type Count struct {
	Count int64 `json:"count" bson:"count"`
}
