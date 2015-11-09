package gw2api

import "testing"

func TestTraits(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testTraits []int
	if testTraits, err = api.Traits(); err != nil {
		t.Error("Failed to fetch traits")
	}

	var traits []Trait
	if traits, err = api.TraitIds("en", testTraits[0], testTraits[1]); err != nil {
		t.Error("Failed to parse the trait data: ", err)
	} else if len(traits) != 2 {
		t.Error("Failed to fetch existing traits")
	}
}

func TestSpecializations(t *testing.T) {
	var err error
	api := NewGW2Api()

	var testSpecializations []int
	if testSpecializations, err = api.Specializations(); err != nil {
		t.Error("Failed to fetch specializations")
	}

	var specializations []Specialization
	if specializations, err = api.SpecializationIds("en", testSpecializations[0], testSpecializations[1]); err != nil {
		t.Error("Failed to parse the specialization data: ", err)
	} else if len(specializations) != 2 {
		t.Error("Failed to fetch existing specializations")
	}
}
