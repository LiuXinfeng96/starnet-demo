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

	sugaredLogger, logger := loggers.InitLogger(conf.LogConfig)

	db, err := db.GormInit(conf.DBConfig, db.TableSlice, sugaredLogger)
	if err != nil {
		panic(err)
	}

	server, err := services.NewServer(services.WithConfig(conf),
		services.WithGinEngin(),
		services.WithGormDb(db),
		services.WithLog(logger),
		services.WithSuLog(sugaredLogger),
	)
	if err != nil {
		panic(err)
	}

	err = server.Start()
	if err != nil {
		panic(err)
	}

}
