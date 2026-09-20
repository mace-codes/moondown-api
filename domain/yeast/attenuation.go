package yeast

import "github.com/mace-codes/moondown-api/domain/common"

// Fermentation represents the known fermentation temperature range for a yeast
type Attenuation struct {
	Range    AttenuationRange
	Strength common.Strength
}

// Marker type — never instantiated, used only to parameterize Attenuationange.
type attenuationMarker struct{}

type AttenuationRange = common.Range[attenuationMarker]

// ParseAttenuation takes in a high and a low string and parses the values into a Attenuation.Range
func ParseAttenuation(h, l string) (AttenuationRange, error) {
	return common.ParseRange[attenuationMarker](h, l)
}
