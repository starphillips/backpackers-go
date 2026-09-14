package loader

import (
	"github.com/starphillips/backpackers-go/algorithm"
	"github.com/starphillips/backpackers-go/models"
)

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

func BuildDestinations(dbDestinations []models.Destinations, tagsByDestinationID map[int][]string, activitiesByDestinationID map[int][]algorithm.Activity) []algorithm.Destination {
	destinations := []algorithm.Destination{}

	for _, d := range dbDestinations {
		destinations = append(destinations, algorithm.Destination{
			ID:          d.DestinationID,
			CountryID:   d.CountryID,
			Name:        d.Name,
			Latitude:    d.Latitude,
			Longitude:   d.Longitude,
			MinDays:     d.MinDays,
			MaxDays:     d.MaxDays,
			Description: d.Description,
			CostLevel:   d.CostLevel,
			Tags:        tagsByDestinationID[d.DestinationID],
			Activities:  activitiesByDestinationID[d.DestinationID],
		})
	}

	return destinations
}

func LoadDestinations(
	dbTags []models.Tag,
	dbDestinationTags []models.DestinationTag,
	dbActivityTags []models.ActivityTag,
	dbActivities []models.Activity,
	dbDestinations []models.Destinations) []algorithm.Destination {

	tagNameByID := BuildTagNameMap(dbTags)
	// this function returns the tagNameByID. so we hold its vaule in the variable it literally is.

	tagsByActivityID := BuildTagsByActivity(dbActivityTags, tagNameByID)
	tagsByDestinationID := BuildTagsByDestination(dbDestinationTags, tagNameByID)
	activitiesByDestinationID := BuildActivitiesByDestination(dbActivities, tagsByActivityID)

	return BuildDestinations(dbDestinations, tagsByDestinationID, activitiesByDestinationID)

}

// BuildTagNameMap - map db tags IDs and to their name for comparisons
// BuildTagsByDestination - grab the destination's tags and append it to the destinations's list (as there can be multiple)
// BuildTagsByActivity - grab the activity's tags and append it to the activity's list (as there can be multiple)
// BuildActivitiesByDestination - map each activity to its destination
// BuildDestinations - map each db destination to the algo-ready destination
// LoadDestinations -
