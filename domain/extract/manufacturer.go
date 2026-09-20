package extract

import "fmt"

// ErrInvalidManufacturer is returned when an invalid manufacturer is provided.
var ErrManufacturer = fmt.Errorf("invalid manufacturer provided")

type Manufacturer string

const (
	Northwestern  Manufacturer = "Northwestern"
	Brewferm      Manufacturer = "Brewferm"
	Briess        Manufacturer = "Briess"
	Muntons       Manufacturer = "Muntons"
	Generic       Manufacturer = "Generic"
	MaillardMalts Manufacturer = "Maillard Malts®"
	Alexanders    Manufacturer = "Alexanders"
	Ireks         Manufacturer = "Irek's"
	JohnBull      Manufacturer = "John Bull"
	Coopers       Manufacturer = "Coopers"
	Weyermann     Manufacturer = "Weyermann"
	Morgans       Manufacturer = "Morgans"
	Mountmellick  Manufacturer = "Mountmellick"
)

// IsValid checks if the Manufacturer is valid.
func (m Manufacturer) IsValid() bool {
	switch m {
	case
		Northwestern,
		Brewferm,
		Briess,
		Muntons,
		Generic,
		MaillardMalts,
		Alexanders,
		Ireks,
		JohnBull,
		Coopers,
		Weyermann,
		Morgans,
		Mountmellick:
		return true
	default:
		return false
	}
}

// ParseManufacturer parses a string into a Manufacturer
func ParseManufacturer(s string) (Manufacturer, error) {
	cat := Manufacturer(s)
	if !cat.IsValid() {
		return "", ErrManufacturer
	}
	return cat, nil
}

// String returns the string value of a Manufacturer
func (m Manufacturer) String() string {
	return string(m)
}

// Equals checks if two Manufacturer values are equal
func (m Manufacturer) Equals(other Manufacturer) bool {
	return m == other
}
