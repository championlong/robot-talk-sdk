# robot-talk-sdk
> 目标：对接市面上常见的通讯软件，让使用者选择自己想要的发送工具做到开箱即用
## 已实现
### 钉钉机器人🤖️
> 支持钉钉群组多机器人配置化发送
* 集成方式: go get github.com/championlong/robot-talk-sdk
* 支持文本发送类型: Text、Markdown、Link、FeedCard、ActionCard
#### 配置文件
以yaml为例，`alert-robot`为自定义的机器人名称，`encrypt`加签值，如果使用钉钉关键字传输的话需自己维护保证消息体存在关键字，
`access-token`为创建机器人中提供的请求token
```yaml
ding-platform:
  alert-robot:
    encrypt: xxxxxx
    access-token: xxxxx
```
加载配置文件到结构体后
```go
package main

import (
	"github.com/championlong/robot-talk-sdk/config"
	"github.com/championlong/robot-talk-sdk/model/ding_talk"
	"github.com/championlong/robot-talk-sdk/platform"
)

func main() {
	global := config.PlatformConfig{} // 要将上方配置加载到结构体
	Init(global) // 将加载的配置文件初始化到SDK
	// 发送消息
	err := SendDingMessage("alert-robot", platform.MsgTypeText, ding_talk.TextMessage{Content: "文本消息"}, ding_talk.At{})
}

```
## 待实现
### 飞书
### 企业微信
### 微信公众号
