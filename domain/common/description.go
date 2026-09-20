package common

import (
	"fmt"
	"regexp"
	"strings"
)

// ErrInvalidDescription is returned when a description is invalid
var ErrInvalidDescription = fmt.Errorf("invalid description")

var validDescription = regexp.MustCompile(`^[a-zA-Z0-9À-ÿ\s\-']*$`)

// Description is the string description of the hops variety
type Description string

// ParseDescription takes in a string and parses it into a Description
func ParseDescription(s string) (Description, error) {
	s = strings.TrimSpace(s)
	s = strings.ReplaceAll(s, "*", "")
	s = strings.TrimSpace(s)

	if !validDescription.MatchString(s) {
		return "", fmt.Errorf("%w: %q", ErrInvalidDescription, s)
	}
	return Description(s), nil
}
