package data

import (
	"log"

	"github.com/PeterBooker/dota2data/internal/db"
)

// Load populates the Data struct with data
func (d *Data) Load() {
	var needsUpdate bool

	abilities, err := db.GetAllFromBucket("abilities")
	if err != nil {
		log.Printf("Error loading Abilities: %s\n", err)
	}
	if len(abilities) == 0 {
		needsUpdate = true
	}
	d.Lock()
	d.Abilities = abilities
	d.Unlock()

	heroes, err := db.GetAllFromBucket("heroes")
	if err != nil {
		log.Printf("Error loading Heroes: %s\n", err)
	}
	if len(heroes) == 0 {
		needsUpdate = true
	}
	d.Lock()
	d.Heroes = heroes
	d.Unlock()

	items, err := db.GetAllFromBucket("items")
	if err != nil {
		log.Printf("Error loading Items: %s\n", err)
	}
	if len(items) == 0 {
		needsUpdate = true
	}
	d.Lock()
	d.Items = items
	d.Unlock()

	if needsUpdate {
		err := d.Update()
		if err != nil {
			log.Printf("Update failed: %s\n", err)
		}
	}
}
