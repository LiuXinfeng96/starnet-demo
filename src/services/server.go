package services

import (
	"starnet-demo/src/configs"

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
	sdkPool   *SdkPool
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

func WithSdkPool(maxEntries int) Option {
	return func(s *Server) {
		s.sdkPool = NewSdkPool(maxEntries)
	}
}

func NewServer(opts ...Option) (*Server, error) {
	server := new(Server)
	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}

func (s *Server) GetGinEngine() *gin.Engine {
	return s.ginEngine
}

func (s *Server) GetSeverPort() string {
	return s.config.ServerPort
}

func (s *Server) GetLogger() *zap.Logger {
	return s.log
}

func (s *Server) GetSuLogger() *zap.SugaredLogger {
	return s.sulog
}
