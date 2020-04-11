package scoring

// WStyle designates the possible wrestling styles
type WStyle int

// Represents the wrestling styles
const (
	Freestyle WStyle = iota
	GrecoRoman
	FemaleWrestling
)

// TODO add language in argument
// String return the name of the style
func (w WStyle) String() string {
	return [...]string{"Freestyle", "Greco-Roman", "Female Wrestling"}[w]
}

// ShortString returns the abbreviated version of the style's name
func (w WStyle) ShortString() string {
	return [...]string{"FS", "GR", "FW"}[w]
}
