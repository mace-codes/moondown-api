package extract

import (
	"fmt"
)

// ErrInvalidType is returned when an invald extract type is provided.
var ErrInvalidType = fmt.Errorf("invalid extract type provided")

type Type string

const (
	DME Type = "Dry Malt Extract (DME)"
	LME Type = "Liquid Malt Extract (LME)"
)

// IsValid checks if the Type is valid.
func (t Type) IsValid() bool {
	switch t {
	case
		DME,
		LME:
		return true
	default:
		return false
	}
}

// ParseType parses a string into an extract Type
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
