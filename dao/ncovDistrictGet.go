package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"strconv"
)

func NcovDistrictGet(name string) []_struct.NcovDistrict {
	db := tools.ConnectToDB()
	var regionList []int
	id := GetRegionId(name)
	var disId int
	rows, err := db.Query("SELECT id FROM region WHERE field_id=" + strconv.Itoa(id))
	errHandle.CheckErr(err, "failed to get district id in get opr")
	if rows != nil {
		for rows.Next() {
			rows.Scan(&disId)
			regionList = append(regionList, disId)
		}
	}
	var districtList []_struct.NcovDistrict
	for _, regionId := range regionList {
		nrows, err := db.Query("SELECT DISTINCT locale, address, lng, lat, source FROM location where region=" + strconv.Itoa(regionId))
		errHandle.CheckErr(err, "failed to select from location by regionId")
		if nrows != nil {
			for nrows.Next() {
				var tmp _struct.NcovDistrict
				nrows.Scan(&tmp.Locale, &tmp.Address, &tmp.Lng, &tmp.Lat, &tmp.Source)
				districtList = append(districtList, tmp)
			}
		}
	}
	return districtList
}

func GetAddressByLocationId(id int) string {
	db := tools.ConnectToDB()
	rows, _ := db.Query("SELECT address FROM location WHERE locationId=" + strconv.Itoa(id))
	var address string
	if rows != nil {
		rows.Next()
		rows.Scan(&address)
	}
	return address
}
