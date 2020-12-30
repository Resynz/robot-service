/**
 * @Author: Resynz
 * @Date: 2020/12/30 11:05
 */
package enums

// 机器人类型
type RobotType string

const (
	DingTalkRobot RobotType = "DingTalk" // 钉钉机器人
	// todo more
)

// 机器人状态
type RobotStatus uint8

const (
	RobotStatusDisable RobotStatus = iota
	RobotStatusEnable
)
