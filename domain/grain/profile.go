package grain

import "github.com/mace-codes/moondown-api/domain/common"

type Profile struct {
	Grain        common.Name
	Origin       common.Origin
	Type         Type
	Manufacturer common.Name
	MustMash     bool
	Color        Color
	Lintner      *float64
	TotalProtein *float64
	ExtractFgMin *float64
	Potential    *float64
	Moisture     *float64
	MaxUsage     *float64
	Description  common.Description
}
