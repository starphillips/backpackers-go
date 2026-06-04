package algorithm

// Algorithm structs represent processed, enriched, ready-to-score data

type Destination struct {
	ID int
	CountryID int
	Name string
	Latitude float64
	Longitude float64
	MinDays int
	MaxDays int
	Description string
	CostLevel string

	Tags []string
	Activities []Activity

	Score float64
	AssignedDays int
}

type Activity struct {
	ID int
	Name string
	Tags []string
	Description string
	PriceAverage float64
}

// IDs are not important to the algorithm for sorting. It is easier to compare based on the string e.g. if "hiking" vs if 1
// They are needed for identifying and matching to the DB

type UserInput struct {
	TotalDays int
	SelectedCountries []int // Goes by the ID
	SelectedTags []int
	SelectedBudget string
	SelectedPace string
}

// 