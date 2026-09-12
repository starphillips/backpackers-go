package questing

import "github.com/starphillips/backpackers-go/algorithm"

type Drink struct {
	ID        int
	Name      string
	BasePrice float64
}

type DrinkAddon struct {
	DrinkID    int
	AddonName  string
	AddonPrice float64
}

type MenuItem struct {
	Name        string
	BasePrice   float64
	Addons      []Addon
	TotalPrices []float64
}

type Addon struct {
	Name  string
	Price float64
}

func BuildMenu(drinks []Drink, addons []DrinkAddon) []MenuItem {

	// 1. Build a lookup map: drinkID → []Addon
	addonMap := make(map[int][]Addon)

	for _, a := range addons {
		addonMap[a.DrinkID] = append(addonMap[a.DrinkID], Addon{
			Name:  a.AddonName,
			Price: a.AddonPrice,
		})
	}

	// 2. Build the final enriched menu items
	var menu []MenuItem

	for _, d := range drinks {
		drinkAddons := addonMap[d.ID]

		// Calculate total prices (base + each addon)
		var totals []float64
		for _, ad := range drinkAddons {
			totals = append(totals, d.BasePrice+ad.Price)
		}

		// Build the final MenuItem
		menu = append(menu, MenuItem{
			Name:        d.Name,
			BasePrice:   d.BasePrice,
			Addons:      drinkAddons,
			TotalPrices: totals,
		})
	}

	return menu
}

func BuildDestinations(
	dbDestinations []models.Destinations,
	dbTags []models.Tag,
	dbDestinationTags []models.DestinationTag,
	dbActivities []models.Activity,
	dbActivityTags []models.ActivityTag) []models.Destination {

	activityMap := make(map[int][]algorithm.Activity)

	for _, a := range dbActivities {
		activityMap[a.DestinationID] = append(activityMap[a.DestinationID], loader.Activity{
			ID:           a.ActivityID,
			Name:         a.Name,
			Description:  a.Description,
			PriceAverage: a.PriceAverage,
			Tags:         activityTagMap[a.ActivityID], // ← you must fill this earlier
		})
	}
}
