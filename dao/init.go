package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
)

func StoreRegion() {
	db := tools.ConnectToDB()
	rows, err := db.Query("SELECT name, id FROM region")
	errHandle.CheckErr(err, "failed to store region")
	var name string
	var id int
	for rows.Next() {
		rows.Scan(&name, &id)
		_struct.Regions[id] = name
	}
	rows, err = db.Query("SELECT field_id, id FROM region")
	errHandle.CheckErr(err, "failed to store region")
	var fieldId int
	for rows.Next() {
		rows.Scan(&fieldId, &id)
		_struct.RegionsIdField[id] = fieldId
	}
	//fmt.Println(_struct.RegionsIdField)
	//fmt.Println(_struct.Regions)
}
