package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
)

func DescGet() _struct.DailyDesc {
	db := tools.ConnectToDB()
	rows, err := db.Query(`SELECT currentConfirmedCount, confirmedCount, suspectedCount, curedCount,
 							deadCount, seriousCount, suspectedIncr, currentConfirmedIncr,confirmedIncr, 
 							curedIncr, deadIncr, seriousIncr FROM description ORDER BY id DESC LIMIT 1`)
	errHandle.CheckErr(err, "failed to getDesc in descGet")
	rows.Next()
	var desc _struct.DailyDesc
	err = rows.Scan(&desc.CurrentConfirmedCount, &desc.ConfirmedCount, &desc.SuspectedCount, &desc.CuredCount,
		&desc.DeadCount, &desc.SeriousCount, &desc.SuspectedIncr, &desc.CurrentConfirmedIncr,
		&desc.ConfirmedIncr, &desc.CuredIncr, &desc.DeadIncr, &desc.SeriousIncr)
	return desc
}
