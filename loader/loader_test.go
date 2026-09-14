package loader

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/starphillips/backpackers-go/algorithm"
	"github.com/starphillips/backpackers-go/models"
)

func TestBuildTagNameMap(t *testing.T) {
	dbTags := []models.Tag{
		{TagID: 1, Name: "Hiking"},
		{TagID: 2, Name: "Watersports"},
	}

	result := BuildTagNameMap(dbTags)

	if result[1] == "Hiking" {
		fmt.Println("Tag Name matches ID")
	} else {
		t.Errorf("Tag Name does not match ID")
	}
}

func TestBuildTagsByDestination(t *testing.T) {
	// define the tags first
	tagNameByID := map[int]string{
		1: "Hiking",
		2: "Watersports",
		3: "Adrenaline",
	}

	dbDestinationTags := []models.DestinationTag{
		{DestinationID: 1, TagID: 2},
		{DestinationID: 1, TagID: 3},
		{DestinationID: 2, TagID: 3},
	}

	result := BuildTagsByDestination(dbDestinationTags, tagNameByID)

	if !slices.Equal(result[1], []string{"Watersports", "Adrenaline"}) {
		t.Errorf("Destinations 1 tags wrong, got %v", result[1])
	}

	if !slices.Equal(result[2], []string{"Adrenaline"}) {
		t.Errorf("Destinations 2 tags wrong, got %v", result[2])
	}

}

func TestBuildTagsByActivity(t *testing.T) {
	tagNameByID := map[int]string{
		1: "Hiking",
		2: "Watersports",
		3: "Adrenaline",
	}

	dbActivityTags := []models.ActivityTag{
		{ActivityID: 1, TagID: 1},
		{ActivityID: 1, TagID: 2},
		{ActivityID: 2, TagID: 3},
	}

	results := BuildTagsByActivity(dbActivityTags, tagNameByID)

	if !slices.Equal(results[1], []string{"Hiking", "Watersports"}) {
		t.Errorf("Activity 1 tags are wrong, got %v", results[1])
	}

	if !slices.Equal(results[2], []string{"Adrenaline"}) {
		t.Errorf("Activity 2 tags are incorrect, got %v", results[2])
	}
}

func TestBuildActivitiesByDestination(t *testing.T) {
	tagsByActivityID := map[int][]string{
		1: {"Culture Rich"},
		2: {"Culture Rich", "Chill/Relaxed"},
		3: {"Hiking"},
	}

	dbActivities := []models.Activity{
		{
			Name:          "Recoleta Walking Tour",
			ActivityID:    1,
			DestinationID: 1,
			Description:   "Guided cultural walk through Recoleta and historic landmarks.",
			PriceAverage:  15.00,
		},
		{
			Name:          "Palermo Café Hopping",
			ActivityID:    2,
			DestinationID: 1,
			Description:   "Relaxed café and food experience in Palermo district.",
			PriceAverage:  20.00,
		},
		{
			Name:          "Laguna de los Tres Trek",
			ActivityID:    3,
			DestinationID: 2,
			Description:   "Full-day Patagonia hike with views of Mount Fitz Roy.",
			PriceAverage:  0.00,
		},
	}

	results := BuildActivitiesByDestination(dbActivities, tagsByActivityID)

	expectedDest1 := []algorithm.Activity{
		{
			ID:           1,
			Name:         "Recoleta Walking Tour",
			Tags:         []string{"Culture Rich"},
			Description:  "Guided cultural walk through Recoleta and historic landmarks.",
			PriceAverage: 15.00,
		},
		{
			ID:           2,
			Name:         "Palermo Café Hopping",
			Tags:         []string{"Culture Rich", "Chill/Relaxed"},
			Description:  "Relaxed café and food experience in Palermo district.",
			PriceAverage: 20.00,
		},
	}

	expectedDest2 := []algorithm.Activity{
		{
			ID:           3,
			Name:         "Laguna de los Tres Trek",
			Tags:         []string{"Hiking"},
			Description:  "Full-day Patagonia hike with views of Mount Fitz Roy.",
			PriceAverage: 0.00,
		},
	}

	if !reflect.DeepEqual(results[1], expectedDest1) {
		t.Errorf("Destination 1 activities wrong, got %v", results[1])
	}

	if !reflect.DeepEqual(results[2], expectedDest2) {
		t.Errorf("Destination 2 activities wrong, got %v", results[2])
	}

}

func TestBuildDestinations(t *testing.T) {
	dbDestinations := []models.Destinations{
		{
			Name:          "Buenos Aires",
			DestinationID: 1,
			CountryID:     1,
			Latitude:      -34.603700,
			Longitude:     -58.381600,
			MinDays:       3,
			MaxDays:       6,
			Description:   "Cultural capital with nightlife, food and historic districts.",
			CostLevel:     "Medium",
		},
		{
			Name:          "El Chaltén",
			DestinationID: 2,
			CountryID:     1,
			Latitude:      -49.331500,
			Longitude:     -72.886300,
			MinDays:       3,
			MaxDays:       5,
			Description:   "Patagonia hiking town known for mountain trails and glaciers.",
			CostLevel:     "Medium",
		},
	}

	tagsByDestinationID := map[int][]string{
		1: {"Culture Rich", "Chilled/Relaxed"},
	}

	activitiesByDestinationID := map[int][]algorithm.Activity{
		1: {
			{
				ID:           1,
				Name:         "Recoleta Walking Tour",
				Tags:         []string{"Culture Rich"},
				Description:  "Guided cultural walk through Recoleta and historic landmarks.",
				PriceAverage: 15.00,
			},
			{
				ID:           2,
				Name:         "Palermo Café Hopping",
				Tags:         []string{"Culture Rich", "Chill/Relaxed"},
				Description:  "Relaxed café and food experience in Palermo district.",
				PriceAverage: 20.00,
			},
		},
	}

	expDestination1 := []algorithm.Destination{
		{
			ID:          1,
			CountryID:   1,
			Name:        "Buenos Aires",
			Latitude:    -34.603700,
			Longitude:   -58.381600,
			MinDays:     3,
			MaxDays:     6,
			Description: "Cultural capital with nightlife, food and historic districts.",
			CostLevel:   "Medium",
			Tags:        tagsByDestinationID[1],
			Activities:  activitiesByDestinationID[1],
		},
	}

	expDestination2 := []algorithm.Destination{
		{
			ID:          2,
			CountryID:   1,
			Name:        "El Chaltén",
			Latitude:    -49.331500,
			Longitude:   -72.886300,
			MinDays:     3,
			MaxDays:     5,
			Description: "Patagonia hiking town known for mountain trails and glaciers.",
			CostLevel:   "Medium",
			Tags:        nil,
			Activities:  nil,
		},
	}

	results := BuildDestinations(dbDestinations, tagsByDestinationID, activitiesByDestinationID)
	if !reflect.DeepEqual(results[0], expDestination1[0]) {
		t.Errorf("Destination 1 built incorrectly, got %v", results[0])
	}

	if !reflect.DeepEqual(results[1], expDestination2[0]) {
		t.Errorf("Destination 2 built incorrectly, got %v", results[1])
	}

}

func TestLoadDestinations(t *testing.T) {
	dbTags := []models.Tag{
		{TagID: 1, Name: "Hiking"},
		{TagID: 2, Name: "Watersports"},
		{TagID: 3, Name: "Culture Rich"},
		{TagID: 4, Name: "Chilled/Relaxed"},
	}

	dbDestinationTags := []models.DestinationTag{
		{DestinationID: 1, TagID: 3},
		{DestinationID: 1, TagID: 4},
		{DestinationID: 2, TagID: 1},
		{DestinationID: 2, TagID: 2},
	}

	dbActivityTags := []models.ActivityTag{
		{ActivityID: 1, TagID: 1},
		{ActivityID: 1, TagID: 3},
		{ActivityID: 2, TagID: 3},
		{ActivityID: 2, TagID: 4},
		{ActivityID: 3, TagID: 1},
	}

	dbActivities := []models.Activity{
		{
			Name:          "Recoleta Walking Tour",
			ActivityID:    1,
			DestinationID: 1,
			Description:   "Guided cultural walk through Recoleta and historic landmarks.",
			PriceAverage:  15.00,
		},
		{
			Name:          "Palermo Café Hopping",
			ActivityID:    2,
			DestinationID: 1,
			Description:   "Relaxed café and food experience in Palermo district.",
			PriceAverage:  20.00,
		},
		{
			Name:          "Laguna de los Tres Trek",
			ActivityID:    3,
			DestinationID: 2,
			Description:   "Full-day Patagonia hike with views of Mount Fitz Roy.",
			PriceAverage:  0.00,
		},
	}

	dbDestinations := []models.Destinations{
		{
			Name:          "Buenos Aires",
			DestinationID: 1,
			CountryID:     1,
			Latitude:      -34.603700,
			Longitude:     -58.381600,
			MinDays:       3,
			MaxDays:       6,
			Description:   "Cultural capital with nightlife, food and historic districts.",
			CostLevel:     "Medium",
		},
		{
			Name:          "El Chaltén",
			DestinationID: 2,
			CountryID:     1,
			Latitude:      -49.331500,
			Longitude:     -72.886300,
			MinDays:       3,
			MaxDays:       5,
			Description:   "Patagonia hiking town known for mountain trails and glaciers.",
			CostLevel:     "Medium",
		},
	}

	expDestination1 := algorithm.Destination{
		ID:          1,
		CountryID:   1,
		Name:        "Buenos Aires",
		Latitude:    -34.603700,
		Longitude:   -58.381600,
		MinDays:     3,
		MaxDays:     6,
		Description: "Cultural capital with nightlife, food and historic districts.",
		CostLevel:   "Medium",
		Tags:        []string{"Culture Rich", "Chilled/Relaxed"},
		Activities: []algorithm.Activity{
			{ID: 1, Name: "Recoleta Walking Tour", Tags: []string{"Hiking", "Culture Rich"}, Description: "Guided cultural walk through Recoleta and historic landmarks.", PriceAverage: 15.00},
			{ID: 2, Name: "Palermo Café Hopping", Tags: []string{"Culture Rich", "Chilled/Relaxed"}, Description: "Relaxed café and food experience in Palermo district.", PriceAverage: 20.00},
		},
	}

	expDestination2 := algorithm.Destination{
		ID:          2,
		CountryID:   1,
		Name:        "El Chaltén",
		Latitude:    -49.331500,
		Longitude:   -72.886300,
		MinDays:     3,
		MaxDays:     5,
		Description: "Patagonia hiking town known for mountain trails and glaciers.",
		CostLevel:   "Medium",
		Tags:        []string{"Hiking", "Watersports"},
		Activities: []algorithm.Activity{
			{ID: 3, Name: "Laguna de los Tres Trek", Tags: []string{"Hiking"}, Description: "Full-day Patagonia hike with views of Mount Fitz Roy.", PriceAverage: 0.00},
		},
	}

	results := LoadDestinations(dbTags, dbDestinationTags, dbActivityTags, dbActivities, dbDestinations)
	if !reflect.DeepEqual(results[0], expDestination1) {
		t.Errorf("Destination 1 built incorrectly, got %v", results[0])
	}

	if !reflect.DeepEqual(results[1], expDestination2) {
		t.Errorf("Destination 2 built incorrectly, got %v", results[1])
	}
}
