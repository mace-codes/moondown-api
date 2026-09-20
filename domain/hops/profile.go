package hops

import (
	"github.com/mace-codes/moondown-api/domain/beer"
	"github.com/mace-codes/moondown-api/domain/common"
)

// Profile represnts a hops detailed profile breakdown
type Profile struct {
	Variety         Variety
	Origin          common.Origin
	Type            Type
	ChemicalProfile ChemicalProfile
	Description     common.Description
	Aroma           common.Description
	Beerstyle       []beer.Style
	Substitutions   []Variety
}
