package scoring

// WAge designates the age class of a bout
type WAge int

// Represents all age categories
const (
	OtherAge WAge = iota
	Rookie
	Piccolo
	JeunesseB
	JeunesseA
	U15
	Cadets
	Juniors
	U23
	Seniors
	Veterans
)
