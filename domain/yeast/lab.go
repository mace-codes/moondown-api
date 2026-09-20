package yeast

import "fmt"

// ErrInvalidLab is returned when an invalid lab is provided
var ErrInvalidLab = fmt.Errorf("invlaid lab provided")

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

// IsValid checks if the Lab is valid.
func (l Lab) IsValid() bool {
	switch l {
	case
		ImperialYeast,
		Danstar,
		Fermentis,
		Brewferm,
		Coopers,
		EastCoastYeast,
		GigaYeast,
		MangroveJack,
		Omega,
		WhiteLabs,
		Wyeast,
		Muntons,
		TheYeastBay:
		return true
	default:
		return false
	}
}

// ParseLab parses a string into a Lab
func ParseOrigin(l string) (Lab, error) {
	lab := Lab(l)
	if !lab.IsValid() {
		return "", ErrInvalidLab
	}
	return lab, nil
}

// String returns the string value of a Lab
func (l Lab) String() string {
	return string(l)
}

// Equals checks if two Lab values are equal
func (l Lab) Equals(other Lab) bool {
	return l == other
}
