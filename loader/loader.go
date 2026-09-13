package loader

import (
	"github.com/starphillips/backpackers-go/algorithm"
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

func BuildTagsByActivity(dbActivityTags []models.ActivityTag, tagNameByID map[int]string) map[int][]string {
	tagsByActivityID := make(map[int][]string)

	for _, d := range dbActivityTags {
		tagsByActivityID[d.ActivityID] = append(tagsByActivityID[d.ActivityID], tagNameByID[d.TagID])
	}
	return tagsByActivityID
}

func BuildActivitiesByDestination(dbActivities []models.Activity, tagsByActivityID map[int][]string) map[int][]algorithm.Activity {
	// Take Name, ID, Description and Price Average directly from the models struct (db)
	// Map the Activity to the destination, append the list of activities at a destination as there can be multiple?
	activitiesByDestinationID := make(map[int][]algorithm.Activity)

	// i will need to start looping over models.activity
	for _, d := range dbActivities {
		// matching fields need to be appended
		activitiesByDestinationID[d.DestinationID] = append(activitiesByDestinationID[d.DestinationID], algorithm.Activity{
			ID:           d.ActivityID,
			Name:         d.Name,
			Tags:         tagsByActivityID[d.ActivityID], // tags will need a map look up
			Description:  d.Description,
			PriceAverage: d.PriceAverage,
		})

	}
	return activitiesByDestinationID
}

// BuildTagNameMap - ability to loop through the different tags IDs and provide the string value
// BuildTagsByDestination - ability to grab the tags for each destination and append it to the destinations's list (as there can be multiple)
// BuildTagsByActivity - ability to grab the tags for each activity and append it to the activity's list (as there can be multiple)
// BuildActivitiesByDestination - abiloity to match each activity to its destination
