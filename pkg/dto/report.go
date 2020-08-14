package dto

type Report struct {
	DepressionNormal   int64   `json:"depressionNormal" bson:"depressionNormal"`
	DepressionMild     int64   `json:"depressionMild" bson:"depressionMild"`
	DepressionModerate int64   `json:"depressionModerate" bson:"depressionModerate"`
	DepressionSevere   int64   `json:"depressionSevere" bson:"depressionSevere"`
	DepressionExtreme  int64   `json:"depressionExtreme" bson:"depressionExtreme"`
	AnxietyNormal      int64   `json:"anxietyNormal" bson:"anxietyNormal"`
	AnxietyMild        int64   `json:"anxietyMild" bson:"anxietyMild"`
	AnxietyModerate    int64   `json:"anxietyModerate" bson:"anxietyModerate"`
	AnxietySevere      int64   `json:"anxietySevere" bson:"anxietySevere"`
	AnxietyExtreme     int64   `json:"anxietyExtreme" bson:"anxietyExtreme"`
	StressNormal       int64   `json:"stressNormal" bson:"stressNormal"`
	StressMild         int64   `json:"stressMild" bson:"stressMild"`
	StressModerate     int64   `json:"stressModerate" bson:"stressModerate"`
	StressSevere       int64   `json:"stressSevere" bson:"stressSevere"`
	StressExtreme      int64   `json:"stressExtreme" bson:"stressExtreme"`
	PtsdNormal         int64   `json:"ptsdNormal" bson:"ptsdNormal"`
	PtsdModerate       int64   `json:"ptsdModerate" bson:"ptsdModerate"`
	PtsdSevere         int64   `json:"ptsdSevere" bson:"ptsdSevere"`
	DailyNormal        int64   `json:"dailyNormal" bson:"dailyNormal"`
	DailySevere        int64   `json:"dailySevere" bson:"dailySevere"`
	DepressionCounts   []int64 `json:"depressionCounts" bson:"depressionCounts"`
	AnxietyCounts      []int64 `json:"anxietyCounts" bson:"anxietyCounts"`
	StressCounts       []int64 `json:"stressCounts" bson:"stressCounts"`
	PtsdCounts         []int64 `json:"ptsdCounts" bson:"ptsdCounts"`
	DailyCounts        []int64 `json:"dailyCounts" bson:"dailyCounts"`
	DepressionStatuses []int64 `json:"depressionStatuses" bson:"depressionStatuses"`
	AnxietyStatuses    []int64 `json:"anxietyStatuses" bson:"anxietyStatuses"`
	StressStatuses     []int64 `json:"stressStatuses" bson:"stressStatuses"`
	PtsdStatuses       []int64 `json:"ptsdStatuses" bson:"ptsdStatuses"`
	DailyStatuses      []int64 `json:"dailyStatuses" bson:"dailyStatuses"`
}
