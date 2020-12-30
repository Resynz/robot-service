/**
 * @Author: Resynz
 * @Date: 2020/12/30 10:57
 */
package main

import (
	"log"
	"robot-service/config"
	"robot-service/db"
	"robot-service/server"
)

func main() {
	if err := config.CheckEnvConf(); err != nil {
		log.Fatalf("check env config failed! error:%v\n", err)
	}
	if err := db.InitDBHandler(); err != nil {
		log.Fatalf("init db handler failed! error:%v\n", err)
	}
	log.Println("\033[42;30m DONE \033[0m[RobotService] Start Success!")
	server.StartService()
}
