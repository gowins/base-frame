package config

import (
	"base-frame/common/env"
	"fmt"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gowins/dionysus/log"
	"github.com/spf13/viper"
)

var (
	FilePath   = "./etc/"
	ViperMap   sync.Map
	ConfigName = ""
)

func Setup() {
	if ConfigName == "" {
		viper.SetConfigName(env.GetRunEnv())
	} else {
		viper.SetConfigName(ConfigName)
	}
	viper.SetConfigType("yaml")
	viper.AddConfigPath(FilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("config has been reloaded at ", time.Now().Format("2006-01-02 15:04:05.000"))
	})
	viper.WatchConfig()
}

func SetupToml() {
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")
	viper.AddConfigPath(FilePath)
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("------------------------------------------------------------")
		fmt.Println("Config file changed:", e.Name)
		fmt.Println("config has been reloaded at ", time.Now().Format("2006-01-02 15:04:05.000"))
	})
	viper.WatchConfig()
}

// viper 包直接使用都是导出函数，这里直接返回 viper 新的实例
func NewViperExtra(log log.Logger, opt ...viper.Option) *viper.Viper {
	v := viper.NewWithOptions(opt...)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %w", err)
	}
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("viper config file changed:", e.Name)
		log.Infof("config has been reloaded at ", time.Now().Format("2006-01-02 15:04:05.000"))
	})
	viper.WatchConfig()
	return v
}

func NewViperWithStore(key string, logger log.Logger, opt ...viper.Option) (v *viper.Viper, err error) {
	v = NewViperExtra(logger, opt...)
	if _, exist := ViperMap.Load(key); exist {
		err = fmt.Errorf("viper key %s already exist in map", key)
		return
	}
	ViperMap.Store(key, v)
	return
}
