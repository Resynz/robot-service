/**
 * @Author: Resynz
 * @Date: 2020/12/30 16:45
 */
package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	db_handler "github.com/yue-best-practices/db-handler"
	"net/http"
	"robot-service/db"
	"robot-service/db/enums"
	"robot-service/db/model"
	"robot-service/lib"
	"strconv"
	"time"
)

func AddRobot(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	var robot model.Robot

	if err = json.Unmarshal(body, &robot); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	if robot.Type != enums.DingTalkRobot {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "invalid robot type",
		})
		return
	}
	if robot.Webhook == "" {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "invalid webhook",
		})
		return
	}

	if robot.Secret == "" {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  "invalid secret",
		})
		return
	}
	robot.Status = enums.RobotStatusEnable
	robot.CreateTime = time.Now()
	robot.UpdateTime = time.Now()

	if err = db.Handler.Save(&robot, "robot"); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
		"data": map[string]int64{
			"robot_id": robot.Id,
		},
	})

}

func ListRobot(ctx *gin.Context) {
	page := 1
	if c, err := strconv.Atoi(ctx.Query("page")); err == nil && c > 1 {
		page = c
	}
	pageSize := 10
	if s, err := strconv.Atoi(ctx.Query("page_size")); err == nil && s > 0 {
		pageSize = s
	}

	var list []model.Robot
	condition := &db_handler.Condition{
		Where:  "",
		Params: []interface{}{},
		Asc:    []string{"id"},
		Offset: (page - 1) * pageSize,
		Limit:  pageSize,
	}

	if ctx.Query("name") != "" {
		condition.Where = fmt.Sprintf("name like '%%%s%%'", ctx.Query("title"))
	}

	err := db.Handler.List(&list, "robot", condition)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	total, err := db.Handler.Count("robot", &db_handler.Condition{Where: condition.Where, Params: condition.Params})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	pagination := map[string]interface{}{
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}
	data := map[string]interface{}{
		"list":       list,
		"pagination": pagination,
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
		"data": data,
	})
}

func RobotInfo(ctx *gin.Context) {
	robotId := ctx.Param("id")
	var robot model.Robot
	has, err := db.Handler.Get(&robot, "robot", robotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  fmt.Sprintf("robot[%s] is not exists", robotId),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
		"data": map[string]model.Robot{
			"robot": robot,
		},
	})
}

func DeleteRobot(ctx *gin.Context) {
	robotId := ctx.Param("id")
	var robot model.Robot
	has, err := db.Handler.Get(&robot, "robot", robotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  fmt.Sprintf("robot[%s] is not exists", robotId),
		})
		return
	}
	if err = db.Handler.Del(&robot, "robot", robotId); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	})
}

func ManageRobot(ctx *gin.Context) {
	robotId := ctx.Param("id")
	_status := ctx.Param("status")
	status := enums.RobotStatusEnable
	if _status == "disable" {
		status = enums.RobotStatusDisable
	}
	var robot model.Robot
	has, err := db.Handler.Get(&robot, "robot", robotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  fmt.Sprintf("robot[%s] is not exists", robotId),
		})
	}

	robot.Status = status
	robot.UpdateTime = time.Now()
	if err = db.Handler.Save(&robot, "robot"); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	})
}

func SayHello(ctx *gin.Context) {
	robotId := ctx.Param("id")
	var robot model.Robot
	has, err := db.Handler.Get(&robot, "robot", robotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  fmt.Sprintf("robot[%s] is not exists", robotId),
		})
		return
	}

	dingTalkRobot := &lib.DingTalkRobot{
		Name:    robot.Name,
		Webhook: robot.Webhook,
		Secret:  robot.Secret,
	}
	if err = dingTalkRobot.SayHello(); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	})
}

func Remind(ctx *gin.Context) {
	robotId := ctx.Param("id")
	var robot model.Robot
	has, err := db.Handler.Get(&robot, "robot", robotId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	if !has {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  fmt.Sprintf("robot[%s] is not exists", robotId),
		})
		return
	}

	dingTalkRobot := &lib.DingTalkRobot{
		Name:    robot.Name,
		Webhook: robot.Webhook,
		Secret:  robot.Secret,
	}

	if err = dingTalkRobot.Remind(lib.RemindType(ctx.Query("type")), ctx.Query("mobile")); err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code": 400,
			"msg":  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	})
}
