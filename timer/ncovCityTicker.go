package timer

import (
	"emafs/config"
	"emafs/store"
	_struct "emafs/struct"
	"emafs/tools"
	"fmt"
	"strings"
	"time"
)

func TickerGetNcovCity() {
	ticker := time.NewTicker(config.TickerNcovCityFreq)
	go GetNcovCity(ticker)
}

func GetNcovCity(ticker *time.Ticker) {
	for _ = range ticker.C {
		getNcovCity()
	}
}

func getNcovCity() {
	url := config.NcovCityURL + config.KEY
	jsonM := tools.HttpReqGetMap(url)
	if jsonM != nil {
		var ncovCity []_struct.NcovCity
		ncovCity = append(ncovCity, _struct.NcovCity{})
		i := 0
		for key, value := range jsonM {
			if strings.Contains(key, "provinceShortName") || strings.Contains(key, "cityName") {
				ncovCity[i].Name = value.(string)
				subStr := key
				subStr = strings.Replace(subStr, ".provinceShortName", ".", 1)
				subStr = strings.Replace(subStr, ".cityName", ".", 1)
				ncovCity[i].CurrentConfirmedCount = tools.ConvertToInt(jsonM[subStr+"currentConfirmedCount"])
				ncovCity[i].ConfirmedCount = tools.ConvertToInt(jsonM[subStr+"confirmedCount"])
				ncovCity[i].DeadCount = tools.ConvertToInt(jsonM[subStr+"deadCount"])
				ncovCity[i].CuredCount = tools.ConvertToInt(jsonM[subStr+"curedCount"])
				ncovCity[i].SuspectedCount = tools.ConvertToInt(jsonM[subStr+"suspectedCount"])
				ncovCity = append(ncovCity, _struct.NcovCity{})
				i++
			} else {
				continue
			}
		}
		store.NcovCityStore(ncovCity)
	} else {
		fmt.Println("can not get ncov city from tianxing")
	}
}
