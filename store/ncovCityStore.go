package store

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
)

func NcovCityStore(ncovs []_struct.NcovCity) int {
	db := tools.ConnectToDB()
	stmt, err := db.Prepare(`UPDATE region set currentConfirmedCount=?, suspectedCount=?, ConfirmedCount=?, 
										CuredCount=?, DeadCount=? WHERE name=?`)
	errHandle.CheckErr(err, "failed to db.Prepare!")
	for _, ncov := range ncovs {
		res, err := stmt.Exec(ncov.CurrentConfirmedCount, ncov.SuspectedCount, ncov.ConfirmedCount, ncov.CuredCount, ncov.DeadCount, ncov.Name)
		errHandle.CheckErr(err, "failed to stmt.Exec!")
		_, err = res.LastInsertId()
		errHandle.CheckErr(err, "failed to res.LastInsertId")
	}
	return 1
}
