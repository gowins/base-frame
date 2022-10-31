package main

import (
	"base-frame/service/foo-multi/internal/handler"

	"github.com/gowins/dionysus"
	"github.com/gowins/dionysus/cmd"
	"github.com/gowins/dionysus/log"
)

func main() {
	log.Setup(log.SetProjectName("Test"))
	log.Debug("haha")
	ginCmd := cmd.NewGinCommand()
	dio := dionysus.NewDio()
	handler.RegisterHandlers(ginCmd)
	_ = dio.DioStart("user", ginCmd)
}
