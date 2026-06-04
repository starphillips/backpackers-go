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
