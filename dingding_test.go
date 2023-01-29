package dingding

import (
	"fmt"
	"testing"

	"github.com/championlong/dingtalk-sdk/config"
	"github.com/championlong/dingtalk-sdk/model"
)

func initDingConfig() {
	query := make(map[string]config.DingdingQueryConfig)
	query["report"] = config.DingdingQueryConfig{
		Encrypt:     "",
		AccessToken: "",
	}
	Init(config.DingdingConfig{
		DingdingQuery: query,
	})
}

func TestSendText(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", MsgTypeText, model.TextMessage{
		Content: "text",
	}, model.At{})
	fmt.Println(err)
}

func TestSendMarkdown(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", MsgTypeMarkdown, model.MarkdownMessage{
		Title: "测试",
		Text:  "# 测试",
	}, model.At{})
	fmt.Println(err)
}

func TestSendLink(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", MsgTypeLink, model.LinkMessage{
		Title:      "测试",
		Text:       "测试文本",
		PicURL:     "",
		MessageURL: "https://open.dingtalk.com/document/group/custom-robot-access",
	}, model.At{})
	fmt.Println(err)
}

func TestSendFeedCard(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", MsgTypeFeedCard, model.FeedCardMessage{
		Links: []model.Links{
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
	}, model.At{})
	fmt.Println(err)
}

func TestSendActionCard(t *testing.T) {
	initDingConfig()
	err := SendDingMessage("report", MsgTypeActionCard, model.ActionCardMessage{
		Title:          "测试",
		Text:           "测试消息",
		BtnOrientation: model.Vertical,
		Btns: []model.ActionCardBtn{
			{
				Title:     "确认",
				ActionURL: "https://www.dingtalk.com/",
			},
			{
				Title:     "取消",
				ActionURL: "https://www.dingtalk.com/",
			},
		},
	}, model.At{})
	fmt.Println(err)
}
