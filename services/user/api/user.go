package main

import (
	"fmt"

	base "base-frame"
	"base-frame/cmd"
	"base-frame/services/user/api/internal/config"
	"base-frame/services/user/api/internal/handler"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile string

func main() {

	ginCmd := cmd.NewGinCommand()
	ginCmd.Flags().StringVarP(&configFile, "config file", "c", "", "")

	var c config.Config
	conf.MustLoad(configFile, &c)

	handler.RegisterHandlers(ginCmd)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	base.Start("user", ginCmd)
}
