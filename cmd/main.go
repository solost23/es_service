package main

import (
	"es_service/configs"
	"es_service/internal/server"
	"es_service/pkg/helper"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"os"
	"path"
)

var (
	WebConfigPath = "configs/conf.yml"
	WebLogPath    = "logs"
	version       = "__BUILD_VERSION_"
	execDir       string
	provider      string
	st, v, V      bool
)

func main() {
	flag.StringVar(&execDir, "d", ".", "项目目录")
	flag.StringVar(&provider, "p", "consul", "项目配置提供者")
	flag.BoolVar(&v, "v", false, "查看版本号")
	flag.BoolVar(&V, "V", false, "查看版本号")
	flag.BoolVar(&st, "s", false, "项目状态")
	flag.Parse()
	if v || V {
		fmt.Println(version)
		os.Exit(-1)
	}
	// run
	// 从consul读取配置
	serverConfig, err := InitConfigFromConsul()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	sl, err := helper.InitLogger(execDir, WebLogPath, serverConfig.Mode)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	server.NewServer(serverConfig, sl).Run()
}

func InitConfig() (serverConfig *configs.ServerConfig, err error) {
	serverConfig = new(configs.ServerConfig)
	confPath := path.Join(execDir, WebConfigPath)
	v := viper.New()
	v.SetConfigFile(confPath)
	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(serverConfig); err != nil {
		return nil, err
	}
	return serverConfig, nil
}

func InitConfigFromConsul() (serverConfig *configs.ServerConfig, err error) {
	serverConfig = new(configs.ServerConfig)
	v := viper.New()
	// 通过项目配置读取基本信息
	v.SetConfigFile(path.Join(execDir, WebConfigPath))
	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}
	if err = v.Unmarshal(serverConfig); err != nil {
		return nil, err
	}

	// 从配置中心读取配置
	err = v.AddRemoteProvider(provider,
		fmt.Sprintf("%s:%d", serverConfig.ConsulConfig.Host, serverConfig.ConsulConfig.Port),
		serverConfig.ConfigPath)
	if err != nil {
		return nil, err
	}

	v.SetConfigType("YAML")

	if err = v.ReadRemoteConfig(); err != nil {
		return nil, err
	}

	if err = v.Unmarshal(serverConfig); err != nil {
		return nil, err
	}
	return serverConfig, nil

}
