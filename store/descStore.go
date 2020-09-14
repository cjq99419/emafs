package store

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
	"time"
)

func DescStore(dailyDesc _struct.DailyDesc) int {
	db := tools.ConnectToDB()
	rows, err := db.Query(`SELECT currentConfirmedCount, confirmedCount, suspectedCount, curedCount,
 							deadCount, seriousCount, suspectedIncr, currentConfirmedIncr,confirmedIncr, 
 							curedIncr, deadIncr, seriousIncr FROM description ORDER BY id DESC LIMIT 1`)
	errHandle.CheckErr(err, "failed to select from desc in descStore")
	var preDesc _struct.DailyDesc
	rows.Next()
	err = rows.Scan(&preDesc.CurrentConfirmedCount, &preDesc.ConfirmedCount, &preDesc.SuspectedCount, &preDesc.CuredCount,
		&preDesc.DeadCount, &preDesc.SeriousCount, &preDesc.SuspectedIncr, &preDesc.CurrentConfirmedIncr,
		&preDesc.ConfirmedIncr, &preDesc.CuredIncr, &preDesc.DeadIncr, &preDesc.SeriousIncr)
	errHandle.CheckErr(err, "failed to scan rows to struct preDesc")
	if dailyDesc.SuspectedIncr != -12345 {
	} else if dailyDesc.SuspectedCount != preDesc.SuspectedCount {
		dailyDesc.SuspectedIncr = dailyDesc.SuspectedCount - preDesc.SuspectedCount
	} else {
		dailyDesc.SuspectedIncr = preDesc.SuspectedIncr
	}
	if dailyDesc.CuredIncr != -12345 {
	} else if dailyDesc.CuredCount != preDesc.CuredCount {
		dailyDesc.CuredIncr = dailyDesc.CuredCount - preDesc.CuredCount
	} else {
		dailyDesc.CuredIncr = preDesc.CuredIncr
	}
	if dailyDesc.CurrentConfirmedIncr != -12345 {
	} else if dailyDesc.CurrentConfirmedCount != preDesc.CurrentConfirmedCount {
		dailyDesc.CurrentConfirmedIncr = dailyDesc.CurrentConfirmedCount - preDesc.CurrentConfirmedCount
	} else {
		dailyDesc.CurrentConfirmedIncr = preDesc.CurrentConfirmedIncr
	}
	if dailyDesc.ConfirmedIncr != -12345 {
	} else if dailyDesc.ConfirmedCount == preDesc.ConfirmedCount {
		dailyDesc.ConfirmedIncr = dailyDesc.ConfirmedCount - preDesc.ConfirmedCount
	} else {
		dailyDesc.ConfirmedIncr = preDesc.ConfirmedIncr
	}
	if dailyDesc.DeadIncr != -12345 {
	} else if dailyDesc.DeadCount == preDesc.DeadCount {
		dailyDesc.DeadIncr = dailyDesc.DeadCount - preDesc.DeadCount
	} else {
		dailyDesc.DeadIncr = preDesc.DeadIncr
	}
	if dailyDesc.SeriousIncr != -12345 {
	} else if dailyDesc.SeriousCount == preDesc.SeriousCount {
		dailyDesc.SeriousIncr = dailyDesc.SeriousCount - preDesc.SeriousCount
	} else {
		dailyDesc.SeriousIncr = preDesc.SeriousIncr
	}

	stmt, err := db.Prepare(`INSERT description (currentConfirmedCount, confirmedCount, suspectedCount, curedCount,
												deadCount, seriousCount, suspectedIncr, currentConfirmedIncr,
												confirmedIncr, curedIncr, deadIncr, seriousIncr, createTime) VALUES (
												?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
												)`)
	errHandle.CheckErr(err, "failed to db.Prepare!")
	res, err := stmt.Exec(dailyDesc.CurrentConfirmedCount, dailyDesc.ConfirmedCount, dailyDesc.SuspectedCount, dailyDesc.CuredCount,
		dailyDesc.DeadCount, dailyDesc.SeriousCount, dailyDesc.SuspectedIncr, dailyDesc.CurrentConfirmedIncr,
		dailyDesc.ConfirmedIncr, dailyDesc.CuredIncr, dailyDesc.DeadIncr, dailyDesc.SeriousIncr, time.Now())
	errHandle.CheckErr(err, "failed to stmt.Exec!")
	id, err := res.LastInsertId()
	errHandle.CheckErr(err, "failed to res.LastInsertId")
	return int(id)
}
