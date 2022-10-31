package main

import (
	"base-frame/service/bar-single/handler"
	"fmt"

	"github.com/gowins/dionysus"
	"github.com/gowins/dionysus/cmd"
)

func main() {
	gcmd := cmd.NewGinCommand()
	handler.RegisterHandler(gcmd)
	d := dionysus.NewDio()
	if err := d.DioStart("ginxdemo", gcmd); err != nil {
		fmt.Printf("dio start error %v\n", err)
	}
}
