package usaddress

import (
	"github.com/xrash/smetrics"
)

func stringSimilarity(s1, s2 string) float64 {
	return smetrics.JaroWinkler(s1, s2, 0.1, 4)
}

type weights struct {
	PrimaryNumber   float64
	StreetName      float64
	ZIPCode         float64
	City            float64
	State           float64
	StreetSuffix    float64
	StreetPredir    float64
	StreetPostdir   float64
	SecondaryUnit   float64
	Plus4           float64
	POBox           float64
	RuralRoute      float64
	HighwayContract float64
}

// Default weights
//
// PrimaryNumber, StreetName, and ZIPCode are given the highest weights because they are critical for identifying a specific location.
// City and State are important but less specific than ZIPCode.
// StreetSuffix, StreetPredir, and StreetPostdir have lower weights as they often have less impact on mail delivery.
// SecondaryUnit and Plus4 are less critical and have the lowest weights.
// POBox, RuralRoute, and HighwayContract are alternative address types and replace the street address components when present.
var defaultWeights = weights{
	PrimaryNumber:   0.25,
	StreetName:      0.25,
	ZIPCode:         0.20,
	City:            0.10,
	State:           0.05,
	StreetSuffix:    0.05,
	StreetPredir:    0.025,
	StreetPostdir:   0.025,
	SecondaryUnit:   0.025,
	Plus4:           0.025,
	POBox:           0.25,
	RuralRoute:      0.25,
	HighwayContract: 0.25,
}

func (a Address) Similarity(other Address) float64 {
	var totalWeight float64
	var similarityScore float64

	// Function to add similarity and weight
	addSimilarity := func(weight float64, s1, s2 string) {
		sim := stringSimilarity(s1, s2)
		similarityScore += sim * weight
		totalWeight += weight
	}

	// Check for POBox, RuralRoute, or HighwayContract
	if a.POBox != "" && other.POBox != "" {
		addSimilarity(defaultWeights.POBox, a.POBox, other.POBox)
	} else if a.RuralRoute != "" && other.RuralRoute != "" {
		addSimilarity(defaultWeights.RuralRoute, a.RuralRoute, other.RuralRoute)
	} else if a.HighwayContract != "" && other.HighwayContract != "" {
		addSimilarity(defaultWeights.HighwayContract, a.HighwayContract, other.HighwayContract)
	} else {
		// Compare street address components
		addSimilarity(defaultWeights.PrimaryNumber, a.PrimaryNumber, other.PrimaryNumber)
		addSimilarity(defaultWeights.StreetPredir, a.StreetPredir, other.StreetPredir)
		addSimilarity(defaultWeights.StreetName, a.StreetName, other.StreetName)
		addSimilarity(defaultWeights.StreetSuffix, a.StreetSuffix, other.StreetSuffix)
		addSimilarity(defaultWeights.StreetPostdir, a.StreetPostdir, other.StreetPostdir)
		addSimilarity(defaultWeights.SecondaryUnit, a.SecondaryUnit, other.SecondaryUnit)
	}

	// Compare City, State, ZIPCode, Plus4
	addSimilarity(defaultWeights.City, a.City, other.City)
	addSimilarity(defaultWeights.State, a.State, other.State)
	addSimilarity(defaultWeights.ZIPCode, a.ZIPCode, other.ZIPCode)
	if a.Plus4 != "" || other.Plus4 != "" {
		addSimilarity(defaultWeights.Plus4, a.Plus4, other.Plus4)
	}

	if totalWeight == 0 {
		return 0.0
	}
	return similarityScore / totalWeight
}
