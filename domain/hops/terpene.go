package hops

// Marker type — never instantiated, used only to parameterize TerpeneRange.
type terpeneMarker struct{}

// Terpene represents the total terpene profile for a hops
type Terpene = Range[terpeneMarker]

// ParseTerpene takes in a high and a low string and parses the values into a Terpene.
func ParseTerpene(h, l string) (Terpene, error) {
	return ParseRange[terpeneMarker](h, l)
}
