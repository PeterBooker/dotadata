package data

import (
	"fmt"
	"strconv"
)

// Hero ...
type Hero struct {
	ID                        string        `json:"id"`
	Name                      string        `json:"name"`
	Title                     string        `json:"title"`
	Roles                     []string      `json:"roles"`
	Armor                     float64       `json:"armor"`
	AttackType                string        `json:"attack_type"`
	AttackDmgMin              uint8         `json:"attack_dmg_min"`
	AttackDmgMax              uint8         `json:"attack_dmg_max"`
	AttackRate                float64       `json:"attack_rate"`
	AttackRange               uint          `json:"attack_range"`
	AttributePrimary          string        `json:"attribute_primary"`
	AttributeBaseStrength     uint8         `json:"attribute_base_str"`
	AttributeStrengthGain     float64       `json:"attribute_str_gain"`
	AttributeBaseIntelligence uint8         `json:"attribute_base_int"`
	AttributeIntelligenceGain float64       `json:"attribute_int_gain"`
	AttributeBaseAgility      uint8         `json:"attribute_base_agi"`
	AttributeAgilityGain      float64       `json:"attribute_agi_gain"`
	MovementSpeed             uint          `json:"movement_speed"`
	MovementTurnRate          float64       `json:"movement_turn_rate"`
	Vision                    []uint        `json:"vision"`
	ImageURL                  string        `json:"img_url"`
	Bio                       string        `json:"bio"`
	Abilities                 []HeroAbility `json:"abilities"`
	Talents                   []HeroTalent  `json:"talents"`
}

type HeroAbility struct {
	Name     string `json:"name"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ImageURL string `json:"img_url"`
}

type HeroTalent struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

func (h *Hero) calcArmor() {
	armorCalc := h.Armor + (float64(h.AttributeBaseAgility) * 0.16)
	armor := fmt.Sprintf("%.2f", armorCalc)
	h.Armor, _ = strconv.ParseFloat(armor, 64)
}

func (h *Hero) calcDamage() {
	var bonus uint8
	switch h.AttributePrimary {
	case "agility":
		bonus = h.AttributeBaseAgility
	case "intelligence":
		bonus = h.AttributeBaseIntelligence
	case "strength":
		bonus = h.AttributeBaseStrength
	}
	h.AttackDmgMin = h.AttackDmgMin + bonus
	h.AttackDmgMax = h.AttackDmgMax + bonus
}

// GetHeroImageURL uses Hero data to produce a URL to a Heroes image.
func GetHeroImageURL(name, imgsize string) string {
	var size string

	// Set image size
	switch imgsize {

	case "full":
		size = "full.png"
	case "large":
		size = "lg.png"
	case "small":
		size = "sb.png"
	case "vertical":
		size = "vert.jpg"
	default:
		size = "lg.png"

	}

	// Produce URL to image
	return fmt.Sprintf("http://cdn.dota2.com/apps/dota2/images/heroes/%s_%s", name, size)
}

func primaryStat(name string) string {
	switch name {
	case "DOTA_ATTRIBUTE_AGILITY":
		name = "agility"
	case "DOTA_ATTRIBUTE_INTELLECT":
		name = "intellect"
	case "DOTA_ATTRIBUTE_STRENGTH":
		name = "strength"
	}

	return name
}

func NameToAlias(name string) string {
	switch name {
	case "rattletrap":
		name = "clockwerk"
	case "nevermore":
		name = "shadow_fiend"
	case "obsidian_destroyer":
		name = "outworld_devourer"
	case "shredder":
		name = "timbersaw"
	case "queenofpain":
		name = "queen_of_pain"
	case "wisp":
		name = "io"
	case "vengefulspirit":
		name = "vengeful_spirit"
	case "zuus":
		name = "zeus"
	case "abyssal_underlord":
		name = "underlord"
	case "necrolyte":
		name = "necrophos"
	case "skeleton_king":
		name = "wraith_king"
	}

	return name
}

func HeroNameFromAlias(name string) string {
	switch name {
	case "clockwerk":
		name = "rattletrap"
	case "shadow_fiend":
		name = "nevermore"
	case "outworld_devourer":
		name = "obsidian_destroyer"
	case "timbersaw":
		name = "shredder"
	case "queen_of_pain":
		name = "queenofpain"
	case "io":
		name = "wisp"
	case "vengeful_spirit":
		name = "vengefulspirit"
	case "zeus":
		name = "zuus"
	case "underlord":
		name = "abyssal_underlord"
	case "necrophos":
		name = "necrolyte"
	case "wraith_king":
		name = "skeleton_king"
	}

	return name
}
