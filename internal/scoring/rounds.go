package scoring

// WRound designates the round of a bout
type WRound int

// Represents all possible rounds
const (
	None WRound = iota
	Repechage
	Qualification
	Final32nd
	Final16th
	EightFinal
	QuarterFinal
	SemiFinal
	Final56
	Final35
	Final34
	Final12
)
