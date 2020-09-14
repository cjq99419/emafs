package config

import "time"

const (
	KEY                    = "e6443a1d71f7a829647d0159b1e7b101"
	DailyDescURL           = "http://api.tianapi.com/txapi/ncov/index?key="
	DailyNewsURL           = "http://api.tianapi.com/txapi/ncov/index?key="
	NcovCityURL            = "http://api.tianapi.com/txapi/ncovcity/index?key="
	NcovDistrictURL        = "http://api.tianapi.com/txapi/ncovnearby/index?key="
	DBDriverName           = "mysql"
	ServerPort             = "80"
	DBDataSource           = DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBSchema
	DBUser                 = "root"
	DBPassword             = "buptMXLY503"
	DBHost                 = "59.110.68.44"
	DBPort                 = "3306"
	DBSchema               = "emafs2.0"
	TickerDailyDescFreq    = 30 * time.Minute
	TickerDailyNewsFreq    = 60 * time.Minute
	TickerNcovCityFreq     = 2 * time.Hour
	TickerNcovDistrictFreq = 5 * time.Hour
)
