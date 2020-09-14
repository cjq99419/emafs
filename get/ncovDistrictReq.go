package get

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func NcovDistrictReq(c *gin.Context) {
	name := c.Query("name")
	ncovDistricts := dao.NcovDistrictGet(name)
	c.JSON(200, gin.H{"districtList": ncovDistricts})
}
