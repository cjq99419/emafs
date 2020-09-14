package broadcast

import (
	"emafs/dao"
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"strconv"
)

func NcovBroadcast(id string) {
	name := dao.GetNameByID(id)
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT trainNum, date FROM rideRecord WHERE passengerId=(SELECT userId FROM user WHERE identityNum='" + id + "')")
	errHandle.CheckErr(err, "failed to get trainNum and date in broadcast")
	if rows != nil {
		for rows.Next() {
			var trainNum, date string
			rows.Scan(&trainNum, &date)
			rows1, err := db.Query("SELECT identityNum FROM user WHERE userId in (SELECT passengerId FROM rideRecord WHERE date='" + date + "' and trainNum='" + trainNum + "')")
			errHandle.CheckErr(err, "failed to get passengerId in broadcast")
			if rows1 != nil {
				for rows1.Next() {
					var push _struct.Push
					var uid string
					rows1.Scan(&uid)
					if uid == id {
						continue
					}
					phoneNum := dao.GetPhoneByID(uid)
					rows2, err := db.Query("SELECT carriageNum, seatNum, posStart, posEnd FROM rideRecord WHERE date='" + date + "' and trainNum='" + trainNum + "' and passengerId=" +
						"(SELECT userId FROM user WHERE identityNum='" + uid + "')")
					errHandle.CheckErr(err, "failed to get p info")
					if rows2 != nil {
						for rows2.Next() {
							var carriageNum, seatNum, posStart, posEnd string
							rows2.Scan(&carriageNum, &seatNum, &posStart, &posEnd)
							push = _struct.Push{
								UserAccount: phoneNum,
								Title:       "注意！您同乘列车出现新增确诊病例！！！",
								Description: "您于" + date + "日乘坐的列车" + trainNum + "出现一例新增确诊病例\n" +
									"乘客" + name + "于" + date + "日乘坐列车" + trainNum + "，乘车期间就坐于" +
									carriageNum + "号车厢" + seatNum + "号座位，乘车区间为" + posStart + "站到" +
									posEnd + "站\n请您及时注意身体情况，如有不适尽快就医！",
							}
						}
					}
					if tools.IsInFTDays(date) {
						tools.HttpReqPostPush(push)
					}
				}
			}
		}
	}
	rows, err = db.Query("SELECT location, Date FROM publicRecord WHERE " +
		"personId=(SELECT userId FROM user WHERE identityNum='" + id + "')")
	errHandle.CheckErr(err, "can not get locationId and date from pr")
	if rows != nil {
		for rows.Next() {
			var locationId int
			var date string
			rows.Scan(&locationId, &date)
			address := dao.GetAddressByLocationId(locationId)
			rows1, err := db.Query("SELECT personId FROM publicRecord WHERE location=" + strconv.Itoa(locationId) +
				" and Date='" + date + "' and personId in " + "(SELECT userId FROM user WHERE isSick=0)")
			errHandle.CheckErr(err, "failed to get person id from pr")
			if rows1 != nil {
				for rows1.Next() {
					var personId int
					var push _struct.Push
					rows1.Scan(&personId)
					rows2, err := db.Query("SELECT phoneNum FROM user WHERE userId=" + strconv.Itoa(personId))
					errHandle.CheckErr(err, "failed to select phoneNum from user")
					if rows2 != nil {
						rows2.Next()
						var phoneNum string
						rows2.Scan(&phoneNum)
						rows3, _ := db.Query("SELECT start, end FROM publicRecord WHERE personId=" + strconv.Itoa(personId) + " and location=" + strconv.Itoa(locationId))
						if rows3 != nil {
							rows3.Next()
							var start, end string
							rows3.Scan(&start, &end)
							push = _struct.Push{
								UserAccount: phoneNum,
								Title:       "注意！您曾去往的公共场所出现新增确诊病例！！！",
								Description: "您于" + date + "日前往的" + address + "出现一个新的确诊病例\n" +
									name + "于" + date + "日" + start + "至" + end + "时间段在此活动\n" +
									"请您及时注意身体情况，如有不适尽快就医！",
							}
						}
					}
					tools.HttpReqPostPush(push)
				}
			}
		}
	}
}
