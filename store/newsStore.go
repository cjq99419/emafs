package store

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
)

func NewsStore(news _struct.DailyNews) int {
	db := tools.ConnectToDB()
	stmt, err := db.Prepare(`INSERT INTO news (pubDate, pubDateStr, title, summary, source, url) VALUES (?, ?, ?, ?, ?, ?)`)
	errHandle.CheckErr(err, "failed to db.Prepare!")
	res, err := stmt.Exec(news.PubDate, news.PubDateStr, news.Title, news.Summary, news.InfoSource, news.SourceURL)
	errHandle.CheckErr(err, "failed to stmt.Exec!")
	id, err := res.LastInsertId()
	errHandle.CheckErr(err, "failed to res.LastInsertId")
	return int(id)
}

func NewsClear() {
	db := tools.ConnectToDB()
	_, err := db.Exec(`DELETE FROM news WHERE newsId>?`, 0)
	errHandle.CheckErr(err, "failed to stmt.Exec!")
	db.Close()
}
