package hops

// ChemicalProfile represents the chemical breakdown of a hops
type ChemicalProfile struct {
	AlphaAcid     Acid
	BetaAcid      Acid
	Cohumulone    Acid
	TotalOil      Oil
	Myrcene       Oil
	Caryophyllene Terpene
	Humulene      Terpene
	Farnesene     Terpene
}
