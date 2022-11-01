package main

import (
	"fmt"
	"log"

	"base-frame/common/config"
	"base-frame/common/env"
	"base-frame/service/bar-single/internal/handler"
	"base-frame/service/bar-single/internal/healthy_checker"

	"github.com/gowins/dionysus/healthy"
	"github.com/gowins/dionysus/step"

	"github.com/gowins/dionysus"
	"github.com/gowins/dionysus/cmd"
)

func main() {
	gcmd := cmd.NewGinCommand()
	handler.RegisterHandler(gcmd)
	d := dionysus.NewDio()
	healthy.RegReadinessCheckers(healthy_checker.CheckUpstreamHealth, healthy_checker.CheckDBHealth)
	preSteps := []step.InstanceStep{
		{
			StepName: "PrePrint1", Func: func() error {
				fmt.Println("init orm")
				return nil
			},
		},
		// 运行环境
		{
			StepName: "env", Func: func() error {
				env.Setup()
				return nil
			},
		},
		// 配置
		{
			StepName: "config", Func: func() error {
				config.ConfigName = "config"
				config.Setup()
				return nil
			},
		},
	}
	err := d.PreRunStepsAppend(preSteps...)
	if err != nil {
		log.Printf("step append error %v", err)
		return
	}
	if err := d.DioStart("ginxdemo", gcmd); err != nil {
		fmt.Printf("dio start error %v\n", err)
	}
}
