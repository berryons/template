package config

import (
	"github.com/berryons/log"
	"github.com/caarlos0/env/v6"
)

var gEnv *Config

func Load() *Config {
	if gEnv == nil {
		load()
	}
	return gEnv
}

func load() {
	// TODO: godotenv 로 파일로 읽기 추가하기.
	gEnv = new(Config)

	parseErr := env.Parse(gEnv)
	if parseErr != nil {
		log.Println("Env contents parsing error: ", parseErr.Error())
	}

	if len(gEnv.ServerName) == 0 {
		gEnv.ServerName = "template"
	}

	if len(gEnv.ServerNetwork) == 0 {
		gEnv.ServerNetwork = "tcp"
	}

	if len(gEnv.ServerAddress) == 0 {
		gEnv.ServerAddress = "127.0.0.1"
	}

	if len(gEnv.ServerPort) == 0 {
		gEnv.ServerPort = "8080"
	}

	if len(gEnv.LogDir) == 0 {
		gEnv.LogDir = "logs"
	}

	if len(gEnv.LogLevel) == 0 {
		gEnv.LogLevel = "DEBUG"
	}

	if len(gEnv.LogFileLevel) == 0 {
		gEnv.LogFileLevel = "DEBUG"
	}

	if len(gEnv.Seed) == 0 {
		gEnv.Seed = "0d6e3a73fc47dfbeb2be572a03debef166ab1c3cb13ee4b125c557fc28f0dcad25385cbf73409e125dd2b45b5ac8c83e"
	}

	if len(gEnv.Salt) == 0 {
		gEnv.Salt = "9p67kZnk2vGrF6YFpnzY07XbTU/66ab1c3cb13ee4b125c557fc28f0dcad"
	}
}

type Config struct {
	ServerName    string `env:"server.name" required:"true"`
	ServerNetwork string `env:"server.network" required:"true"`
	ServerAddress string `env:"server.addr" required:"true"`
	ServerPort    string `env:"server.port" required:"true"`

	LogDir       string `env:"log.dir" required:"true"`
	LogLevel     string `env:"log.level" required:"true"`
	LogFileLevel string `env:"log.file.level" required:"true"`

	Seed string `env:"app.seed" required:"true"`
	Salt string `env:"app.salt" required:"true"`
}
