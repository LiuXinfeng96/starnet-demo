package main

import (
	"starnet-demo/src/configs"
	"starnet-demo/src/db"
	"starnet-demo/src/handlers"
	"starnet-demo/src/loggers"
	"starnet-demo/src/routers"
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

	err = Start(server)
	if err != nil {
		panic(err)
	}

}

func Start(s *services.Server) error {
	//loading middleware
	s.GetGinEngine().Use(handlers.Cors())
	s.GetGinEngine().Use(loggers.GinLogger(s.GetLogger()))
	//s.GetGinEngine().Use(loggers.GinRecovery(s.GetLogger(), true))

	//loading route
	routers.LoadUserRouter(s)

	routers.LoadControlRouter(s)

	routers.LoadExecRouter(s)

	err := s.GetGinEngine().Run(":" + s.GetSeverPort())
	if err != nil {
		return err
	}

	return nil
}
