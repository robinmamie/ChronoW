package scoring

import "github.com/robinmamie/ChronoW/internal/timer"

// MaxCautions is the number of cautions required for one wrestler to be
// eliminated
const MaxCautions uint32 = 3

// Bout stores all necessary information about a given bout
type Bout struct {
	ID                   uint32
	Style                WStyle
	Weight               uint32
	Round                WRound
	AgeCategory          WAge
	BoutTimer            timer.Timer
	RedWrestler          Score
	BlueWrestler         Score
	TechnicalSuperiority uint32
	SpecialVictory       Color
	LastScore            Color
}

// Totals returns the total number of points of both wrestlers
func (b *Bout) Totals() (uint32, uint32) {
	return b.RedWrestler.Total(), b.BlueWrestler.Total()
}

// Winner returns the current winner of the bout, and if it is finished or not
func (b *Bout) Winner() (Color, bool) {
	// Base case
	winner := b.LastScore
	isFinished := b.BoutTimer.TenthOfSecondsRemaining() == 0 // TODO implement "IsFinished" in timer
	scoreR, scoreB := b.Totals()

	switch {
	// Special victories (fall, disqualification, ...) overwrites everything
	case b.SpecialVictory != NoWrestler:
		winner, isFinished = b.SpecialVictory, true
	// If a wrestler has MaxCautions, they are eliminated
	case b.RedWrestler.Cautions == MaxCautions:
		winner, isFinished = Blue, true
	case b.BlueWrestler.Cautions == MaxCautions:
		winner, isFinished = Red, true
	// First, check if a wrestler has more points than another one
	// Also check if there is a tech. sup.
	case scoreR > scoreB:
		winner, isFinished = Red, scoreR >= scoreB+b.TechnicalSuperiority
	case scoreR < scoreB:
		winner, isFinished = Blue, scoreR+b.TechnicalSuperiority <= scoreB
	// TODO check who has the more "biggest" points
	//case
	case b.RedWrestler.Cautions < b.BlueWrestler.Cautions:
		winner = Red
	case b.RedWrestler.Cautions > b.BlueWrestler.Cautions:
		winner = Blue
	}

	return winner, isFinished
}
