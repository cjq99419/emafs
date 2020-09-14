package tools

import (
	"math"
	"time"
)

func IsInFTDays(date string) bool {
	timeTemplate := "2006-01-02"

	stamp, _ := time.ParseInLocation(timeTemplate, date, time.Local)
	sec := math.Abs(float64(stamp.Unix() - time.Now().Unix()))
	if sec <= 1209600 {
		return true
	} else {
		return false
	}

}
