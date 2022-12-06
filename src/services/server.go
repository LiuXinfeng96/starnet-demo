package services

import (
	"starnet-demo/src/configs"
	"starnet-demo/src/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"

	sdk "chainmaker.org/chainmaker/sdk-go/v2"
)

type Server struct {
	ginEngine        *gin.Engine
	config           *configs.Config
	log              *zap.Logger
	sulog            *zap.SugaredLogger
	gormDb           *gorm.DB
	sdkPool          *SdkPool
	satelliteChainId string
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
	server.satelliteChainId = "testchain"
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

func (s *Server) InitChainClient() {

	masterName := s.config.BCConfig[0].UserName
	user := new(db.User)
	err := s.QueryObjectByCondition(user, "user_name", masterName)
	if err != nil {
		materUser := &db.User{
			UserName:     masterName,
			UserRole:     db.CONTROL,
			UserPwd:      s.config.BCConfig[0].UserPwd,
			UserNickName: "主链用户",
			UserPhoneNum: "",
			UserEmail:    "",
		}
		err := s.InsertOneObjertToDB(materUser)
		if err != nil {
			panic(err)
		}
	}

	execName := s.config.BCConfig[1].UserName
	user = new(db.User)
	err = s.QueryObjectByCondition(user, "user_name", execName)
	if err != nil {
		execUser := &db.User{
			UserName:     execName,
			UserRole:     db.EXEC,
			UserPwd:      s.config.BCConfig[1].UserPwd,
			UserNickName: "星座链用户",
			UserPhoneNum: "",
			UserEmail:    "",
		}
		err := s.InsertOneObjertToDB(execUser)
		if err != nil {
			panic(err)
		}
	}

	masterClient, err := sdk.NewChainClient(
		sdk.WithConfPath(s.config.BCConfig[0].SdkConfigPath))
	if err != nil {
		panic(err)
	}

	s.sdkPool.Add(masterName, masterClient)

	execClient, err := sdk.NewChainClient(
		sdk.WithConfPath(s.config.BCConfig[1].SdkConfigPath))
	if err != nil {
		panic(err)
	}
	s.sdkPool.Add(execName, execClient)

}
