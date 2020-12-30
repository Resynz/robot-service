/**
 * @Author: Resynz
 * @Date: 2020/12/30 11:01
 */
package model

import (
	"robot-service/db/enums"
	"time"
)

type Robot struct {
	Id         int64             `json:"id"`
	Name       string            `json:"name"`
	Type       enums.RobotType   `json:"type"`
	Status     enums.RobotStatus `json:"status"`
	Webhook    string            `json:"webhook"`
	Secret     string            `json:"secret"`
	CreateTime time.Time         `json:"create_time"`
	UpdateTime time.Time         `json:"update_time"`
}
