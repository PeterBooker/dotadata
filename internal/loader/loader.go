package loader

import (
	"fmt"
	"net"
	"net/http"
	"runtime"
	"time"
)

var (
	abilitiesInfoURL = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/npc_abilities.json"
	heroesInfoURL    = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/npc_heroes.json"
	itemsInfoURL     = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/items.json"
	unitsInfoURL     = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/scripts/npc/npc_units.json"

	abilitiesLangURL = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/localization/abilities_%s.json"
	heroesLangURL    = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/localization/hero_lore_%s.txt"
	itemsLangURL     = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/localization/items_%s.json"

	heropickerURL = "http://www.dota2.com/jsfeed/heropickerdata?l=%s"
	heropediaURL  = "http://www.dota2.com/jsfeed/heropediadata?feeds=%s&l=%s"

	dotaLangURL = "https://raw.githubusercontent.com/dotabuff/d2vpkr/master/dota/resource/dota_%s.txt"
)

var (
	client *http.Client
)

func init() {
	var netTransport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
	}

	client = &http.Client{
		Timeout:   time.Second * time.Duration(30),
		Transport: netTransport,
	}
}

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

// getRequest performs the HTTP GET request using the provided URL
func getRequest(URL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	// Set User Agent
	req.Header.Set("User-Agent", "dota2-tooltips")

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	// Check status code is 2XX
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("Invalid HTTP status code: %d", resp.StatusCode)
	}

	return resp, nil
}
