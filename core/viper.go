package core

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"myGo/global"
	"os"
	"path/filepath"
)

const (
	ConfigEnv  = "ServerConfig"
	ConfigFile = "config.yaml"
)

func Viper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv(ConfigEnv); configEnv == "" {
				config = ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
			} else {
				config = configEnv
				fmt.Printf("您正在使用GVA_CONFIG环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.ServerConfig); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		fmt.Println(err)
	}
	// root 适配性
	// 根据root位置去找到对应迁移位置,保证root路径有效
	global.ServerConfig.AutoCode.Root, _ = filepath.Abs("..")
	return v
}
