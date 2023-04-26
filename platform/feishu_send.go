package platform

import (
	"github.com/championlong/robot-talk-sdk/config"
)

type FeishuMasterJob struct {
	Url       string                    `json:"-"` //请求url
	KindRobot string                    `json:"-"` //机器人种类
	Query     config.FeishuRobotsConfig `json:"-"`
}

func NewFeishuMasterJob() SendMessage {
	return &FeishuMasterJob{}
}

// SendMessage 发送消息
func (job *FeishuMasterJob) SendMessage(message interface{}) error {
	return nil
}
