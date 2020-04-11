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
	scoreR, scoreB := b.Totals()

	// Special victories (fall, disqualification, ...) overwrites everything
	if b.SpecialVictory != NoWrestler {
		return b.SpecialVictory, true
	}

	// If a wrestler has MaxCautions, they are eliminated
	if b.RedWrestler.Cautions == MaxCautions {
		return Blue, true
	}
	if b.BlueWrestler.Cautions == MaxCautions {
		return Red, true
	}

	// First, check if a wrestler has more points than another one (+ tech sup)
	if scoreR > scoreB {
		return Red, scoreR >= scoreB+b.TechnicalSuperiority
	}
	if scoreR < scoreB {
		return Blue, scoreR+b.TechnicalSuperiority <= scoreB
	}
	// If no points were scored, no wrestlers are winning yet
	if scoreR == 0 {
		return NoWrestler, false
	}

	// TODO check who has the more "biggest" points
	// TODO check number of cautions

	// TODO implement "IsFinished" in timer
	return b.LastScore, b.BoutTimer.TenthOfSecondsRemaining() == 0
}
