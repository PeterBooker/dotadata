package data

import (
	"errors"
	"sync"
)

var (
	abilitiesInfoURL = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/npc_abilities.json"
	heroesInfoURL    = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/npc_heroes.json"
	itemsInfoURL     = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/items.json"
	unitsInfoURL     = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/npc_units.json"

	abilitiesLangURL = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/localization/abilities_%s.json"
	heroesLangURL    = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/localization/hero_lore_%s.txt"
	itemsLangURL     = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/localization/items_%s.json"

	dotaLangURL = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/dota_%s.txt"

	data *Data
)

func parseDamageType(t string) string {
	switch t {
	case "DAMAGE_TYPE_MAGICAL":
		return "Magical"
	case "DAMAGE_TYPE_PHYSICAL":
		return "Physical"
	default:
		return "Unknown"
	}
}

// Data ...
type Data struct {
	Abilities map[string][]byte
	Heroes    map[string][]byte
	Items     map[string][]byte
	Units     map[string][]byte
	sync.RWMutex
}

// New ...
func New() *Data {
	data = &Data{}
	return data
}

// GetHeroByID ...
func (d *Data) GetHeroByID(ID string) (string, error) {
	var name []byte
	var found bool
	d.RLock()
	name, found = d.Heroes[ID]
	if !found {
		d.RUnlock()
		return string(name), errors.New("Hero ID not recognized")
	}

	d.RUnlock()
	return string(name), nil
}

// GetHero ...
func (d *Data) GetHero(name, lang string) ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Heroes[name+"_"+lang]
	if !found {
		d.RUnlock()
		return data, errors.New("Hero name not recognized")
	}

	d.RUnlock()
	return data, nil
}

// GetHeroes ...
func (d *Data) GetHeroes() ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Heroes["list"]
	if !found {
		d.RUnlock()
		return data, errors.New("Hero list not found")
	}

	d.RUnlock()
	return data, nil
}

// GetAbility ...
func (d *Data) GetAbility(name, lang string) ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Abilities[name+"_"+lang]
	if !found {
		d.RUnlock()
		return data, errors.New("Ability name not recognized")
	}

	d.RUnlock()
	return data, nil
}

// GetAbilities ...
func (d *Data) GetAbilities() ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Abilities["list"]
	if !found {
		d.RUnlock()
		return data, errors.New("Ability list not found")
	}

	d.RUnlock()
	return data, nil
}

// GetItem ...
func (d *Data) GetItem(name, lang string) ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Items[name+"_"+lang]
	if !found {
		d.RUnlock()
		return data, errors.New("Item name not recognized")
	}

	d.RUnlock()
	return data, nil
}

// GetItems ...
func (d *Data) GetItems() ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Items["list"]
	if !found {
		d.RUnlock()
		return data, errors.New("Item list not found")
	}

	d.RUnlock()
	return data, nil
}

// GetUnit ...
func (d *Data) GetUnit(name, lang string) ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Units[name+"_"+lang]
	if !found {
		d.RUnlock()
		return data, errors.New("Unit name not recognized")
	}

	d.RUnlock()
	return data, nil
}

// GetUnits ...
func (d *Data) GetUnits() ([]byte, error) {
	var data []byte
	var found bool
	d.RLock()
	data, found = d.Units["list"]
	if !found {
		d.RUnlock()
		return data, errors.New("Unit list not found")
	}

	d.RUnlock()
	return data, nil
}
