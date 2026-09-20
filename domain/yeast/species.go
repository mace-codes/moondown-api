package yeast

// Species represents the specific species of a strain of yeast
type Species string

const (
	Ale           Species = "Ale"
	Blend         Species = "Blend"
	Brettanomyces Species = "Brettanomyces"
	Lactobacillus Species = "Lactobacillus"
	Lager         Species = "Lager"
	MaloLactic    Species = "Malo-Lactic"
)

// String returns the string value of a Species
func (s Species) String() string {
	return string(s)
}

// Equals checks if two Species values are equal
func (s Species) Equals(other Species) bool {
	return s == other
}
