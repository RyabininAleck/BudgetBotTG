package config

type ServerConfig struct {
	Chat map[string]ChatConfig
	Bot  BotConfig
	DB   string
}

type ChatConfig struct {
	HelloMsg string
	HelpMsg  string
}

type BotConfig struct {
	Debug   bool
	Timeout int
}

func GetServerConfig(pathToConfig string) ServerConfig {
	ruChatCfg := ChatConfig{
		HelloMsg: "Добро пожаловать",
		HelpMsg:  "тут могут помочь 89113867588",
	}

	botCfg := BotConfig{
		Debug:   true,
		Timeout: 60,
	}
	cfg := ServerConfig{
		Chat: map[string]ChatConfig{"ru": ruChatCfg},
		Bot:  botCfg,
		DB:   "",
	}

	return cfg
}
