package yeast

import (
	"github.com/mace-codes/moondown-api/domain/beer"
	"github.com/mace-codes/moondown-api/domain/common"
)

// Profile represents a yeast profile breakdown
type Profile struct {
	Strain           Strain
	Name             common.Name
	Species          Species
	Lab              Lab
	Fermentation     Fermentation
	Attenuation      Attenuation
	AlcoholTolerance *float64
	Notes            common.Description
	Styles           []beer.Style
	Substitutions    []Strain
}
