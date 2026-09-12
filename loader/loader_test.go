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
