package store

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"fmt"
	"strings"
)

func NcovDistrictStore(ncovs [40][40][40]_struct.NcovDistricts) int {
	db := tools.ConnectToDB()
	for _, ncov2 := range ncovs {
		if len(ncov2[0][0].Ncovs) != 0 {
			for _, ncov1 := range ncov2 {
				if len(ncov1[0].Ncovs) != 0 {
					for _, ncov0 := range ncov1 {
						if len(ncov0.Ncovs) != 0 {
							for _, ncov := range ncov0.Ncovs {
								if len(ncov.Address) == 0 {
									continue
								} else {
									stmt, err := db.Prepare(`REPLACE INTO location (locale, address, lng, lat, source, region) VALUES(?, ?, ?, ?, ?, ?)`)
									errHandle.CheckErr(err, "failed to db.Prepare!")
									res, err := stmt.Exec(ncov.Locale, ncov.Address, ncov.Lng, ncov.Lat, ncov.Source, ncov.Region)
									errHandle.CheckErr(err, "failed to stmt.Exec!")
									_, err = res.LastInsertId()
									errHandle.CheckErr(err, "failed to res.LastInsertId")
								}
							}
						}
					}
				}
			}
		}
	}
	_, err := db.Exec("delete from location where (lng, lat) in (select lng, lat from (select lng, lat from location group by lng, lat having count(*) > 1) temp) and locationId not in (select * from (select min(locationId) from location group by lng,lat having count(*)>1) temp)")
	errHandle.CheckErr(err, "failed to do duplicate removal")
	return 1
}

func splitAddress(address string) []string {
	var province []string
	var city []string
	var district []string
	var cnAddressSpl []string
	if strings.Contains(address, "北京") || strings.Contains(address, "上海") || strings.Contains(address, "天津") || strings.Contains(address, "重庆") {
		c := strings.Split(address, "市")
		province = []string{c[0], address}
	} else if strings.Contains(address, "省") {
		province = strings.Split(address, "省")
	}
	fmt.Println(province)
	if strings.Contains(province[1], "市") {
		city = strings.Split(province[1], "市")
	}
	fmt.Println(city)
	if strings.Contains(city[1], "市") {
		district = strings.Split(city[1], "市")
		cnAddressSpl = []string{province[0], city[0], district[0] + "市", district[1]}
	} else if strings.Contains(city[1], "区") {
		district = strings.Split(city[1], "区")
		cnAddressSpl = []string{province[0], city[0], district[0] + "区", district[1]}
	} else if strings.Contains(city[1], "县") {
		district = strings.Split(city[1], "县")
		cnAddressSpl = []string{province[0], city[0], district[0] + "县", district[1]}
	} else if strings.Contains(city[1], "街道") {
		district = strings.Split(city[1], "街道")
		cnAddressSpl = []string{province[0], city[0], district[0] + "街道", district[1]}
	}
	fmt.Println(district)
	return cnAddressSpl
}
