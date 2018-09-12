package loader

import (
	"encoding/json"
	"io/ioutil"
)

type ItemsInfoResponse struct {
	Items map[string]interface{} `json:"DOTAAbilities"`
}

func GetItemsInfo() (map[string]interface{}, error) {
	var info ItemsInfoResponse

	resp, err := getRequest(itemsInfoURL)
	if err != nil {
		return info.Items, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info.Items, err
	}

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return info.Items, err
	}

	return info.Items, nil
}
