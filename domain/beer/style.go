package beer

import "fmt"

// ErrInvalidStyle is returned when a style provided is invalid
var ErrInvalidStyle = fmt.Errorf("invalid style provided")

// Style represents a canonical BJCP beer style name.
type Style string

const (
	Altbier                            Style = "Altbier"
	AlternativeGrainBeer               Style = "Alternative Grain Beer"
	AlternativeSugarBeer               Style = "Alternative Sugar Beer"
	AmericanAmberAle                   Style = "American Amber Ale"
	AmericanBarleywine                 Style = "American Barleywine"
	AmericanBrownAle                   Style = "American Brown Ale"
	AmericanIPA                        Style = "American IPA"
	AmericanLager                      Style = "American Lager"
	AmericanLightLager                 Style = "American Light Lager"
	AmericanPaleAle                    Style = "American Pale Ale"
	AmericanPorter                     Style = "American Porter"
	AmericanStout                      Style = "American Stout"
	AmericanStrongAle                  Style = "American Strong Ale"
	AmericanWheatBeer                  Style = "American Wheat Beer"
	AustralianSparklingAle             Style = "Australian Sparkling Ale"
	AutumnSeasonalBeer                 Style = "Autumn Seasonal Beer"
	BalticPorter                       Style = "Baltic Porter"
	BelgianBlondAle                    Style = "Belgian Blond Ale"
	BelgianDarkStrongAle               Style = "Belgian Dark Strong Ale"
	BelgianDubbel                      Style = "Belgian Dubbel"
	BelgianGoldenStrongAle             Style = "Belgian Golden Strong Ale"
	BelgianPaleAle                     Style = "Belgian Pale Ale"
	BelgianSingle                      Style = "Belgian Single"
	BelgianTripel                      Style = "Belgian Tripel"
	BerlinerWeisse                     Style = "Berliner Weisse"
	BestBitter                         Style = "Best Bitter"
	BiereDeGarde                       Style = "Bière de Garde"
	BlondeAle                          Style = "Blonde Ale"
	BrettBeer                          Style = "Brett Beer"
	BritishBrownAle                    Style = "British Brown Ale"
	BritishGoldenAle                   Style = "British Golden Ale"
	BritishStrongAle                   Style = "British Strong Ale"
	CaliforniaCommon                   Style = "California Common"
	ClassicStyleSmokedBeer             Style = "Classic Style Smoked Beer"
	CommercialSpecialtyBeer            Style = "Commercial Specialty Beer"
	CreamAle                           Style = "Cream Ale"
	CzechAmberLager                    Style = "Czech Amber Lager"
	CzechDarkLager                     Style = "Czech Dark Lager"
	CzechPaleLager                     Style = "Czech Pale Lager"
	CzechPremiumPaleLager              Style = "Czech Premium Pale Lager"
	DarkMild                           Style = "Dark Mild"
	Doppelbock                         Style = "Doppelbock"
	DoubleIPA                          Style = "Double IPA"
	DunklesBock                        Style = "Dunkles Bock"
	DunklesWeissbier                   Style = "Dunkles Weissbier"
	Eisbock                            Style = "Eisbock"
	EnglishBarleyWine                  Style = "English Barley Wine"
	EnglishIPA                         Style = "English IPA"
	EnglishPorter                      Style = "English Porter"
	ExperimentalBeer                   Style = "Experimental Beer"
	Festbier                           Style = "Festbier"
	FlandersRedAle                     Style = "Flanders Red Ale"
	ForeignExtraStout                  Style = "Foreign Extra Stout"
	FruitBeer                          Style = "Fruit Beer"
	FruitLambic                        Style = "Fruit Lambic"
	FruitAndSpiceBeer                  Style = "Fruit and Spice Beer"
	GermanHellesExportbier             Style = "German Helles Exportbier"
	GermanLeichtbier                   Style = "German Leichtbier"
	GermanPils                         Style = "German Pils"
	Gose                               Style = "Gose"
	GrapeAle                           Style = "Grape Ale"
	Gueuze                             Style = "Gueuze"
	HazyIPA                            Style = "Hazy IPA"
	HellesBock                         Style = "Helles Bock"
	HistoricalBeerKellerbier           Style = "Historical Beer: Kellerbier"
	HistoricalBeerKentuckyCommon       Style = "Historical Beer: Kentucky Common"
	HistoricalBeerLichtenhainer        Style = "Historical Beer: Lichtenhainer"
	HistoricalBeerLondonBrownAle       Style = "Historical Beer: London Brown Ale"
	HistoricalBeerPiwoGrodziskie       Style = "Historical Beer: Piwo Grodziskie"
	HistoricalBeerPreProhibitionLager  Style = "Historical Beer: Pre-Prohibition Lager"
	HistoricalBeerPreProhibitionPorter Style = "Historical Beer: Pre-Prohibition Porter"
	HistoricalBeerRoggenbier           Style = "Historical Beer: Roggenbier"
	HistoricalBeerSahti                Style = "Historical Beer: Sahti"
	ImperialStout                      Style = "Imperial Stout"
	InternationalAmberLager            Style = "International Amber Lager"
	InternationalDarkLager             Style = "International Dark Lager"
	InternationalPaleLager             Style = "International Pale Lager"
	IrishExtraStout                    Style = "Irish Extra Stout"
	IrishRedAle                        Style = "Irish Red Ale"
	IrishStout                         Style = "Irish Stout"
	Kolsch                             Style = "Kolsch"
	Lambic                             Style = "Lambic"
	Marzen                             Style = "Marzen"
	MixedFermentationSourBeer          Style = "Mixed-Fermentation Sour Beer"
	MixedStyleBeer                     Style = "Mixed-Style Beer"
	MunichDunkel                       Style = "Munich Dunkel"
	MunichHelles                       Style = "Munich Helles"
	OatmealStout                       Style = "Oatmeal Stout"
	OldAle                             Style = "Old Ale"
	OrdinaryBitter                     Style = "Ordinary Bitter"
	OudBruin                           Style = "Oud Bruin"
	Rauchbier                          Style = "Rauchbier"
	Saison                             Style = "Saison"
	Schwarzbier                        Style = "Schwarzbier"
	ScottishExport                     Style = "Scottish Export"
	ScottishHeavy                      Style = "Scottish Heavy"
	ScottishLight                      Style = "Scottish Light"
	SpecialtyFruitBeer                 Style = "Specialty Fruit Beer"
	SpecialtyIPA                       Style = "Specialty IPA"
	SpecialtySmokedBeer                Style = "Specialty Smoked Beer"
	SpecialtySpiceBeer                 Style = "Specialty Spice Beer"
	SpecialtyWoodAgedBeer              Style = "Specialty Wood-Aged Beer"
	SpiceHerbOrVegetableBeer           Style = "Spice, Herb, or Vegetable Beer"
	StraightSourBeer                   Style = "Straight Sour Beer"
	StrongBitter                       Style = "Strong Bitter"
	SweetStout                         Style = "Sweet Stout"
	TropicalStout                      Style = "Tropical Stout"
	ViennaLager                        Style = "Vienna Lager"
	WeeHeavy                           Style = "Wee Heavy"
	Weissbier                          Style = "Weissbier"
	Weizenbock                         Style = "Weizenbock"
	Wheatwine                          Style = "Wheatwine"
	WildSpecialtyBeer                  Style = "Wild Specialty Beer"
	WinterSeasonalBeer                 Style = "Winter Seasonal Beer"
	Witbier                            Style = "Witbier"
	WoodAgedBeer                       Style = "Wood-Aged Beer"
)

// IsValid checks if the Style is valid.
func (s Style) IsValid() bool {
	switch s {
	case
		Altbier,
		AlternativeGrainBeer,
		AlternativeSugarBeer,
		AmericanAmberAle,
		AmericanBarleywine,
		AmericanBrownAle,
		AmericanIPA,
		AmericanLager,
		AmericanLightLager,
		AmericanPaleAle,
		AmericanPorter,
		AmericanStout,
		AmericanStrongAle,
		AmericanWheatBeer,
		AustralianSparklingAle,
		AutumnSeasonalBeer,
		BalticPorter,
		BelgianBlondAle,
		BelgianDarkStrongAle,
		BelgianDubbel,
		BelgianGoldenStrongAle,
		BelgianPaleAle,
		BelgianSingle,
		BelgianTripel,
		BerlinerWeisse,
		BestBitter,
		BiereDeGarde,
		BlondeAle,
		BrettBeer,
		BritishBrownAle,
		BritishGoldenAle,
		BritishStrongAle,
		CaliforniaCommon,
		ClassicStyleSmokedBeer,
		CommercialSpecialtyBeer,
		CreamAle,
		CzechAmberLager,
		CzechDarkLager,
		CzechPaleLager,
		CzechPremiumPaleLager,
		DarkMild,
		Doppelbock,
		DoubleIPA,
		DunklesBock,
		DunklesWeissbier,
		Eisbock,
		EnglishBarleyWine,
		EnglishIPA,
		EnglishPorter,
		ExperimentalBeer,
		Festbier,
		FlandersRedAle,
		ForeignExtraStout,
		FruitBeer,
		FruitLambic,
		FruitAndSpiceBeer,
		GermanHellesExportbier,
		GermanLeichtbier,
		GermanPils,
		Gose,
		GrapeAle,
		Gueuze,
		HazyIPA,
		HellesBock,
		HistoricalBeerKellerbier,
		HistoricalBeerKentuckyCommon,
		HistoricalBeerLichtenhainer,
		HistoricalBeerLondonBrownAle,
		HistoricalBeerPiwoGrodziskie,
		HistoricalBeerPreProhibitionLager,
		HistoricalBeerPreProhibitionPorter,
		HistoricalBeerRoggenbier,
		HistoricalBeerSahti,
		ImperialStout,
		InternationalAmberLager,
		InternationalDarkLager,
		InternationalPaleLager,
		IrishExtraStout,
		IrishRedAle,
		IrishStout,
		Kolsch,
		Lambic,
		Marzen,
		MixedFermentationSourBeer,
		MixedStyleBeer,
		MunichDunkel,
		MunichHelles,
		OatmealStout,
		OldAle,
		OrdinaryBitter,
		OudBruin,
		Rauchbier,
		Saison,
		Schwarzbier,
		ScottishExport,
		ScottishHeavy,
		ScottishLight,
		SpecialtyFruitBeer,
		SpecialtyIPA,
		SpecialtySmokedBeer,
		SpecialtySpiceBeer,
		SpecialtyWoodAgedBeer,
		SpiceHerbOrVegetableBeer,
		StraightSourBeer,
		StrongBitter,
		SweetStout,
		TropicalStout,
		ViennaLager,
		WeeHeavy,
		Weissbier,
		Weizenbock,
		Wheatwine,
		WildSpecialtyBeer,
		WinterSeasonalBeer,
		Witbier,
		WoodAgedBeer:
		return true
	default:
		return false
	}
}

// ParseStyle parses a string into a  beer Style
func ParseStyle(s string) (Style, error) {
	style := Style(s)
	if !style.IsValid() {
		return "", ErrInvalidStyle
	}
	return style, nil
}

// String returns the string value of a beer Style.
func (s Style) String() string {
	return string(s)
}

// Equals checks if two Style values are equal
func (s Style) Equals(other Style) bool {
	return s == other
}
