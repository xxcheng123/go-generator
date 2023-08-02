package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(AppConfig)

type ViperConfig struct {
	ConfigFile string
	ConfigName string
	ConfigType string
	ConfigPath string
}

func Init(viperConfig *ViperConfig) (err error) {
	//viper.SetConfigFile("config.yaml")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath(".")
	//如果配置了 `SetConfigFile`，`viper` 不会检查任何路径
	if viperConfig.ConfigFile != "" {
		viper.SetConfigFile(viperConfig.ConfigFile)
	} else {
		viper.AddConfigPath(viperConfig.ConfigPath)
		viper.AddConfigPath(viperConfig.ConfigName)
		if viperConfig.ConfigType != "" {
			viper.AddConfigPath(viperConfig.ConfigType)
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		//panic(err)
		return err
	}
	if err = viper.Unmarshal(Config); err != nil {
		return err
	}
	//监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("%s changed.\n", in.Name)
		if err = viper.Unmarshal(Config); err != nil {
			fmt.Println("config Unmarshal failed.", err)
		}
	})
	return err
}
