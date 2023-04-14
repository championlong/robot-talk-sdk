package config

// PlatformConfig 平台机器人发送配置
type PlatformConfig struct {
	// DingRobots 机器人列表 key机器人名称
	DingRobots   map[string]DingRobotsConfig   `mapstructure:"ding_talk-platform" json:"ding_talk-platform" yaml:"ding_talk-platform"`
	FeishuRobots map[string]FeishuRobotsConfig `mapstructure:"feishu-platform" json:"feishu-platform" yaml:"feishu-platform"`
}

// DingRobotsConfig 钉钉机器人配置
type DingRobotsConfig struct {
	// KeysWord 关键字
	KeysWord string `mapstructure:"keys-word" json:"keys-word" yaml:"keys-word"`
	// Encrypt 加签信息
	Encrypt string `mapstructure:"encrypt" json:"encrypt" yaml:"encrypt"`
	// AccessToken 请求token信息
	AccessToken string `mapstructure:"access-token" json:"access-token" yaml:"access-token"`
}

// FeishuRobotsConfig 飞书机器人配置
type FeishuRobotsConfig struct {
}
