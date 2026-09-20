package hops

import (
	"github.com/mace-codes/moondown-api/domain/common"
)

// Marker type — never instantiated, used only to parameterize TerpeneRange.
type terpeneMarker struct{}

// Terpene represents the total terpene profile for a hops
type Terpene = common.Range[terpeneMarker]

// ParseTerpene takes in a high and a low string and parses the values into a Terpene.
func ParseTerpene(h, l string) (Terpene, error) {
	return common.ParseRange[terpeneMarker](h, l)
}
