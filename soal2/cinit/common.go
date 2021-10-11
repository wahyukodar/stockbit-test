package cinit

import (
	"log"
	"stockbit_test/soal2/pkg/pprof"

	"github.com/jinzhu/configor"
)

const (
	MySQL = "MySQL"
)

var Config = struct {
	Service struct {
		Name      string `default:""`
		Version   string `default:"v1.0"`
	}
	Mysql struct {
		DbName   string `default:"omdb"`
		Addr     string `default:"127.0.0.1"`
		User     string `default:"root"`
		Password string `default:"root"`
		Port     int    `default:"3306"`
		IDleConn int    `default:"4"`
		MaxConn  int    `default:"20"`
	}
	RestServiceMovie struct {
		Port    string `default:":8080"`
		Address string `default:"127.0.0.1:8080"`
	}
	GrpcServiceMovie struct {
		Port    string `default:":4040"`
		Address string `default:"127.0.0.1:4040"`
	}
}{}

func configInit(sn string) {
	err := configor.New(&configor.Config{ErrorOnUnmatchedKeys: true}).Load(&Config, "config.yml")
	if err != nil {
		log.Printf("load config error:%+v", err)
	}
	if Config.Service.Name == "" {
		Config.Service.Name = sn
	}
	log.Printf("config: %+v\n", Config)
}

var closeArgs []string

func InitOption(sn string, args ...string) {
	go pprof.Run()
	closeArgs = args
	configInit(sn)

	for _, o := range args {
		switch o {
		case MySQL:
			mysqlInit()
			// others init such Redis, Mongo, etc..
		}
	}
}

func Close() {
	for _, o := range closeArgs {
		switch o {
		case MySQL:
			mysqlClose()
		}
	}
}
