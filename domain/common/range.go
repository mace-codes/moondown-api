package common

import (
	"fmt"
	"strconv"
	"strings"
)

// Range represents a High/Low percentage range for a given type K.
// K is a phantom type parameter used only to make Range instantiations
// distinct from one another at compile time (e.g. AlphaAcid vs BetaAcid vs Lovibond etc.).
type Range[K any] struct {
	High float64
	Low  float64
}

// ParseRange takes a number for high and a number for low and parses them into a valid range.
func ParseRange[K any](h, l string) (Range[K], error) {
	high, err := strconv.ParseFloat(strings.TrimSpace(h), 64)
	if err != nil {
		return Range[K]{}, fmt.Errorf("invalid high value %q: %w", h, err)
	}

	low, err := strconv.ParseFloat(strings.TrimSpace(h), 64)
	if err != nil {
		return Range[K]{}, fmt.Errorf("invalid low value %q: %w", l, err)
	}

	return Range[K]{High: high, Low: low}, nil
}
