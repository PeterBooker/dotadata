package data

import "fmt"

// Item ...
type Item struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Title      string          `json:"title"`
	Cooldown   []float64       `json:"cd,omitempty"`
	ManaCost   []float64       `json:"mc,omitempty"`
	Cost       uint            `json:"cost"`
	ShopTags   []string        `json:"shop_tags,omitempty"`
	Quality    string          `json:"quality,omitempty"`
	Components []ItemComponent `json:"components,omitempty"`
	SideShop   bool            `json:"sideshop,omitempty"`
	SecretShop bool            `json:"secretshop,omitempty"`
	Desc       string          `json:"desc"`
	Notes      string          `json:"notes"`
	Lore       string          `json:"lore"`
	Level      uint            `json:"level,omitempty"`
	Created    bool            `json:"created,omitempty"`
	ImageURL   string          `json:"img_url"`
}

type ItemComponent struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Cost     uint   `json:"cost"`
	ImageURL string `json:"img_url"`
}

// GetItemImageURL uses Ability data to produce a URL to a Heroes image.
func GetItemImageURL(name string) string {
	if len(name) >= 6 && name[:6] == "recipe" {
		return "http://cdn.dota2.com/apps/dota2/images/items/recipe_lg.png"
	}
	return fmt.Sprintf("http://cdn.dota2.com/apps/dota2/images/items/%s_lg.png", name)
}

func ItemNameFromAlias(name string) string {
	switch name {
	case "battle_fury":
		name = "bfury"
	case "healing_salve":
		name = "flask"
	case "town_portal_scroll":
		name = "tpscroll"
	case "daedalus":
		name = "greater_crit"
	case "ghost_scepter":
		name = "ghost"
	case "mango":
		name = "enchanted_mango"
	case "band_of_elvenskin":
		name = "boots_of_elves"
	case "gem_of_true_sight":
		name = "gem"
	case "shadow_blade":
		name = "invis_sword"
	case "eaglesong":
		name = "eagle"
	case "linkens_sphere":
		name = "sphere"
	case "slippers_of_agility":
		name = "slippers"
	case "sentry_ward":
		name = "ward_sentry"
	case "observer_ward":
		name = "ward_observer"
	case "drum_of_endurance":
		name = "ancient_janggo"
	case "crystalys":
		name = "lesser_crit"
	case "perseverance":
		name = "pers"
	case "aghanims_scepter":
		name = "ultimate_scepter"
	case "pipe_of_insight":
		name = "pipe"
	case "heart_of_tarrasque":
		name = "heart"
	case "manta_style":
		name = "manta"
	case "vladmirs_offering":
		name = "vladmir"
	case "assault_cuirass":
		name = "assault"
	case "scythe_of_vyse":
		name = "sheepstick"
	case "tango_shared":
		name = "tango_single"
	case "aegis_of_the_immortal":
		name = "aegis"
	case "iron_branch":
		name = "branches"
	case "observer_and_sentry_wards":
		name = "ward_dispenser"
	}

	return name
}
