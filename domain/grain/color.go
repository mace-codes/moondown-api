package grain

import "github.com/mace-codes/moondown-api/domain/common"

// Color represents the color profile range for a given grain
type Color struct {
	Lovibond Lovibond
	SRM      SRM
}

// Marker types — never instantiated, used only to parameterize Lovibond and SRM.
type lovibondMarker struct{}
type srmMarker struct{}

// Lovibond represents the lovibond color range for a grain
type Lovibond = common.Range[lovibondMarker]

// SRM represents the srm color range for a grain
type SRM = common.Range[srmMarker]

// ParseLovibond takes in a high and a low string and parses the values into a Lovibond.
func ParseLovibond(h, l string) (Lovibond, error) {
	return common.ParseRange[lovibondMarker](h, l)
}

// ParseSRM takes in a high and a low string and parses the values into an SRM.
func ParseSRM(h, l string) (SRM, error) {
	return common.ParseRange[srmMarker](h, l)
}
