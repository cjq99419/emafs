package dao

import (
	errHandle "emafs/error"
	_struct "emafs/struct"
	"emafs/tools"
)

func NewsGet() []_struct.DailyNews {
	db := tools.ConnectToDB()
	rows, err := db.Query(`SELECT pubDate, pubDateStr, title, summary, source, url FROM news ORDER BY newsId DESC LIMIT 5`)
	errHandle.CheckErr(err, "failed to getDesc in descGet")
	var news [5]_struct.DailyNews
	i := 0
	for rows.Next() {
		err := rows.Scan(&news[i].PubDate, &news[i].PubDateStr, &news[i].Title, &news[i].Summary, &news[i].InfoSource, &news[i].SourceURL)
		errHandle.CheckErr(err, "failed to ScanNews")
		i++
	}
	return news[:]
}
