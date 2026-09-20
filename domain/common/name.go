package common

import (
	"fmt"
	"regexp"
)

// ErrInvalidName is returned when an invalid name is provided
var ErrInvalidName = fmt.Errorf("invalid name provided")

// Name represents the name of an ingredient
type Name string

var validName = regexp.MustCompile(`^[a-zA-Z0-9À-ÿ\s\-_'®™]+$`)

// ParseName takes in a string and parses it to a Name.
func ParseName(s string) (Name, error) {
	if s == "" {
		return "", fmt.Errorf("%w: name cannot be empty", ErrInvalidName)
	}
	if !validName.MatchString(s) {
		return "", fmt.Errorf("%w: %q", ErrInvalidName, s)
	}
	return Name(s), nil
}

// String returns the string value of a Name
func (n Name) String() string {
	return string(n)
}

// Equals checks if two Name values are equal
func (n Name) Equals(other Name) bool {
	return n == other
}
