package common

import (
	"fmt"
	"regexp"
	"strings"
)

var varietyPattern = regexp.MustCompile(`^[a-zA-Z0-9À-ÿ\s\-_'®™]+$`)

// ErrInvalidVariety is returned when  an invlaid variety is provided
var ErrInvalidVariety = fmt.Errorf("invalid variety provided")

// Variety represents the name of a variety of ingredient
type Variety string

// ParseVariety parses a string to a Variety
func ParseVariety(s string) (Variety, error) {
	trimmed := strings.TrimSpace(s)

	if trimmed == "" || len(trimmed) > 50 {
		return "", ErrInvalidVariety
	}

	if !varietyPattern.MatchString(trimmed) {
		return "", fmt.Errorf("%w: msut match pattern: %s", ErrInvalidVariety, varietyPattern.String())
	}

	return Variety(trimmed), nil
}
