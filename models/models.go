package models

type Countries struct {
	Name string
	CountryID int
}

type Destinations struct {
	Name string
	CountryID Countries.CountryID
	Latitude float64
	Longitude float64
	MinDays int
	MaxDays int
	Description string
	CostLevel string
}

