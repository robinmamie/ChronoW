package scoring

// PossiblePoints designates all points that can be scored during a bout
var PossiblePoints [4]uint32 = [4]uint32{1, 2, 4, 5}

// Score represents the score of a given wrestler during a given bout
type Score struct {
	WrestlerInfo Wrestler
	Points       []uint32
	Cautions     uint32
}

// Total returns the total number of points of a given wrestler
func (s *Score) Total() uint32 {
	sum := uint32(0)
	for _, p := range s.Points {
		sum += p
	}
	return sum
}

// TODO implement method giving the repartition of points for scoring
