package common

import "fmt"

// ErrInvalidOrigin is returned when a provided origin is invalid
var ErrInvalidOrigin = fmt.Errorf("invalid origin")

// Origin represents the country of origin for an ingredient.
type Origin string

const (
	OriginGermany       Origin = "Germany"
	OriginUnitedStates  Origin = "United States"
	OriginBelgium       Origin = "Belgium"
	OriginUnitedKingdom Origin = "United Kingdom"
	OriginFrance        Origin = "France"
	OriginCanada        Origin = "Canada"
	OriginIreland       Origin = "Ireland"
	OriginChile         Origin = "Chile"
)

// IsValid checks if the Origin is valid.
func (o Origin) IsValid() bool {
	switch o {
	case
		OriginGermany,
		OriginUnitedStates,
		OriginBelgium,
		OriginUnitedKingdom,
		OriginFrance,
		OriginCanada,
		OriginIreland,
		OriginChile:
		return true
	default:
		return false
	}
}

// ParseOrigin parses a string into an Origin
func ParseOrigin(s string) (Origin, error) {
	origin := Origin(s)
	if !origin.IsValid() {
		return "", ErrInvalidOrigin
	}
	return origin, nil
}

// String returns the string value of a Origin
func (o Origin) String() string {
	return string(o)
}

// Equals checks if two Origin values are equal
func (o Origin) Equals(other Origin) bool {
	return o == other
}
