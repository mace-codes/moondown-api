package yeast

// Lab represents the lab in which the  yeast was developed
type Lab string

const (
	ImperialYeast  Lab = "Imperial Yeast"
	Danstar        Lab = "Danstar"
	Fermentis      Lab = "Fermentis"
	Brewferm       Lab = "Brewferm"
	Coopers        Lab = "Coopers"
	EastCoastYeast Lab = "East Coast Yeast"
	GigaYeast      Lab = "GigaYeast"
	MangroveJack   Lab = "Mangrove Jack"
	Omega          Lab = "Omega"
	WhiteLabs      Lab = "White Labs"
	Wyeast         Lab = "Wyeast"
	Muntons        Lab = "Munton's"
	TheYeastBay    Lab = "The Yeast Bay"
)

// String returns the string value of a Lab
func (l Lab) String() string {
	return string(l)
}

// Equals checks if two Lab values are equal
func (l Lab) Equals(other Lab) bool {
	return l == other
}
