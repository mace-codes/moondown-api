package hops

// Marker type — never instantiated, used only to parameterize OilRange.
type oilMarker struct{}

// Oil represents the total oil profile for a hops
type Oil = Range[oilMarker]

// ParseOil takes in a high and a low string and parses the values into an Oil.
func ParseOil(h, l string) (Oil, error) {
	return ParseRange[oilMarker](h, l)
}
