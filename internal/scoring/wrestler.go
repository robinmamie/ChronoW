package scoring

// Wrestler stores all necessary information for a given wrestler
type Wrestler struct {
	ID      uint32
	Name    string
	Country string
}

type Color int

const (
	NoWrestler Color = iota
	Red
	Blue
)
