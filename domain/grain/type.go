package grain

import (
	"fmt"
)

// ErrInvalidType is returned when an invald grain type is provided.
var ErrInvalidType = fmt.Errorf("invalid grain type provided")

type Type string

const (
	Base             Type = "Base"
	ForColor         Type = "Color"
	Crystal          Type = "Crystal"
	Roasted          Type = "Roasted"
	RoastedMalt      Type = "Roasted Malt"
	Specialty        Type = "Specialty"
	SpecialtyAdjunct Type = "Specialty/Adjunct"
	DextrineStyle    Type = "Dextrine-Style"
)

// IsValid checks if the Type is valid.
func (t Type) IsValid() bool {
	switch t {
	case
		Base,
		ForColor,
		Crystal,
		Roasted,
		RoastedMalt,
		Specialty,
		SpecialtyAdjunct,
		DextrineStyle:
		return true
	default:
		return false
	}
}

// ParseType parses a string into a grain Type
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
