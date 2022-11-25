package main

import (
	"starnet-demo/src/configs"
	"starnet-demo/src/db"
	"starnet-demo/src/loggers"
	"starnet-demo/src/services"
)

func main() {
	var err error

	conf, err := configs.InitConfig(configs.GetConfigEnv())
	if err != nil {
		panic(err)
	}

	sugaredLogger, _ := loggers.InitLogger(conf.LogConfig)

	_, err = db.GormInit(conf.DBConfig, db.SchemaSlice, sugaredLogger)
	if err != nil {
		panic(err)
	}

	services.NewServer(w)
}
