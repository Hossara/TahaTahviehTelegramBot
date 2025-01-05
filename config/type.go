package config

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
}

type DBConfig struct {
	Host   string `json:"host"`
	Port   uint   `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type ServerConfig struct {
	Token           string `json:"token"`
	Proxy           string `json:"proxy"`
	Port            uint   `json:"port"`
	NewUpdateOffset uint   `json:"new_update_offset"`
	BotTimeout      uint   `json:"bot_timeout"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type Constants struct {
	Admins      []string `json:"admins"`
	SuperAdmins []string `json:"super_admins"`
}
