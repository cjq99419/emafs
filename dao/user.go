package dao

import (
	errHandle "emafs/error"
	"emafs/tools"
	"fmt"
)

func GetPasswordByPN(phoneNum string) (int, string) {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT password FROM user WHERE phoneNum='" + phoneNum + "'")
	errHandle.CheckErr(err, "failed to get pwd by phoneNum")
	if rows == nil {
		return 0, ""
	}
	rows.Next()
	var pwd string
	rows.Scan(&pwd)
	return 1, pwd
}

func GetPasswordByID(id string) (int, string) {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT password FROM user WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to get pwd by id")
	if rows == nil {
		return 0, ""
	}
	rows.Next()
	var pwd string
	rows.Scan(&pwd)
	return 1, pwd
}

func UpdatePhoneByID(id, phoneNum string) int {
	db := tools.ConnectToDB()
	fmt.Println(id)
	phone := GetPhoneByID(id)
	fmt.Println(phone)
	res, err := db.Exec("UPDATE user SET phoneNum='" + phoneNum + "' WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to update phone of user")
	if res == nil {
		return 0
	}
	aff, err := res.RowsAffected()
	if aff == 0 && len(phone) == 0 {
		return 0
	} else {
		return 1
	}
}

func GetIDByPhone(phoneNum string) string {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT identityNum FROM user WHERE phoneNum='" + phoneNum + "'")
	errHandle.CheckErr(err, "failed to get pwd by id")
	if rows == nil {
		return ""
	}
	rows.Next()
	var identityNum string
	rows.Scan(&identityNum)
	return identityNum
}

func GetPhoneByID(id string) string {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT phoneNum FROM user WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to get pwd by id")
	if rows == nil {
		return ""
	}
	rows.Next()
	var phoneNum string
	rows.Scan(&phoneNum)
	return phoneNum
}

func GetIsSickByID(id string) int {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT isSick FROM user WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to get pwd by id")
	if rows == nil {
		return -1
	}
	rows.Next()
	var isSick int
	rows.Scan(&isSick)
	return isSick
}

func GetUserIdByIdentityNum(id string) int {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT userId FROM user WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to get pwd by id")
	if rows == nil {
		return -1
	}
	rows.Next()
	var userId int
	rows.Scan(&userId)
	return userId
}

func UpdateIsSickByID(id, date string) int {
	db := tools.ConnectToDB()
	isSick := GetIsSickByID(id)
	if isSick == 1 {
		return 0
	}
	fmt.Println(id)
	res, err := db.Exec("UPDATE user SET isSick=1, confirmTime='" + date + "' WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to update user by id")
	if res == nil {
		return 1
	}
	aff, err := res.RowsAffected()
	if aff == 0 {
		return 1
	} else {
		return 2
	}
}

func GetNameByID(id string) string {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT userName FROM user WHERE identityNum='" + id + "'")
	errHandle.CheckErr(err, "failed to get userName by id")
	if rows == nil {
		return ""
	}
	rows.Next()
	var userName string
	rows.Scan(&userName)
	return userName
}
