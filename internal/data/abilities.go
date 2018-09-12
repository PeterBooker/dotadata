package data

import "fmt"

// Ability ...
type Ability struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Title      string    `json:"title"`
	Cooldown   []float64 `json:"cd,omitempty"`
	ManaCost   []float64 `json:"mc,omitempty"`
	Desc       string    `json:"desc"`
	Affects    string    `json:"affects"`
	Notes      string    `json:"notes"`
	Damage     string    `json:"dmg,omitempty"`
	Attributes string    `json:"attributes"`
	Lore       string    `json:"lore"`
	ImageURL   string    `json:"img_url"`
}

// GetAbilityImageURL uses Ability data to produce a URL to a Heroes image.
func GetAbilityImageURL(name string) string {
	return fmt.Sprintf("http://cdn.dota2.com/apps/dota2/images/abilities/%s_hp2.png", name)
}
