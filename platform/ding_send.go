package platform

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/championlong/dingtalk-sdk/model/ding_talk"
	"net/url"
	"strconv"
	"time"

	"github.com/championlong/dingtalk-sdk/config"
	"github.com/championlong/dingtalk-sdk/utils"
)

var dingClient = utils.NewHttpClient()

type MsgType string

const (
	MsgTypeText       MsgType = "text"
	MsgTypeMarkdown   MsgType = "markdown"
	MsgTypeLink       MsgType = "link"
	MsgTypeActionCard MsgType = "actionCard"
	MsgTypeFeedCard   MsgType = "feedCard"
)

type DingMasterJob struct {
	Url        string                       `json:"-"`       //请求url
	KindRobot  string                       `json:"-"`       //机器人种类
	Msgtype    MsgType                      `json:"msgtype"` //发送类型
	Text       *ding_talk.TextMessage       `json:"text,omitempty"`
	Markdown   *ding_talk.MarkdownMessage   `json:"markdown,omitempty"`
	Link       *ding_talk.LinkMessage       `json:"link,omitempty"`
	ActionCard *ding_talk.ActionCardMessage `json:"actionCard,omitempty"`
	FeedCard   *ding_talk.FeedCardMessage   `json:"feedCard,omitempty"`
	At         ding_talk.At                 `json:"at,omitempty"`
	Query      config.DingRobotsConfig      `json:"-"`
}

type SendMessage interface {
	SendMessage(message interface{}) error
}

// SendMessage 发送消息
func (job *DingMasterJob) SendMessage(message interface{}) error {
	switch job.Msgtype {
	case MsgTypeText:
		if value, ok := message.(ding_talk.TextMessage); ok {
			job.Text = &value
		}
	case MsgTypeMarkdown:
		if value, ok := message.(ding_talk.MarkdownMessage); ok {
			job.Markdown = &value
		}
	case MsgTypeLink:
		if value, ok := message.(ding_talk.LinkMessage); ok {
			job.Link = &value
		}
	case MsgTypeActionCard:
		if value, ok := message.(ding_talk.ActionCardMessage); ok {
			job.ActionCard = &value
		}
	case MsgTypeFeedCard:
		if value, ok := message.(ding_talk.FeedCardMessage); ok {
			job.FeedCard = &value
		}
	}
	return job.PostDingWebHookMsg()
}

// PostDingWebHookMsg 发送请求
func (job *DingMasterJob) PostDingWebHookMsg() error {
	ctx := fmt.Sprintf("[%s]", job.KindRobot)
	queries := make(map[string]string)
	queries["access_token"] = job.Query.AccessToken
	if job.Query.Encrypt != "" {
		timestampInt := time.Now().UnixNano() / 1e6
		timestamp := strconv.FormatInt(timestampInt, 10)
		secret := job.Query.Encrypt
		signString := timestamp + "\n" + secret
		h := hmac.New(sha256.New, []byte(secret))
		h.Write([]byte(signString))
		sha := h.Sum(nil)
		sign := base64.StdEncoding.EncodeToString(sha)

		queries["sign"] = url.QueryEscape(sign)
		queries["timestamp"] = timestamp
	}

	var resBody, err = dingClient.PostJson(utils.PackUrl(job.Url, queries), job)
	if err != nil {
		return fmt.Errorf("error while %s. err: %s", ctx, err.Error())
	}

	var data = new(ding_talk.DingDingResponse)
	err = json.Unmarshal(resBody, data)
	if err != nil {
		return fmt.Errorf("error while unmarshal %s. err: %s", ctx, err.Error())
	}
	if data.ErrCode != 0 {
		return fmt.Errorf("failed to %s. response: %s", ctx, resBody)
	}
	return nil
}
