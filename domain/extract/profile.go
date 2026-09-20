package extract

import "github.com/mace-codes/moondown-api/domain/common"

// Profile represents the detailed breakdown of an extract
type Profile struct {
	Name         common.Name
	Origin       common.Origin
	Manufacturer Manufacturer
	Type         Type
	Color        *float64
	PPG          *float64
	PH           *float64
	Description  common.Description
}
