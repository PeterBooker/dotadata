package loader

import (
	"encoding/json"
	"io/ioutil"
)

type AbilitiesInfoResponse struct {
	Abilities map[string]interface{} `json:"DOTAAbilities"`
}

func GetAbilitiesInfo() (map[string]interface{}, error) {
	var info AbilitiesInfoResponse

	resp, err := getRequest(abilitiesInfoURL)
	if err != nil {
		return info.Abilities, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info.Abilities, err
	}

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return info.Abilities, err
	}

	return info.Abilities, nil
}
