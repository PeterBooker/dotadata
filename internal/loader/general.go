package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func GetDotaLangInfo(lang string) (string, error) {
	var info string

	URL := fmt.Sprintf(dotaLangURL, lang)

	resp, err := getRequest(URL)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	info = string(bytes)

	return info, nil
}

func GetHeropediaInfo(t, lang string) (map[string]interface{}, error) {
	var info map[string]interface{}

	URL := fmt.Sprintf(heropediaURL, t, lang)

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

	info = info[t].(map[string]interface{})

	return info, nil
}
