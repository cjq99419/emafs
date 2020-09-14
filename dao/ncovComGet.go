package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"strconv"
)

func NcovComGet(id string) _struct.NcovCom {
	var ncovCom _struct.NcovCom
	ncovCom.DistrictList = GetDistrictListById(id)
	ncovCom.RideList = GetRideListById(id)
	return ncovCom
}

func GetRideListById(id string) []_struct.NcovRide {
	userId := GetUserIdByIdentityNum(id)
	db := tools.ConnectToDB()
	var rideList []_struct.NcovRide
	rows, err := db.Query("SELECT trainNum, date FROM rideRecord WHERE passengerId=" + strconv.Itoa(userId))
	errHandle.CheckErr(err, "failed to get trainNum and date from rideRecord")
	if rows != nil {
		for rows.Next() {
			var trainNum, date string
			rows.Scan(&trainNum, &date)
			rows1, err := db.Query("SELECT passengerId, date, start, end, trainNum, carriageNum, seatNum, posStart, posEnd" +
				" FROM rideRecord WHERE trainNum='" + trainNum + "' and date='" + date +
				"' and passengerId in (SELECT userId FROM user WHERE isSick=1)")
			errHandle.CheckErr(err, "failed to get rideInf by trainNum and date")
			if rows1 != nil {
				for rows1.Next() {
					var tmpR _struct.NcovRide
					rows1.Scan(&tmpR.PassengerName, &tmpR.Date, &tmpR.Start, &tmpR.End, &tmpR.TrainNum, &tmpR.CarriageNum,
						&tmpR.SeatNum, &tmpR.PosStart, &tmpR.PosEnd)
					rows2, err := db.Query("SELECT userName FROM user WHERE userId='" + tmpR.PassengerName + "'")
					errHandle.CheckErr(err, "can not get name of user")
					if rows2 != nil {
						rows2.Next()
						rows2.Scan(&tmpR.PassengerName)
					}
					rideList = append(rideList, tmpR)
				}
			}
		}
	}
	return rideList
}

func GetDistrictListById(id string) []_struct.NcovDistrict {
	var ncovDistrictList []_struct.NcovDistrict
	userId := GetUserIdByIdentityNum(id)
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT DISTINCT location, date FROM publicRecord WHERE personId=" + strconv.Itoa(userId))
	errHandle.CheckErr(err, "failed to get locationId from publicRecord by person id")
	if rows != nil {
		for rows.Next() {
			var locationId int
			var date string
			rows.Scan(&locationId, &date)
			rows1, err := db.Query("SELECT locale, address, lng, lat, source FROM location where locationId=" + strconv.Itoa(locationId))
			errHandle.CheckErr(err, "failed to get locationInf by id")
			if rows1 != nil {
				for rows1.Next() {
					var tmpD _struct.NcovDistrict
					rows1.Scan(&tmpD.Locale, &tmpD.Address, &tmpD.Lng, &tmpD.Lat, &tmpD.Source)
					var NcovPublicList []_struct.NcovPublic
					rows2, err := db.Query("SELECT personId, start, end, Date FROM publicRecord WHERE location=" + strconv.Itoa(locationId) +
						" and Date='" + date + "'")
					errHandle.CheckErr(err, "can not get publicRecord by location and date")
					if rows2 != nil {
						for rows2.Next() {
							var tmpP _struct.NcovPublic
							var personId int
							tmpP.Address = tmpD.Address
							rows2.Scan(&personId, &tmpP.Start, &tmpP.End, &tmpP.Date)
							rows3, err := db.Query("SELECT userName FROM user WHERE isSick=1 and userId=" + strconv.Itoa(personId))
							errHandle.CheckErr(err, "failed to select from user by id")
							if rows3 != nil {
								rows3.Next()
								rows3.Scan(&tmpP.PersonName)
							}
							if len(tmpP.PersonName) != 0 {
								NcovPublicList = append(NcovPublicList, tmpP)
							}
						}
					}
					tmpD.NcovPublicList = NcovPublicList
					ncovDistrictList = append(ncovDistrictList, tmpD)
				}
			}
		}
	}
	return ncovDistrictList
}
