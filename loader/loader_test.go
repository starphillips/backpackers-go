package loader

import (
	"fmt"
	"slices"
	"testing"

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
