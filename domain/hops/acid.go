package hops

import (
	"github.com/mace-codes/moondown-api/domain/common"
)

// Marker type — never instantiated, used only to parameterize AcidRange.
type acidMarker struct{}

// Acid represents the data for an acid profile of a hops
type Acid = common.Range[acidMarker]

// ParseAcid takes in a high and a low string and parses the values into an Acid
func ParseAcid(h, l string) (Acid, error) {
	return common.ParseRange[acidMarker](h, l)
}
