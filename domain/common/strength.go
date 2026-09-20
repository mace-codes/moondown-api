package common

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var ErrInvalidStrength = fmt.Errorf("invalid strength")

// Strength represents a relative intensity rating on an ordinal scale.
type Strength int

const (
	StrengthUnknown Strength = iota
	StrengthVeryLow
	StrengthLow
	StrengthMediumLow
	StrengthMedium
	StrengthMediumHigh
	StrengthHigh
	StrengthVeryHigh
)

var strengthNames = map[Strength]string{
	StrengthVeryLow:    "Very Low",
	StrengthLow:        "Low",
	StrengthMediumLow:  "Medium-Low",
	StrengthMedium:     "Medium",
	StrengthMediumHigh: "Medium-High",
	StrengthHigh:       "High",
	StrengthVeryHigh:   "Very High",
}

// String returns the string value for a Strength
func (s Strength) String() string {
	if name, ok := strengthNames[s]; ok {
		return name
	}
	return "Unknown"
}

var titleCaser = cases.Title(language.English)

// ParsesStrength take sin a string and parses it to a Strength
func ParseStrength(s string) (Strength, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return StrengthUnknown, nil
	}
	normalized := titleCaser.String(strings.ToLower(s))
	for strength, name := range strengthNames {
		if name == normalized {
			return strength, nil
		}
	}
	return StrengthUnknown, fmt.Errorf("%w: %q", ErrInvalidStrength, s)
}
