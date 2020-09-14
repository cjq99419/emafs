package get

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func NcovPublicReq(c *gin.Context) {
	cityName := c.Query("cityName")
	locale := c.Query("locale")
	district := dao.NcovPublicGet(cityName, locale)
	c.JSON(200, gin.H{
		"districtList": district,
	})

}
