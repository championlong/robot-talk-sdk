package dingding

import (
	"github.com/championlong/dingtalk-sdk/config"
	"github.com/championlong/dingtalk-sdk/model/ding_talk"
	"github.com/championlong/dingtalk-sdk/platform"
)

var platformConfig *config.PlatformConfig

// GetConfig 获取钉钉配置文件
func GetConfig() *config.PlatformConfig {
	return platformConfig
}

// Init 初始化钉钉配置
func Init(config config.PlatformConfig) {
	platformConfig = &config
}

// SendDingMessage 发送钉钉消息
func SendDingMessage(kindRobot string, messageType platform.MsgType, message interface{}, at ding_talk.At) error {
	return platform.NewDingMasterJob(kindRobot, messageType, platformConfig.DingRobots[kindRobot], at).SendMessage(message)
}

// SendFeishuMessage 发送飞书消息
func SendFeishuMessage(kindRobot string, messageType platform.MsgType, message interface{}, at ding_talk.At) error {
	job := &platform.DingMasterJob{}
	job.KindRobot = kindRobot
	job.Url = ""
	job.Msgtype = messageType
	job.Query = platformConfig.DingRobots[kindRobot]
	job.At = at
	return job.SendMessage(message)
}
