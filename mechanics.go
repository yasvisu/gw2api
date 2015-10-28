package gw2api

type Trait struct {
	ID             int           `json:"id"`
	Name           string        `json:"icon"`
	Icon           string        `json:"icon"`
	Description    string        `json:"description"`
	Specialization int           `json:"specialization"`
	Tier           int           `json:"tier"`
	Slot           string        `json:"slot"`
	Facts          []Fact        `json:"facts"`
	TraitedFacts   []TraitedFact `json:"traited_facts"`
	Skills         []Skill       `json:"skills"`
}

type Skill struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Icon         string        `json:"icon"`
	Facts        []Fact        `json:"facts"`
	TraitedFacts []TraitedFact `json:"traited_facts"`
}

type FactType struct {
	//Common fields
	Text        string `json:"text"`
	Type        string `json:"type"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	Percent     int    `json:"percent"`
	Target      string `json:"target"`
	Duration    int    `json:"duration"`
	Status      string `json:"status"`
	ApplyCount  int    `json:"apply_count"`
	Value       int    `json:"value"`

	// AttributeAdjustFact

	// BuffFact

	// BuffConversionFact
	Source string `json:"source"`

	// ComboFieldFact
	FinisherType string `json:"finisher_type"`

	// DamageFact
	HitCount int `json:"hit_count"`

	// DistanceFact
	Distance int `json:"distance"`

	// NoDataFact

	// PercentFact

	// PrefixedBuffFact
	Prefix Fact `json:"prefix"`

	// RadiusFact

	// RangeFact

	// RechargeFact

	// TimeFact

	// UnblockableFact
}

type Fact struct {
	Text  string `json:"text"`
	Icon  string `json:"icon"`
	Type  string `json:"type"`
	Facts []Fact `json:"facts"`

	// Only as TraitedFact
	RequiresTrait int `json:"requires_trait"`
	Overrides     int `json:"overrides"`
}

type TraitedFact Fact

func Traits(lang string) (res []int, err error) {
	ver := "v2"
	tag := "traits"
	err = fetchEndpoint(ver, tag, lang, &res)
	return
}

func TraitIds(lang string, ids ...int) (traits []Trait, err error) {
	ver := "v2"
	tag := "traits"
	err = fetchDetailEndpoint(ver, tag, lang, stringSlice(ids), &traits)
	return
}

type Specialization struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Elite       bool   `json:"elite"`
	Icon        string `json:"icon"`
	Background  string `json:"background"`
	MinorTraits []int  `json:"minor_traits"`
	MajorTraits []int  `json:"major_traits"`
}

func Specializations(lang string) (res []string, err error) {
	ver := "v2"
	tag := "specializations"
	err = fetchEndpoint(ver, tag, lang, &res)
	return
}

func SpecializationIds(lang string, ids ...int) (specs []Specialization, err error) {
	ver := "v2"
	tag := "specializations"
	err = fetchDetailEndpoint(ver, tag, lang, stringSlice(ids), &specs)
	return
}
