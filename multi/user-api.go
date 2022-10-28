package main

import (
	"fmt"

	base "base-frame"
	"base-frame/multi/internal/config"
	"base-frame/multi/internal/handler"

	"github.com/gowins/dionysus/log"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile string

func main() {
	log.Setup(log.SetProjectName("Test"))
	log.Debug("haha")
	ginCmd := cmd.NewGinCommand()
	ginCmd.Flags().StringVarP(&configFile, "config file", "c", "", "")

	var c config.Config
	conf.MustLoad(configFile, &c)

	handler.RegisterHandlers(ginCmd)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	base.Start("user", ginCmd)
}
