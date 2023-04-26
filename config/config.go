package config

var PlatformConfigGlobal *PlatformConfig

// PlatformConfig 平台机器人发送配置
type PlatformConfig struct {
	// DingRobots 钉钉机器人列表
	DingRobots map[string]DingRobotsConfig `mapstructure:"ding-platform" json:"ding-platform" yaml:"ding-platform"`
	// FeishuRobots 飞书机器人列表
	FeishuRobots map[string]FeishuRobotsConfig `mapstructure:"feishu-platform" json:"feishu-platform" yaml:"feishu-platform"`
}

// DingRobotsConfig 钉钉机器人配置
type DingRobotsConfig struct {
	// Encrypt 加签信息
	Encrypt string `mapstructure:"encrypt" json:"encrypt" yaml:"encrypt"`
	// AccessToken 请求token信息
	AccessToken string `mapstructure:"access-token" json:"access-token" yaml:"access-token"`
}

// FeishuRobotsConfig 飞书机器人配置
type FeishuRobotsConfig struct {
}
