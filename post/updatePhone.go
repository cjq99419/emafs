package post

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func UpdatePhone(c *gin.Context) {
	id := c.Query("id")
	phoneNum := c.Query("phoneNum")
	if len(id) == 0 || len(phoneNum) == 0 {
		c.JSON(200, map[string]int{
			"state": 2,
		})
		return
	}
	state := dao.UpdatePhoneByID(id, phoneNum)
	if state == 0 {
		c.JSON(200, map[string]int{
			"state": 0,
		})
	} else {
		c.JSON(200, map[string]int{
			"state": 1,
		})
	}
}
