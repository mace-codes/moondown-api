package hops

import (
	"fmt"
)

// ErrInvalidType is returned when an invald hosp type is provided.
var ErrInvalidType = fmt.Errorf("invalid hops type provided")

type Type string

const (
	Aroma       Type = "aroma"
	Bittering   Type = "bittering"
	DualPurpose Type = "dual purpose"
)

// IsValid checks if the Type is valid.
func (t Type) IsValid() bool {
	switch t {
	case
		Aroma,
		Bittering,
		DualPurpose:
		return true
	default:
		return false
	}
}

// ParseType parses a string into a hops Type
func ParseType(s string) (Type, error) {
	typ := Type(s)
	if !typ.IsValid() {
		return "", ErrInvalidType
	}
	return typ, nil
}

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}

// Equals checks if two Type values are equal
func (t Type) Equals(other Type) bool {
	return t == other
}
