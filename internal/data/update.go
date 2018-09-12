package data

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PeterBooker/dota2data/internal/db"
	"github.com/PeterBooker/dota2data/internal/loader"
)

var (
	abilitiesInfo map[string]interface{}
	heroesInfo    map[string]interface{}
	itemsInfo     map[string]interface{}
	unitsInfo     map[string]interface{}

	heroPickerInfo map[string]map[string]interface{}

	heropediaAbilityInfo map[string]map[string]interface{}
	heropediaHeroInfo    map[string]map[string]interface{}
	heropediaItemInfo    map[string]map[string]interface{}
)

// Update ...
func (d *Data) Update() error {
	var err error

	ls := *GetLangs()

	heroPickerInfo = make(map[string]map[string]interface{})

	heropediaAbilityInfo = make(map[string]map[string]interface{})
	heropediaHeroInfo = make(map[string]map[string]interface{})
	heropediaItemInfo = make(map[string]map[string]interface{})

	// Get Language Data
	for _, lang := range ls {
		var data map[string]interface{}
		var err error

		data, err = loader.GetHeroPickerInfo(lang.Name)
		if err != nil {
			return err
		}
		heroPickerInfo[lang.Name] = data

		data, err = loader.GetHeropediaInfo("abilitydata", lang.Name)
		if err != nil {
			return err
		}
		heropediaAbilityInfo[lang.Name] = data

		data, err = loader.GetHeropediaInfo("herodata", lang.Name)
		if err != nil {
			return err
		}
		heropediaHeroInfo[lang.Name] = data

		data, err = loader.GetHeropediaInfo("itemdata", lang.Name)
		if err != nil {
			return err
		}
		heropediaItemInfo[lang.Name] = data
	}

	abilitiesInfo, err = loader.GetAbilitiesInfo()
	if err != nil {
		return err
	}

	heroesInfo, err = loader.GetHeroesInfo()
	if err != nil {
		return err
	}

	itemsInfo, err = loader.GetItemsInfo()
	if err != nil {
		return err
	}

	unitsInfo, err = loader.GetUnitsInfo()
	if err != nil {
		return err
	}

	// Update Hero Data
	heroes := make(map[string][]byte)
	var heroList []string

	tmpInfo := heroesInfo["npc_dota_hero_base"].(map[string]interface{})

	dayVision, _ := strconv.Atoi(tmpInfo["VisionDaytimeRange"].(string))
	nightVision, _ := strconv.Atoi(tmpInfo["VisionNighttimeRange"].(string))

	log.Println("Starting Heroes Update")

	for k, v := range heroesInfo {
		if k == "Version" || k == "npc_dota_hero_base" || k == "npc_dota_hero_target_dummy" {
			continue
		}
		info := v.(map[string]interface{})
		hero := Hero{}
		name := k[14:]
		hero.ID = info["HeroID"].(string)
		hero.Name = name
		as := getAbilities(info, name)
		ts := getTalents(info, name)

		hero.AttributePrimary = primaryStat(info["AttributePrimary"].(string))

		dmgMin, _ := strconv.Atoi(info["AttackDamageMin"].(string))
		hero.AttackDmgMin = uint8(dmgMin)

		dmgMax, _ := strconv.Atoi(info["AttackDamageMax"].(string))
		hero.AttackDmgMax = uint8(dmgMax)

		hero.AttackRate, _ = strconv.ParseFloat(info["AttackRate"].(string), 64)

		atkRange, _ := strconv.Atoi(info["AttackRange"].(string))
		hero.AttackRange = uint(atkRange)

		baseStr, _ := strconv.Atoi(info["AttributeBaseStrength"].(string))
		hero.AttributeBaseStrength = uint8(baseStr)

		hero.AttributeStrengthGain, _ = strconv.ParseFloat(info["AttributeStrengthGain"].(string), 64)

		baseInt, _ := strconv.Atoi(info["AttributeBaseIntelligence"].(string))
		hero.AttributeBaseIntelligence = uint8(baseInt)

		hero.AttributeIntelligenceGain, _ = strconv.ParseFloat(info["AttributeIntelligenceGain"].(string), 64)

		baseAgi, _ := strconv.Atoi(info["AttributeBaseAgility"].(string))
		hero.AttributeBaseAgility = uint8(baseAgi)

		hero.AttributeAgilityGain, _ = strconv.ParseFloat(info["AttributeAgilityGain"].(string), 64)

		hero.Armor, _ = strconv.ParseFloat(info["ArmorPhysical"].(string), 64)

		// Update Damage and Armor based on
		hero.calcDamage()
		hero.calcArmor()

		moveSpd, _ := strconv.Atoi(info["MovementSpeed"].(string))
		hero.MovementSpeed = uint(moveSpd)

		hero.MovementTurnRate, _ = strconv.ParseFloat(info["MovementTurnRate"].(string), 64)

		vision := []uint{uint(dayVision), uint(nightVision)}
		if info["VisionDaytimeRange"] != nil {
			day, _ := strconv.Atoi(info["VisionDaytimeRange"].(string))
			vision[0] = uint(day)
		}
		if info["VisionNighttimeRange"] != nil {
			night, _ := strconv.Atoi(info["VisionNighttimeRange"].(string))
			vision[1] = uint(night)
		}
		hero.Vision = vision

		hero.ImageURL = GetHeroImageURL(name, "vertical")

		// Add Localized Data
		for _, v := range ls {
			lang := v.Name
			hpinfo := heroPickerInfo[lang][name].(map[string]interface{})
			hero.Title = hpinfo["name"].(string)
			hero.AttackType = hpinfo["atk_l"].(string)
			hero.Bio = hpinfo["bio"].(string)

			hero.Roles = nil
			for _, v := range hpinfo["roles_l"].([]interface{}) {
				hero.Roles = append(hero.Roles, v.(string))
			}

			hero.Abilities = nil
			for _, name := range as {
				hpAbInfo, found := heropediaAbilityInfo[lang][name].(map[string]interface{})
				if !found {
					continue
				}
				ha := HeroAbility{
					Name:     name,
					Title:    hpAbInfo["dname"].(string),
					Desc:     hpAbInfo["desc"].(string),
					ImageURL: GetAbilityImageURL(name),
				}
				hero.Abilities = append(hero.Abilities, ha)
			}

			hero.Talents = nil
			for _, name := range ts {
				hpAbInfo, found := heropediaAbilityInfo[lang][name].(map[string]interface{})
				if !found {
					continue
				}
				ht := HeroTalent{
					Name:  name,
					Title: hpAbInfo["dname"].(string),
				}
				hero.Talents = append(hero.Talents, ht)
			}

			bytes, err := json.Marshal(hero)
			if err != nil {
				return err
			}

			// Add Hero Data
			heroes[name+"_"+lang] = bytes
			err = db.PutToBucket(name+"_"+lang, bytes, "heroes")
			if err != nil {
				log.Printf("Failed to save Hero %s to DB: %s\n", name+"_"+lang, err)
			}

			// Add Hero map ID to Name
			err = db.PutToBucket(hero.ID, []byte(name), "heroes")
			if err != nil {
				log.Printf("Failed to save Hero mapping %s to DB: %s\n", hero.ID, err)
			}
		}
		heroList = append(heroList, name)
	}

	bytes, err := json.Marshal(heroList)
	if err != nil {
		return err
	}

	heroes["list"] = bytes
	err = db.PutToBucket("list", bytes, "heroes")
	if err != nil {
		log.Printf("Failed to save Hero %s to DB: %s\n", "list", err)
	}

	log.Printf("%d Heroes Found", len(heroes))

	data.Lock()
	data.Heroes = heroes
	data.Unlock()

	// Update Ability Data
	abilities := make(map[string][]byte)
	var abilityList []string

	log.Println("Starting Abilities Update")

	for k := range heropediaAbilityInfo["english"] {
		if k[:7] == "special" {
			continue
		}
		fullname := k
		name := k

		if newName, found := checkAndStripHero(name, heroList); !found {
			continue
		} else {
			name = newName
		}

		ability := Ability{}

		abInfo, found := abilitiesInfo[k].(map[string]interface{})
		if !found {
			continue
		}
		ability.ID = abInfo["ID"].(string)
		ability.Name = name

		if abInfo["AbilityCooldown"] != nil {
			ability.Cooldown = parseCDorMC(abInfo["AbilityCooldown"].(string))
		}

		if abInfo["AbilityManaCost"] != nil {
			ability.ManaCost = parseCDorMC(abInfo["AbilityManaCost"].(string))
		}

		ability.ImageURL = GetAbilityImageURL(fullname)

		// Add Localized Data
		for _, v := range ls {
			lang := v.Name

			hpAbInfo, found := heropediaAbilityInfo[lang][fullname].(map[string]interface{})
			if !found {
				continue
			}

			ability.Title = hpAbInfo["dname"].(string)
			ability.Desc = hpAbInfo["desc"].(string)
			ability.Affects = hpAbInfo["affects"].(string)
			ability.Notes = hpAbInfo["notes"].(string)
			ability.Damage = hpAbInfo["dmg"].(string)
			ability.Attributes = hpAbInfo["attrib"].(string)
			ability.Lore = hpAbInfo["lore"].(string)

			bytes, err := json.Marshal(ability)
			if err != nil {
				return err
			}

			abilities[name+"_"+lang] = bytes
			err = db.PutToBucket(name+"_"+lang, bytes, "abilities")
			if err != nil {
				log.Printf("Failed to save Ability %s to DB: %s\n", name+"_"+lang, err)
			}
		}
		abilityList = append(abilityList, name)
	}

	bytes, err = json.Marshal(abilityList)
	if err != nil {
		return err
	}

	abilities["list"] = bytes
	err = db.PutToBucket("list", bytes, "abilities")
	if err != nil {
		log.Printf("Failed to save Ability %s to DB: %s\n", "list", err)
	}

	log.Printf("%d Abilities Found", len(abilities))

	data.Lock()
	data.Abilities = abilities
	data.Unlock()

	// Update Item Data
	items := make(map[string][]byte)
	var itemList []string

	log.Println("Starting Items Update")

	for k, v := range itemsInfo {
		if k == "Version" {
			continue
		}

		// Check for no matching language info
		if heropediaItemInfo["english"][k[5:]] == nil {
			continue
		}

		info := v.(map[string]interface{})
		item := Item{}
		name := k[5:]

		// Ignore Odd Items
		if name == "trident" || name == "mutation_tombstone" || name == "pocket_tower" || name == "combo_breaker" || name == "recipe_iron_talon" {
			continue
		}

		// Ignore recipes
		if len(name) > 7 && name[7:] == "recipe_" {
			continue
		}

		item.ID = info["ID"].(string)
		item.Name = name

		// Parse ShopTags
		switch v := info["ItemShopTags"].(type) {
		case string:
			item.ShopTags = strings.Split(v, ";")
		}

		// Parse Cost
		switch v := info["ItemCost"].(type) {
		case string:
			cost, _ := strconv.Atoi(v)
			item.Cost = uint(cost)
		default:
			item.Cost = 0
		}

		// Parse Quality
		switch v := info["ItemQuality"].(type) {
		case string:
			item.Quality = v
		}

		if info["AbilityCooldown"] != nil {
			item.Cooldown = parseCDorMC(info["AbilityCooldown"].(string))
		}

		if info["AbilityManaCost"] != nil {
			item.ManaCost = parseCDorMC(info["AbilityManaCost"].(string))
		}

		if info["SideShop"] != nil {
			ss, _ := strconv.ParseBool(info["SideShop"].(string))
			item.SideShop = ss
		}

		if info["SecretShop"] != nil {
			ss, _ := strconv.ParseBool(info["SecretShop"].(string))
			item.SecretShop = ss
		}

		// Parse Level
		if info["ItemBaseLevel"] != nil {
			lvl, _ := strconv.Atoi(info["ItemBaseLevel"].(string))
			item.Level = uint(lvl)
		}

		item.ImageURL = GetItemImageURL(name)

		itemInfo := heropediaItemInfo["english"][name].(map[string]interface{})

		var components []string
		if itemInfo["components"] != nil {
			for _, v := range itemInfo["components"].([]interface{}) {
				components = append(components, v.(string))
			}
		}

		// Add Localized Data
		for _, v := range ls {
			lang := v.Name

			hpItInfo, found := heropediaItemInfo[lang][name].(map[string]interface{})
			if !found {
				continue
			}

			item.Components = nil
			if v, found := itemsInfo["item_recipe_"+name]; found {
				info := v.(map[string]interface{})
				str := info["ItemCost"].(string)
				cost, _ := strconv.ParseInt(str, 10, 64)
				if cost > 0 {
					ic := ItemComponent{
						Name:     "recipe",
						Title:    "Recipe",
						Cost:     uint(cost),
						ImageURL: GetItemImageURL("recipe"),
					}
					item.Components = append(item.Components, ic)
				}
			}
			for _, name := range components {
				hpItemInfo, found := heropediaItemInfo[lang][name].(map[string]interface{})
				if !found {
					continue
				}
				info := itemsInfo["item_"+name].(map[string]interface{})
				str := info["ItemCost"].(string)
				cost, _ := strconv.ParseInt(str, 10, 64)
				ic := ItemComponent{
					Name:     name,
					Title:    hpItemInfo["dname"].(string),
					Cost:     uint(cost),
					ImageURL: GetItemImageURL(name),
				}
				item.Components = append(item.Components, ic)
			}

			item.Title = hpItInfo["dname"].(string)
			item.Desc = hpItInfo["desc"].(string)
			item.Notes = hpItInfo["notes"].(string)
			item.Lore = hpItInfo["lore"].(string)

			bytes, err := json.Marshal(item)
			if err != nil {
				return err
			}

			items[name+"_"+lang] = bytes
			err = db.PutToBucket(name+"_"+lang, bytes, "items")
			if err != nil {
				log.Printf("Failed to save Item %s to DB: %s\n", name+"_"+lang, err)
			}
		}
		itemList = append(itemList, name)
	}

	bytes, err = json.Marshal(itemList)
	if err != nil {
		return err
	}

	items["list"] = bytes
	err = db.PutToBucket("list", bytes, "items")
	if err != nil {
		log.Printf("Failed to save Item %s to DB: %s\n", "list", err)
	}

	log.Printf("%d Items Found", len(items))

	data.Lock()
	data.Items = items
	data.Unlock()

	log.Println("Update Complete")

	return nil
}

func getAbilities(hero map[string]interface{}, name string) []string {
	var abilities []string

	if name == "invoker" {
		// Invoker
		if ability := hero["Ability1"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability2"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability3"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability6"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability7"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability8"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability9"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability10"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability11"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability12"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability13"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability14"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability15"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability16"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
	} else if name == "morphling" {
		// Morphling
		if ability := hero["Ability1"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability2"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability3"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability4"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability5"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability6"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability7"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability8"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
	} else {
		// All Others
		if ability := hero["Ability1"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability2"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability3"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability4"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability5"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
		if ability := hero["Ability6"].(string); ability != "generic_hidden" {
			abilities = append(abilities, ability)
		}
	}

	return abilities
}

func getTalents(hero map[string]interface{}, name string) []string {
	var talents []string

	if name == "invoker" {
		// Invoker
		if talent := hero["Ability17"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability18"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability19"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability20"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability21"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability22"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability23"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability24"].(string); talent != "" {
			talents = append(talents, talent)
		}
	} else if name == "morphling" {
		// Morphling
		if talent := hero["Ability15"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability16"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability17"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability18"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability19"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability20"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability21"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability22"].(string); talent != "" {
			talents = append(talents, talent)
		}
	} else {
		// All Others
		if talent := hero["Ability10"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability11"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability12"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability13"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability14"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability15"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability16"].(string); talent != "" {
			talents = append(talents, talent)
		}
		if talent := hero["Ability17"].(string); talent != "" {
			talents = append(talents, talent)
		}
	}

	return talents
}

func inArray(s string, slist []string) bool {
	for _, val := range slist {
		if val == s {
			return true
		}
	}
	return false
}

func inArrayContains(s string, slist map[string][]byte) bool {
	for _, val := range slist {
		if strings.Contains(s, string(val)) {
			return true
		}
	}
	return false
}

func checkAndStripHero(s string, slist []string) (string, bool) {
	for i := 0; i < len(slist); i++ {
		if strings.Contains(s, slist[i]) {
			return s[len(slist[i])+1:], true
		}
	}
	return "", false
}

func parseCDorMC(s string) (items []float64) {
	if strings.Contains(s, " ") {
		ss := strings.Split(s, " ")
		for _, v := range ss {
			f, _ := strconv.ParseFloat(v, 64)
			var str string
			if f >= 1.0 {
				str = fmt.Sprintf("%.0f", f)
			} else {
				str = fmt.Sprintf("%.2f", f)
			}
			i, _ := strconv.ParseFloat(str, 64)
			items = append(items, i)
		}
	} else {
		f, _ := strconv.ParseFloat(s, 64)
		var str string
		if f >= 1.0 {
			str = fmt.Sprintf("%.0f", f)
		} else {
			str = fmt.Sprintf("%.2f", f)
		}
		i, _ := strconv.ParseFloat(str, 64)
		items = []float64{i}
	}

	return items
}
