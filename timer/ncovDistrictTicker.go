package timer

import (
	"emafs/config"
	"emafs/dao"
	"emafs/store"
	_struct "emafs/struct"
	"emafs/tools"
	"strings"
	"time"
)

func TickerGetNcovDistrict() {
	ticker := time.NewTicker(config.TickerNcovDistrictFreq)
	go GetNcovDistrict(ticker)
}

func GetNcovDistrict(ticker *time.Ticker) {
	time.Sleep(time.Duration(20) * time.Minute)
	for _ = range ticker.C {
		getNcovDistrict()
	}
}

func getNcovDistrict() {
	var ncov [40][40][40]_struct.NcovDistricts
	i, j, k, l := 0, 0, 0, 0
	provinces := dao.GetProvinceList()
	for _, province := range provinces {
		provinceName := dao.GetRegionName(province)
		cities := dao.GetCityList(province)
		for _, city := range cities {
			cityName := dao.GetRegionName(city)
			districts := dao.GetDistrictList(city)
			for _, district := range districts {
				districtName := dao.GetRegionName(district)
				var url string
				if provinceName == "北京" || provinceName == "上海" || provinceName == "天津" || provinceName == "重庆" {
					url = config.NcovDistrictURL + config.KEY + "&province=" + provinceName +
						"市&city=" + cityName + "市&district=" + districtName
				} else {
					url = config.NcovDistrictURL + config.KEY + "&province=" + provinceName +
						"省&city=" + cityName + "市&district=" + districtName
				}
				jsonM := tools.HttpReqGetMap(url)
				time.Sleep(time.Duration(1) * time.Second)
				for key, value := range jsonM {
					if strings.Contains(key, "address") {
						ncov[i][j][k].Ncovs = append(ncov[i][j][k].Ncovs, _struct.NcovDistrict{})
						subStr := strings.Replace(key, ".address", ".", 1)
						ncov[i][j][k].Ncovs[l].Address = tools.ConvertToString(value)
						ncov[i][j][k].Ncovs[l].Source = tools.ConvertToString(jsonM[subStr+"source"])
						ncov[i][j][k].Ncovs[l].Lat = tools.ConvertToString(jsonM[subStr+"lat"])
						ncov[i][j][k].Ncovs[l].Lng = tools.ConvertToString(jsonM[subStr+"lng"])
						ncov[i][j][k].Ncovs[l].Locale = tools.ConvertToString(jsonM[subStr+"locale"])
						ncov[i][j][k].Ncovs[l].Region = district
						l++
					}
				}
				if len(jsonM) != 0 {
					k++
				}
				l = 0
			}
			j++
			k = 0
		}
		i++
		j = 0
	}
	store.NcovDistrictStore(ncov)
}
