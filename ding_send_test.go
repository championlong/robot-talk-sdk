package dingding

import (
	"fmt"
	"github.com/championlong/robot-talk-sdk/model/ding_talk"
	"github.com/championlong/robot-talk-sdk/platform"
	"testing"

	"github.com/championlong/robot-talk-sdk/config"
)

func initDingConfig() {
	query := make(map[string]config.DingRobotsConfig)
	query["report"] = config.DingRobotsConfig{
		Encrypt:     "",
		AccessToken: "",
	}
	Init(config.PlatformConfig{
		DingRobots: query,
	})
}

func TestSendText(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", platform.MsgTypeText, ding_talk.TextMessage{
		Content: "text",
	}, ding_talk.At{})
	fmt.Println(err)
}

func TestSendMarkdown(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", platform.MsgTypeMarkdown, ding_talk.MarkdownMessage{
		Title: "测试",
		Text:  "# 测试",
	}, ding_talk.At{})
	fmt.Println(err)
}

func TestSendLink(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", platform.MsgTypeLink, ding_talk.LinkMessage{
		Title:      "测试",
		Text:       "测试文本",
		PicURL:     "",
		MessageURL: "https://open.dingtalk.com/document/group/custom-robot-access",
	}, ding_talk.At{})
	fmt.Println(err)
}

func TestSendFeedCard(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", platform.MsgTypeFeedCard, ding_talk.FeedCardMessage{
		Links: []ding_talk.Links{
			{
				Title:      "测试1",
				PicURL:     "",
				MessageURL: "https://open.dingtalk.com/document/group/custom-robot-access",
			},
			{
				Title:      "测试2",
				PicURL:     "",
				MessageURL: "https://open.dingtalk.com/document/group/custom-robot-access",
			},
		},
	}, ding_talk.At{})
	fmt.Println(err)
}

func TestSendActionCard(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", platform.MsgTypeActionCard, ding_talk.ActionCardMessage{
		Title:          "测试",
		Text:           "测试消息",
		BtnOrientation: ding_talk.Vertical,
		Btns: []ding_talk.ActionCardBtn{
			{
				Title:     "确认",
				ActionURL: "https://www.dingtalk.com/",
			},
			{
				Title:     "取消",
				ActionURL: "https://www.dingtalk.com/",
			},
		},
	}, ding_talk.At{})
	fmt.Println(err)
}
