package loader

import (
	"encoding/json"
	"io/ioutil"
)

type UnitsInfoResponse struct {
	Units map[string]interface{} `json:"DOTAUnits"`
}

func GetUnitsInfo() (map[string]interface{}, error) {
	var info UnitsInfoResponse

	resp, err := getRequest(unitsInfoURL)
	if err != nil {
		return info.Units, err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return info.Units, err
	}

	err = json.Unmarshal(bytes, &info)
	if err != nil {
		return info.Units, err
	}

	return info.Units, nil
}
