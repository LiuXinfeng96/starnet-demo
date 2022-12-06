package configs

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"starnet-demo/src/db"
	"starnet-demo/src/loggers"

	"github.com/spf13/viper"
)

type BlockChainConfig struct {
	UserName      string `mapstructure:"user_name"`
	UserPwd       string `mapstructure:"user_pwd"`
	SdkConfigPath string `mapstructure:"sdk_config_path"`
}

type Config struct {
	ServerPort string              `mapstructure:"server_port"`
	LogConfig  *loggers.LogConfig  `mapstructure:"log_config"`
	DBConfig   *db.DBConfig        `mapstructure:"db_config"`
	BCConfig   []*BlockChainConfig `mapstructure:"block_chain_config"`
	MaterC     *MasterContract     `mapstructure:"master_contract"`
	ExecC      *ExecContract       `mapstructure:"exec_contract"`
}

type Contract struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"name"`
}

type MasterContract struct {
	Debris        *Contract `mapstructure:"debris"`
	Instruction   *Contract `mapstructure:"instruction"`
	Orbit         *Contract `mapstructure:"orbit"`
	Constellation *Contract `mapstructure:"constellation"`
	Loginlog      *Contract `mapstructure:"loginlog"`
	Operation     *Contract `mapstructure:"operation"`
}

type ExecContract struct {
	Debris      *Contract `mapstructure:"debris"`
	Satellite   *Contract `mapstructure:"satellite"`
	Instruction *Contract `mapstructure:"instruction"`
	Exec        *Contract `mapstructure:"exec"`
	Control     *Contract `mapstructure:"control"`
	Fault       *Contract `mapstructure:"fault"`
	Commstate   *Contract `mapstructure:"commstate"`
	Netstate    *Contract `mapstructure:"netstate"`
	Operation   *Contract `mapstructure:"operation"`
}

const DEFAULT_SERVER_PORT = "8096"

// GetConfigEnv --Specify the path and name of the configuration file (Env)
func GetConfigEnv() string {
	var env string
	n := len(os.Args)
	for i := 1; i < n-1; i++ {
		if os.Args[i] == "-e" || os.Args[i] == "--env" {
			env = os.Args[i+1]
			break
		}
	}
	fmt.Println("[env]:", env)
	if env == "" {
		fmt.Println("env is empty, set default")
		env = ""
	}
	return env
}

// GetFlagPath --Specify the path and name of the configuration file (flag)
func GetFlagPath() string {
	var configPath string
	flag.StringVar(&configPath, "config", "../conf/system_config.yaml", "please input the system config file path")
	flag.Parse()
	return configPath
}

// InitConfig --Set config path and file name
func InitConfig(configPath string) (*Config, error) {
	var err error
	if len(configPath) == 0 {
		configPath = GetFlagPath()
	}

	viper.SetConfigFile(configPath)
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var conf Config
	err = viper.Unmarshal(&conf)
	if err != nil {
		return nil, err
	}

	if conf.LogConfig == nil {
		conf.LogConfig = new(loggers.LogConfig)
	}

	if conf.DBConfig == nil {
		return nil, errors.New("not found the db config")
	}

	if conf.BCConfig == nil || len(conf.BCConfig) < 2 {
		return nil, errors.New("not found the block chain config")
	}

	if len(conf.ServerPort) == 0 {
		conf.ServerPort = DEFAULT_SERVER_PORT
	}

	return &conf, nil
}
