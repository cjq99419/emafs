package post

import (
	"emafs/broadcast"
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func UpdateIsSick(c *gin.Context) {
	id := c.Query("id")
	date := c.Query("date")
	state := dao.UpdateIsSickByID(id, date)
	switch state {
	case 0:
		broadcast.NcovBroadcast(id)
		c.JSON(200, gin.H{
			"msg": "已经确诊",
		})
	case 1:
		c.JSON(200, gin.H{
			"msg": "查无id",
		})
	case 2:
		broadcast.NcovBroadcast(id)
		c.JSON(200, gin.H{
			"msg": "成功",
		})
	}

}
