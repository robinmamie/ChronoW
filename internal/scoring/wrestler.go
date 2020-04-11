package scoring

// Wrestler stores all necessary information for a given wrestler
type Wrestler struct {
	ID      uint32
	Name    string
	Country string
}

// Color designates the color of a wrestler
type Color int

// Represents 3 possibilities of color: either a wrestler, or none
const (
	NoWrestler Color = iota
	Red
	Blue
)
