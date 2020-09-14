package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"strconv"
)

func NcovPublicGet(city, locale string) []_struct.NcovDistrict {
	db := tools.ConnectToDB()
	var districtIdList []int
	var locationList []int
	var i int
	id := GetRegionId(city)
	districtIdList = GetDistrictList(id)
	for _, regionId := range districtIdList {
		rows, err := db.Query("SELECT DISTINCT locationId FROM location WHERE region=" + strconv.Itoa(regionId) + " and locale like '%" + locale + "%'")
		errHandle.CheckErr(err, "failed to select locationId from location by region and locale")
		if rows != nil {
			for rows.Next() {
				rows.Scan(&i)
				locationList = append(locationList, i)
			}
		}
	}
	var districtList []_struct.NcovDistrict
	for _, locationId := range locationList {
		rows, err := db.Query("SELECT DISTINCT locale, address, lng, lat, source FROM location WHERE locationId=" + strconv.Itoa(locationId))
		errHandle.CheckErr(err, "failed to select from location by regionId")
		if rows != nil {
			var publicList []_struct.NcovPublic
			var tmpD _struct.NcovDistrict
			rows.Next()
			rows.Scan(&tmpD.Locale, &tmpD.Address, &tmpD.Lng, &tmpD.Lat, &tmpD.Source)

			rows, err = db.Query("SELECT personId, start, end, Date FROM publicRecord WHERE location=" + strconv.Itoa(locationId))
			errHandle.CheckErr(err, "failed to select from pbr by locationId")
			if rows != nil {
				for rows.Next() {
					var personId int
					var tmpP _struct.NcovPublic
					rows.Scan(&personId, &tmpP.Start, &tmpP.End, &tmpP.Date)
					if len(tmpD.Address) != 0 {
						tmpP.Address = tmpD.Address
					}
					rows2, err := db.Query("SELECT userName FROM user WHERE isSick=1 and userId=" + strconv.Itoa(personId))
					errHandle.CheckErr(err, "failed to select from user by id")
					if rows2 != nil {
						rows2.Next()
						rows2.Scan(&tmpP.PersonName)
						if len(tmpP.PersonName) == 0 {
							continue
						}
					}
					rows3, err := db.Query("SELECT address FROM location WHERE locationId=" + strconv.Itoa(locationId))
					errHandle.CheckErr(err, "failed to select from user by id")
					if rows3 != nil {
						rows3.Next()
						rows3.Scan(&tmpP.Address)
					}
					publicList = append(publicList, tmpP)
				}
			}
			tmpD.NcovPublicList = publicList
			districtList = append(districtList, tmpD)
		}
	}
	return districtList
}
