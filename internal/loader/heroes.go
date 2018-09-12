package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type HeroesInfoResponse struct {
	Heroes map[string]interface{} `json:"DOTAHeroes"`
}

func GetHeroesInfo() (map[string]interface{}, error) {
	var info HeroesInfoResponse

	resp, err := getRequest(heroesInfoURL)
	if err != nil {
		return info.Heroes, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info.Heroes, err
	}

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return info.Heroes, err
	}

	return info.Heroes, nil
}

func GetHeroPickerInfo(lang string) (map[string]interface{}, error) {
	var info map[string]interface{}

	URL := fmt.Sprintf(heropickerURL, lang)

	resp, err := getRequest(URL)
	if err != nil {
		return info, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info, err
	}

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return info, err
	}

	return info, nil
}
