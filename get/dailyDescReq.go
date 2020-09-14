package get

import (
	"emafs/dao"
	"github.com/gin-gonic/gin"
)

func DailyDescReq(c *gin.Context) {
	dailyDesc := dao.DescGet()
	c.JSON(200, dailyDesc)
}
