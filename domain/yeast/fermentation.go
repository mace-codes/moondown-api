package yeast

import (
	"github.com/mace-codes/moondown-api/domain/common"
)

// Marker type — never instantiated, used only to parameterize TerpeneRange.
type fermentationMarker struct{}

// Fermentation represents the known fermentation temperature range for a yeast
type Fermentation = common.Range[fermentationMarker]

// ParseFermentation takes in a high and a low string and parses the values into a Fermentation
func ParseFermentation(h, l string) (Fermentation, error) {
	return common.ParseRange[fermentationMarker](h, l)
}
