package yeast

import (
	"fmt"
	"regexp"
)

// ErrInvalidStrain is returned when a strain is invalid
var ErrInvalidStrain = fmt.Errorf("invalid strain provided")

// Strain repsents the strain ID for a yeast ingredient
type Strain string

var validStrainCode = regexp.MustCompile(`^[a-zA-Z0-9\s\-]+$`)

// ParseStrain takes in a string and parses it to a Strain.
func ParseStrain(s string) (Strain, error) {
	if s == "" {
		return "", fmt.Errorf("%w: namstraine cannot be empty", ErrInvalidStrain)
	}
	if !validStrainCode.MatchString(s) {
		return "", fmt.Errorf("%w: %q", ErrInvalidStrain, s)
	}
	return Strain(s), nil
}
