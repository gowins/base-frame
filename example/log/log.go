package main

import (
	"github.com/gowins/dionysus/log"
)

func main() {
	opts := []log.Option{ // 根据实际需求添加option
		log.WithLevelEnabler(log.DebugLevel),
		log.WithEncoderCfg(log.NewEncoderConfig()),
		log.AddCallerSkip(1),
		log.AddCaller(),
	}
	//todo: opts隐藏与否
	logger, _ := log.New(log.ZapLogger, opts...)
	logger.WithField("url", "http://localhost:8080/hello").Infof("infow foo")
	logger.WithField("url", "http://localhost:8080/hello").Errorf("errorw foo")

	logger.Error("error")
	logger.Errorf("error %s", "foo")
}
