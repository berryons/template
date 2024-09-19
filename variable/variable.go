package variable

import (
	"github.com/caarlos0/env/v6"
	"log"
	"os"
)

var gEnv *EnvVariables

func New() *EnvVariables {
	if gEnv == nil {
		gEnv = new(EnvVariables)
		gEnv.init(gEnv)
	}
	return gEnv
}

type EnvVariables struct {
	ServerName    string `env:"server.name" required:"true"`
	ServerNetwork string `env:"server.network" required:"true"`
	ServerAddress string `env:"server.addr" required:"true"`
	ServerPort    string `env:"server.port" required:"true"`

	LogDir          string `env:"log.dir" required:"true"`
	LogLevel        string `env:"log.level" required:"true"`
	LogEnableStdout bool   `env:"log.enable.stdout" required:"true"`

	Seed string `env:"app.seed" required:"true"`
	Salt string `env:"app.salt" required:"true"`
}

func (pSelf *EnvVariables) init(envVar *EnvVariables) {
	parseErr := env.Parse(envVar)
	if parseErr != nil {
		log.Println("Env contents parsing error: ", parseErr.Error())
	}

	if len(pSelf.ServerName) == 0 {
		pSelf.ServerName = "template"
	}

	if len(pSelf.ServerNetwork) == 0 {
		pSelf.ServerNetwork = "tcp"
	}

	if len(pSelf.ServerAddress) == 0 {
		pSelf.ServerAddress = "127.0.0.1"
	}

	if len(pSelf.ServerPort) == 0 {
		pSelf.ServerPort = "8080"
	}

	if len(pSelf.LogDir) == 0 {
		pSelf.LogDir = "logs"
	}

	if len(pSelf.LogLevel) == 0 {
		pSelf.LogLevel = "DEBUG"
	}

	if len(os.Getenv("log.enable.stdout")) == 0 {
		pSelf.LogEnableStdout = true
	}

	if len(pSelf.Seed) == 0 {
		pSelf.Seed = "0d6e3a73fc47dfbeb2be572a03debef166ab1c3cb13ee4b125c557fc28f0dcad25385cbf73409e125dd2b45b5ac8c83e"
	}

	if len(pSelf.Salt) == 0 {
		pSelf.Salt = "9p67kZnk2vGrF6YFpnzY07XbTU/66ab1c3cb13ee4b125c557fc28f0dcad"
	}
}
