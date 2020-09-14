package get

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func NcovCityReq(c *gin.Context) {
	name := c.Query("name")
	ncov := dao.NcovCityGet(name)
	c.JSON(200, ncov)
}
