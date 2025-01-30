package config

type Config struct {
	DB        DBConfig     `json:"db"`
	Server    ServerConfig `json:"server"`
	Redis     RedisConfig  `json:"redis"`
	Constants Constants    `json:"constants"`
	Minio     MinioConfig  `json:"minio"`
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

type MinioConfig struct {
	Endpoint        string `json:"endpoint"`
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	SSL             bool   `json:"ssl"`
}

type Constants struct {
	Channel     string   `json:"channel"`
	Admins      []string `json:"admins"`
	SuperAdmins []string `json:"super_admins"`
}
