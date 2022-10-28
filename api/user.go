package main

import (
	"fmt"

	"base-frame/api/internal/config"
	"base-frame/api/internal/handler"

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
