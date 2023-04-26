package dingding

import (
	"github.com/championlong/dingtalk-sdk/config"
	"github.com/championlong/dingtalk-sdk/model/ding_talk"
	"github.com/championlong/dingtalk-sdk/platform"
)

// GetConfig 获取平台机器人配置文件
func GetConfig() *config.PlatformConfig {
	return config.PlatformConfigGlobal
}

// Init 初始化平台机器人配置文件
func Init(global config.PlatformConfig) {
	config.PlatformConfigGlobal = &global
}

// SendDingMessage 发送钉钉消息
func SendDingMessage(kindRobot string, messageType platform.MsgType, message interface{}, at ding_talk.At) error {
	return platform.NewDingMasterJob(kindRobot, messageType, at).SendMessage(message)
}

// SendFeishuMessage 发送飞书消息
func SendFeishuMessage(message interface{}) error {
	return platform.NewFeishuMasterJob().SendMessage(message)
}

// SendRobotMessage 自己初始化选择平台发送
func SendRobotMessage(send platform.SendMessage, message interface{}) error {
	return send.SendMessage(message)
}
