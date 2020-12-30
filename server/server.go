/**
 * @Author: Resynz
 * @Date: 2020/12/30 16:43
 */
package server

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"robot-service/config"
	"robot-service/controller"
)

func StartService() {
	gin.SetMode(config.Mode)
	app := gin.New()
	app.MaxMultipartMemory = 8 << 20 // 8mb
	app.Use(gzip.Gzip(gzip.DefaultCompression))
	app.GET("/ping", controller.Ping)
	app.POST("/robot", controller.AddRobot)
	app.GET("/robot", controller.ListRobot)
	app.GET("/robot/:id", controller.RobotInfo)
	app.DELETE("/robot/:id", controller.DeleteRobot)
	app.GET("/robot/:id/manage/:status", controller.ManageRobot)
	app.GET("/robot/:id/sayHello", controller.SayHello)
	app.GET("/robot/:id/remind", controller.Remind)

	if err := app.Run(fmt.Sprintf(":%d", config.AppPort)); err != nil {
		log.Fatalf("start server failed! error:%v\n", err)
	}
}
