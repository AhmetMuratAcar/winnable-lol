package riot

import "fmt"

var getPuuidRegionMap = map[string]string{
	// Americas
	"BR1": "americas",
	"LA1": "americas",
	"LA2": "americas",
	"NA1": "americas",
	"OC1": "americas",

	// Europe
	"EUN1": "europe",
	"EUW1": "europe",
	"TR1":  "europe",
	"RU":   "europe",
	"ME1":  "europe",

	// Asia
	"JP1": "asia",
	"KR":  "asia",
	"TH2": "asia",
	"TW2": "asia",
	"VN2": "asia",
	"PH2": "asia",
	"SG2": "asia",
}

var matchDataRegionMap = map[string]string{
	// Americas
	"BR1": "americas",
	"LA1": "americas",
	"LA2": "americas",
	"NA1": "americas",

	// Europe
	"EUN1": "europe",
	"EUW1": "europe",
	"TR1":  "europe",
	"RU":   "europe",
	"ME1":  "europe",

	// Asia
	"JP1": "asia",
	"KR":  "asia",

	// SEA
	"OC1": "sea",
	"PH2": "sea",
	"SG2": "sea",
	"VN2": "sea",
	"TH2": "sea",
	"TW2": "sea",
}

func GetPuuidRegionRoute(region string) string {
	route, ok := getPuuidRegionMap[region]
	if !ok {
		// defaulting instead of returning an error because it doesn't actually
		// matter where you route to for this endpoint
		return "americas"
	}
	return route
}

func GetMatchDataRegionRoute(region string) (string, error) {
	route, ok := matchDataRegionMap[region]
	if !ok {
		return "", fmt.Errorf("region not in matchDataRegionMap")
	}

	return route, nil
}
