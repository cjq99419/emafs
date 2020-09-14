package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"fmt"
	"strings"
)

func NcovCityGet(name string) _struct.NcovCity {
	db := tools.ConnectToDB()
	rows, err := db.Query(`SELECT currentConfirmedCount, confirmedCount, suspectedCount, curedCount, deadCount 
									FROM region WHERE name=?`, name)
	errHandle.CheckErr(err, "failed to getNcovCity in  ncovCityGet")
	rows.Next()
	var ncov _struct.NcovCity
	ncov.Name = name
	err = rows.Scan(&ncov.CurrentConfirmedCount, &ncov.ConfirmedCount, &ncov.SuspectedCount, &ncov.CuredCount, &ncov.DeadCount)
	return ncov
}

func GetProvinceList() []int {
	var list []int
	for key, value := range _struct.RegionsIdField {
		if value == 1 {
			list = append(list, key)
		}
	}
	fmt.Println("ccc", list)
	return list
}

func GetCityList(fieldId int) []int {
	var list []int
	for key, value := range _struct.RegionsIdField {
		if value == fieldId {
			list = append(list, key)
		}
	}
	return list
}

func GetDistrictList(fieldId int) []int {
	return GetCityList(fieldId)
}

func GetRegionName(regionId int) string {
	return _struct.Regions[regionId]
}

func GetRegionId(regionName string) int {
	if strings.Compare(regionName, "北京") == 0 {
		return 52
	} else if strings.Compare(regionName, "上海") == 0 {
		return 321
	} else if strings.Compare(regionName, "天津") == 0 {
		return 343
	} else if strings.Compare(regionName, "重庆") == 0 {
		return 394
	} else {
		for key, value := range _struct.Regions {
			if value == regionName {
				return key
			}
		}
	}
	return 0
}
