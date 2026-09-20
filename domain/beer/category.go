package beer

import "fmt"

// ErrInvalidCategory is returned when an invalid category is provided
var ErrInvalidCategory = fmt.Errorf("invalid category provided")

// Category represents a canonical BJCP beer style category.
type Category string

const (
	AlternativeFermentablesBeer Category = "Alternative Fermentables Beer"
	AmberAndBrownAmericanBeer   Category = "Amber And Brown American Beer"
	AmberBitterEuropeanBeer     Category = "Amber Bitter European Beer"
	AmberMaltyEuropeanLager     Category = "Amber Malty European Lager"
	AmericanPorterAndStout      Category = "American Porter And Stout"
	AmericanWildAle             Category = "American Wild Ale"
	BelgianAle                  Category = "Belgian Ale"
	BritishBitter               Category = "British Bitter"
	BrownBritishBeer            Category = "Brown British Beer"
	CzechLager                  Category = "Czech Lager"
	DarkBritishBeer             Category = "Dark British Beer"
	DarkEuropeanLager           Category = "Dark European Lager"
	EuropeanSourAle             Category = "European Sour Ale"
	FruitBeers                  Category = "Fruit Beer"
	GermanWheatBeer             Category = "German Wheat Beer"
	HistoricalBeer              Category = "Historical Beer"
	InternationalLager          Category = "International Lager"
	Ipa                         Category = "Ipa"
	IrishBeer                   Category = "Irish Beer"
	MonasticAle                 Category = "Monastic Ale"
	PaleAmericanAle             Category = "Pale American Ale"
	PaleBitterEuropeanBeer      Category = "Pale Bitter European Beer"
	PaleCommonwealthBeer        Category = "Pale Commonwealth Beer"
	PaleMaltyEuropeanLager      Category = "Pale Malty European Lager"
	ScottishAle                 Category = "Scottish Ale"
	SmokedBeer                  Category = "Smoked Beer"
	SpecialtyBeer               Category = "Specialty Beer"
	SpicedBeer                  Category = "Spiced Beer"
	StandardAmericanBeer        Category = "Standard American Beer"
	StrongAmericanAle           Category = "Strong American Ale"
	StrongBelgianAle            Category = "Strong Belgian Ale"
	StrongBritishAle            Category = "Strong British Ale"
	StrongEuropeanBeer          Category = "Strong European Beer"
	WoodBeer                    Category = "Wood Beer"
)

// IsValid checks if the Category is valid.
func (c Category) IsValid() bool {
	switch c {
	case
		AlternativeFermentablesBeer,
		AmberAndBrownAmericanBeer,
		AmberBitterEuropeanBeer,
		AmberMaltyEuropeanLager,
		AmericanPorterAndStout,
		AmericanWildAle,
		BelgianAle,
		BritishBitter,
		BrownBritishBeer,
		CzechLager,
		DarkBritishBeer,
		DarkEuropeanLager,
		EuropeanSourAle,
		FruitBeers,
		GermanWheatBeer,
		HistoricalBeer,
		InternationalLager,
		Ipa,
		IrishBeer,
		MonasticAle,
		PaleAmericanAle,
		PaleBitterEuropeanBeer,
		PaleCommonwealthBeer,
		PaleMaltyEuropeanLager,
		ScottishAle,
		SmokedBeer,
		SpecialtyBeer,
		SpicedBeer,
		StandardAmericanBeer,
		StrongAmericanAle,
		StrongBelgianAle,
		StrongBritishAle,
		StrongEuropeanBeer,
		WoodBeer:
		return true
	default:
		return false
	}
}

// ParseCategory parses a string into a Category
func ParseCategory(s string) (Category, error) {
	cat := Category(s)
	if !cat.IsValid() {
		return "", ErrInvalidCategory
	}
	return cat, nil
}

// String returns the string value of a beer category
func (c Category) String() string {
	return string(c)
}

// Equals checks if two Category values are equal
func (c Category) Equals(other Category) bool {
	return c == other
}
