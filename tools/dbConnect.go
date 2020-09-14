package tools

import (
	"database/sql"
	"emafs/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open(config.DBDriverName, config.DBDataSource)
	if err != nil {
		fmt.Printf("connect mysql failed! [%s]", err)
	} else {
		return db
	}
	return nil
}
