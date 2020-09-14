package main

import (
	"emafs/config"
	"emafs/dao"
	"emafs/get"
	"emafs/post"
	"emafs/timer"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	dao.StoreRegion()
	timer.TickerGetDailyDesc()
	timer.TickerGetDailyNews()
	timer.TickerGetNcovCity()
	timer.TickerGetNcovDistrict()
	r.GET("/emafs/DailyDesc", get.DailyDescReq)
	r.GET("/emafs/DailyNews", get.DailyNewsReq)
	r.GET("/emafs/NcovCity", get.NcovCityReq)
	r.GET("/emafs/NcovDistrict", get.NcovDistrictReq)
	r.POST("/emafs/Login", post.Login)
	r.POST("/emafs/UpdatePhone", post.UpdatePhone)
	r.GET("/emafs/NcovPublic", get.NcovPublicReq)
	r.GET("/emafs/NcovRide", get.NcovRideReq)
	r.GET("/emafs/NcovCom", get.NcovComReq)
	r.POST("/emafs/UpdateIsSick", post.UpdateIsSick)
	r.Run(":" + config.ServerPort)

}
