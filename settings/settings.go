package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Config = new(AppConfig)

func Init() (err error) {
	viper.SetConfigFile("config.yaml")
	//viper.SetConfigName("config")
	//viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
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
