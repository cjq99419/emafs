package timer

import (
	"emafs/config"
	"emafs/store"
	_struct "emafs/struct"
	"emafs/tools"
	"fmt"
	"time"
)

func TickerGetDailyDesc() {
	ticker := time.NewTicker(config.TickerDailyDescFreq)
	go GetDailyDesc(ticker)
}

func GetDailyDesc(ticker *time.Ticker) {
	for _ = range ticker.C {
		getDailyDesc()
	}
}

func getDailyDesc() {
	url := config.DailyDescURL + config.KEY
	jsonM := tools.HttpReqGetMap(url)
	var dailyDesc _struct.DailyDesc
	if jsonM != nil {
		dailyDesc = _struct.DailyDesc{
			CurrentConfirmedCount: tools.ConvertToInt(jsonM["newslist.0.desc.currentConfirmedCount"]),
			ConfirmedCount:        tools.ConvertToInt(jsonM["newslist.0.desc.confirmedCount"]),
			SuspectedCount:        tools.ConvertToInt(jsonM["newslist.0.desc.suspectedCount"]),
			CuredCount:            tools.ConvertToInt(jsonM["newslist.0.desc.curedCount"]),
			DeadCount:             tools.ConvertToInt(jsonM["newslist.0.desc.deadCount"]),
			SeriousCount:          tools.ConvertToInt(jsonM["newslist.0.desc.seriousCount"]),
			CurrentConfirmedIncr:  tools.ConvertToInt(jsonM["newslist.0.desc.currentConfirmedIncr"]),
			ConfirmedIncr:         tools.ConvertToInt(jsonM["newslist.0.desc.confirmedIncr"]),
			SuspectedIncr:         tools.ConvertToInt(jsonM["newslist.0.desc.suspectedIncr"]),
			CuredIncr:             tools.ConvertToInt(jsonM["newslist.0.desc.curedIncr"]),
			DeadIncr:              tools.ConvertToInt(jsonM["newslist.0.desc.deadIncr"]),
			SeriousIncr:           tools.ConvertToInt(jsonM["newslist.0.desc.seriousIncr"]),
		}
		_ = store.DescStore(dailyDesc)
	} else {
		fmt.Println("can not get daily desc from tianxing")
	}
}
