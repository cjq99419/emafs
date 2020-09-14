package get

import (
	"emafs/dao"
	_struct "emafs/struct"
	"github.com/gin-gonic/gin"
)

func DailyNewsReq(c *gin.Context) {
	var dailyNews []_struct.DailyNews
	dailyNews = dao.NewsGet()
	c.JSON(200, gin.H{"news": dailyNews})
}
