package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"fmt"
)

func NcovRideGet(trainNum, date string) []_struct.NcovRide {
	db := tools.ConnectToDB()
	var rideLists []_struct.NcovRide
	sql := "SELECT passengerId, date, start, end, trainNum, carriageNum, seatNum, posStart, posEnd " +
		"FROM rideRecord WHERE trainNum='" + trainNum + "' and date='" + date + "'"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	errHandle.CheckErr(err, "failed to get rideRecord by trainNum and date")
	if rows != nil {
		for rows.Next() {
			var tmpR _struct.NcovRide
			rows.Scan(&tmpR.PassengerName, &tmpR.Date, &tmpR.Start, &tmpR.End, &tmpR.TrainNum, &tmpR.CarriageNum,
				&tmpR.SeatNum, &tmpR.PosStart, &tmpR.PosEnd)
			row, err := db.Query("SELECT userName FROM user WHERE isSick=1 and userId='" + tmpR.PassengerName + "'")
			errHandle.CheckErr(err, "can not get name of user")
			if row != nil {
				row.Next()
				row.Scan(&tmpR.PassengerName)
			} else {
				tmpR.PassengerName = ""
			}
			if len(tmpR.PassengerName) > 2 {
				rideLists = append(rideLists, tmpR)
			}
		}
	}
	return rideLists
}
