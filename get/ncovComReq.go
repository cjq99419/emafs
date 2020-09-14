package get

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func NcovComReq(c *gin.Context) {
	id := c.Query("id")
	ncovCom := dao.NcovComGet(id)
	c.JSON(200, gin.H{
		"ncovCom": ncovCom,
	})
}
