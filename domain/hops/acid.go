package hops

// Marker type — never instantiated, used only to parameterize AcidRange.
type acidMarker struct{}

// Acid represents the data for an acid profile of a hops
type Acid = Range[acidMarker]

// ParseAcid takes in a high and a low string and parses the values into an Acid
func ParseAcid(h, l string) (Acid, error) {
	return ParseRange[acidMarker](h, l)
}
