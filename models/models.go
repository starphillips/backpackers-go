package models

type Countries struct {
	Name      string
	CountryID int
}

type Destinations struct {
	Name          string
	DestinationID int // only reference other structs when creating instances of it i.e. Countries.CountryID
	CountryID     int
	Latitude      float64
	Longitude     float64
	MinDays       int
	MaxDays       int
	Description   string
	CostLevel     string
}

type Tag struct {
	Name  string
	TagID int
}

type DestinationTag struct {
	DestinationID int
	TagID         int
}

type Activity struct {
	Name          string
	ActivityID    int
	DestinationID int
	TagID         int
	Description   string
	PriceAverage  float64
}

// Database structs mirror tables.
// Database structs represent raw data
