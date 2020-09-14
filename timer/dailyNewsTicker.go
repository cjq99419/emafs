package timer

import (
	"emafs/config"
	"emafs/store"
	_struct "emafs/struct"
	"emafs/tools"
	"fmt"
	"strconv"
	"time"
)

func TickerGetDailyNews() {
	ticker := time.NewTicker(config.TickerDailyNewsFreq)
	go GetDailyNews(ticker)
}

func GetDailyNews(ticker *time.Ticker) {
	for _ = range ticker.C {
		getDailyNews()
	}
}

func getDailyNews() {
	store.NewsClear()
	url := config.DailyNewsURL + config.KEY
	jsonM := tools.HttpReqGetMap(url)
	if jsonM != nil {
		var dailyNews [5]_struct.DailyNews
		for i := 0; i < 5; i++ {
			str := "newslist.0.news." + strconv.Itoa(i)
			dailyNews[i] = _struct.DailyNews{
				PubDate:    tools.ConvertToInt(jsonM[str+".pubDate"]),
				PubDateStr: tools.ConvertToString(jsonM[str+".pubDateStr"]),
				Title:      tools.ConvertToString(jsonM[str+".title"]),
				Summary:    tools.ConvertToString(jsonM[str+".summary"]),
				InfoSource: tools.ConvertToString(jsonM[str+".infoSource"]),
				SourceURL:  tools.ConvertToString(jsonM[str+".sourceUrl"]),
			}
			store.NewsStore(dailyNews[i])
		}
	} else {
		fmt.Println("can not get daily news from tianxing")
	}
}
