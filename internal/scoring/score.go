package scoring

const PossiblePoints []uint32 = []uint32{1, 2, 4, 5}

// Score represents the score of a given wrestler during a given bout
type Score struct {
	WrestlerInfo Wrestler
	Points       []uint32
	Cautions     uint32
}

func (s Score) Total() uint32 {
	uint32 sum = 0
	for p := range s.Points {
		sum += p
	}
	return sum
}