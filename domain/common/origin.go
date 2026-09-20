package common

import (
	"fmt"
	"regexp"
	"strings"
)

var countryPattern = regexp.MustCompile(`^[a-zA-ZÀ-ÿ\s]+$`)

// ErrInvalidOrigin is returned when an invalid origin is provided
var ErrInvalidOrigin = fmt.Errorf("invalid origin provided")

// Origin represents the origin of a variety of ingredient
type Origin string

// ParseOrigin parses a string to an origin of an ingredient
func ParseOrigin(s string) (Origin, error) {
	trimmed := strings.TrimSpace(s)

	if trimmed == "" || len(trimmed) > 50 {
		return "", ErrInvalidOrigin
	}

	if !countryPattern.MatchString(trimmed) {
		return "", fmt.Errorf("%w: msut match pattern: %s", ErrInvalidOrigin, countryPattern.String())
	}

	return Origin(trimmed), nil
}
