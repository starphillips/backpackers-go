package loader

import (
	"github.com/starphillips/backpackers-go/models"
)

// algorithm.Destination.Tags is []string, but DestinationTag/ActivityTag only gives IDs
func BuildTagNameMap(dbTags []models.Tag) map[int]string {
	tagNameByID := make(map[int]string)

	// We are looping over every tag id in the destinations tag and grabbing the name of it
	for _, d := range dbTags {
		tagNameByID[d.TagID] = d.Name
	}

	return tagNameByID
}

func BuildTagsByDestination(dbDestinationTags []models.DestinationTag, tagNameByID map[int]string) map[int][]string {
	tagsByDestinationID := make(map[int][]string)

	// append tags to an existing list, as a destination has multiple tags
	for _, d := range dbDestinationTags {
		tagsByDestinationID[d.DestinationID] = append(tagsByDestinationID[d.DestinationID], tagNameByID[d.TagID])
	}

	return tagsByDestinationID
}

// func BuildDestinationNameMap(dbTags []models.DestinationTag) map[int]int {
// 	tagNameByDestinationID
// }
