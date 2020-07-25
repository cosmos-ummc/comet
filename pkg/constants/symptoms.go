package constants

// Cough Symptom
const (
	NoCough = iota + 1
	CoughBetter
	CoughNoChange
	StartCoughing
	CoughWorse
)

// Sore Throat Symptom
const (
	NoSoreThroat = iota + 1
	SoreThroatBetter
	SoreThroatNoChange
	StartSoreThroat
	SoreThroatWorse
)

// Fever Symptom
const (
	NoFever       = 1
	FeverNoChange = 3
	StartFever    = 4
	FeverWorse    = 5
)

// Breathing Difficulty Symptom
const (
	NoBreathDifficulty  = 1
	HasBreathDifficulty = 4
)

// Chest Pain Symptom
const (
	NoChestPain  = 1
	HasChestPain = 4
)

// Blue Face / lip Symptom
const (
	NoBlue  = 1
	HasBlue = 4
)

// Drowsy Symptom
const (
	NoDrowsy  = 1
	HasDrowsy = 4
)

// Loss of Taste or Smell
const (
	NoLoss  = 1
	HasLoss = 4
)
