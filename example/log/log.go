package main

import (
	logger "base-frame/common/logger"
	"context"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	logger.Setup()

	logx.Info("88888")

	logger.Infow("infow foo",
		logger.Field("url", "http://localhost:8080/hello"),
		logger.Field("attempt", 3),
		logger.Field("backoff", time.Second),
	)
	logger.Errorw("errorw foo",
		logger.Field("url", "http://localhost:8080/hello"),
		logger.Field("attempt", 3),
		logger.Field("backoff", time.Second),
	)
	logger.Sloww("sloww foo",
		logger.Field("url", "http://localhost:8080/hello"),
		logger.Field("attempt", 3),
		logger.Field("backoff", time.Second),
	)
	logger.Error("error")
	logger.Infov(map[string]interface{}{
		"url":     "localhost:8080/hello",
		"attempt": 3,
		"backoff": time.Second,
		"value":   "foo",
	})

	logger.Error("error")
	logger.Errorf("error %s", "foo")
	logger.Errorv(map[string]interface{}{})
	logger.Errorw("errorw foo")
	logger.Info("info")
	log := logger.WithContext(context.Background())
	log.Info()

	logger.WithDuration(1100*time.Microsecond).Infow("infow withduration",
		logger.Field("url", "localhost:8080/hello"),
		logger.Field("attempt", 3),
		logger.Field("backoff", time.Second),
	)
	logger.WithContext(context.Background()).WithDuration(1100*time.Microsecond).Errorw(
		"errorw withcontext withduration",
		logger.Field("url", "localhost:8080/hello"),
		logger.Field("attempt", 3),
		logger.Field("backoff", time.Second),
	)
	logger.WithDuration(1100*time.Microsecond).WithContext(context.Background()).Errorw(
		"errorw withduration withcontext",
		logger.Field("url", "localhost:8080/hello"),
		logger.Field("attempt", 3),
		logger.Field("backoff", time.Second),
	)

}
