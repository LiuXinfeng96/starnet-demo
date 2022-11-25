package services

import (
	"starnet-demo/src/configs"
	"starnet-demo/src/loggers"
	"starnet-demo/src/routers"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Server struct {
	ginEngine *gin.Engine
	config    *configs.Config
	log       *zap.Logger
	sulog     *zap.SugaredLogger
	gormDb    *gorm.DB
}

type Option func(s *Server)

func WithGinEngin() Option {
	return func(s *Server) {
		g := gin.New()
		s.ginEngine = g
	}
}

func WithConfig(cfg *configs.Config) Option {
	return func(s *Server) {
		s.config = cfg
	}
}

func WithSuLog(log *zap.SugaredLogger) Option {
	return func(s *Server) {
		s.sulog = log
	}
}

func WithLog(log *zap.Logger) Option {
	return func(s *Server) {
		s.log = log
	}
}

func WithGormDb(db *gorm.DB) Option {
	return func(s *Server) {
		s.gormDb = db
	}
}

func NewServer(opts ...Option) error {
	server := new(Server)
	for _, opt := range opts {
		opt(server)
	}

	return nil
}

func (s *Server) Start() error {
	//loading middleware
	s.ginEngine.Use(loggers.GinLogger(s.log), loggers.GinRecovery(s.log, true))

	//loading route
	routers.LoadUserRouter(s.ginEngine)

	routers.LoadControlRouter(s.ginEngine)

	routers.LoadExecRouter(s.ginEngine)

	err := s.ginEngine.Run(s.config.ServerPort)
	if err != nil {
		return err
	}

	return nil
}
