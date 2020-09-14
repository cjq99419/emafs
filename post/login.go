package post

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
	"strings"
)

func Login(c *gin.Context) {
	loginType := c.Query("type")
	password := c.Query("password")
	if len(loginType) == 0 || len(password) == 0 {
		c.JSON(200, gin.H{
			"state": 3,
		})
	} else if loginType == "0" {
		phoneNum := c.Query("phoneNum")
		if len(phoneNum) == 0 {
			c.JSON(200, gin.H{
				"state": 3,
			})
		}
		state, pwd := dao.GetPasswordByPN(phoneNum)
		if state == 0 {
			c.JSON(200, gin.H{
				"state": 0,
			})
		} else if strings.Compare(password, pwd) == 0 {
			id := dao.GetIDByPhone(phoneNum)
			c.JSON(200, gin.H{
				"state":       2,
				"identityNum": id,
				"phoneNum":    phoneNum,
			})
		} else {
			c.JSON(200, gin.H{
				"state": 1,
			})
		}
	} else if loginType == "1" {
		id := c.Query("id")
		if len(id) == 0 {
			c.JSON(200, gin.H{
				"state": 3,
			})
		}
		state, pwd := dao.GetPasswordByID(id)
		if state == 0 {
			c.JSON(200, gin.H{
				"state": 0,
			})
		} else if strings.Compare(password, pwd) == 0 {
			phoneNum := dao.GetPhoneByID(id)
			c.JSON(200, gin.H{
				"state":       2,
				"identityNum": id,
				"phoneNum":    phoneNum,
			})
		} else {
			c.JSON(200, gin.H{
				"state": 1,
			})
		}
	} else {
		c.JSON(200, gin.H{
			"state": 3,
		})
	}
}
