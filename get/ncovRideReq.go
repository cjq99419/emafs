package get

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func NcovRideReq(c *gin.Context) {
	trainNum := c.Query("trainNum")
	date := c.Query("date")
	rideList := dao.NcovRideGet(trainNum, date)
	c.JSON(200, gin.H{
		"ride_list": rideList,
	})
}
