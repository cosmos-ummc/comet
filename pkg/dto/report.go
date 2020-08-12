package dto

type Report struct {
	DepressionNormal   int64 `json:"depressionNormal" bson:"depressionNormal"`
	DepressionMild     int64 `json:"depressionMild" bson:"depressionMild"`
	DepressionModerate int64 `json:"depressionModerate" bson:"depressionModerate"`
	DepressionSevere   int64 `json:"depressionSevere" bson:"depressionSevere"`
	DepressionExtreme  int64 `json:"depressionExtreme" bson:"depressionExtreme"`
	AnxietyNormal      int64 `json:"anxietyNormal" bson:"anxietyNormal"`
	AnxietyMild        int64 `json:"anxietyMild" bson:"anxietyMild"`
	AnxietyModerate    int64 `json:"anxietyModerate" bson:"anxietyModerate"`
	AnxietySevere      int64 `json:"anxietySevere" bson:"anxietySevere"`
	AnxietyExtreme     int64 `json:"anxietyExtreme" bson:"anxietyExtreme"`
	StressNormal       int64 `json:"stressNormal" bson:"stressNormal"`
	StressMild         int64 `json:"stressMild" bson:"stressMild"`
	StressModerate     int64 `json:"stressModerate" bson:"stressModerate"`
	StressSevere       int64 `json:"stressSevere" bson:"stressSevere"`
	StressExtreme      int64 `json:"stressExtreme" bson:"stressExtreme"`
	PtsdNormal         int64 `json:"ptsdNormal" bson:"ptsdNormal"`
	PtsdSevere         int64 `json:"ptsdSevere" bson:"ptsdSevere"`
	DepressionCount1   int64 `json:"depressionCount1" bson:"depressionCount1"`
	DepressionCount2   int64 `json:"depressionCount2" bson:"depressionCount2"`
	AnxietyCount1      int64 `json:"anxietyCount1" bson:"anxietyCount1"`
	AnxietyCount2      int64 `json:"anxietyCount2" bson:"anxietyCount2"`
	StressCount1       int64 `json:"stressCount1" bson:"stressCount1"`
	StressCount2       int64 `json:"stressCount2" bson:"stressCount2"`
	PtsdCount1         int64 `json:"ptsdCount1" bson:"ptsdCount1"`
	PtsdCount2         int64 `json:"ptsdCount2" bson:"ptsdCount2"`
	DepressionStatus1  int64 `json:"depressionStatus1" bson:"depressionStatus1"`
	DepressionStatus2  int64 `json:"depressionStatus2" bson:"depressionStatus2"`
	AnxietyStatus1     int64 `json:"anxietyStatus1" bson:"anxietyStatus1"`
	AnxietyStatus2     int64 `json:"anxietyStatus2" bson:"anxietyStatus2"`
	StressStatus1      int64 `json:"stressStatus1" bson:"stressStatus1"`
	StressStatus2      int64 `json:"stressStatus2" bson:"stressStatus2"`
	PtsdStatus1        int64 `json:"ptsdStatus1" bson:"ptsdStatus1"`
	PtsdStatus2        int64 `json:"ptsdStatus2" bson:"ptsdStatus2"`
}
