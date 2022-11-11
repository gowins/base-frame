package main

import (
	"base-frame/service/foo-multi/internal/rpc_server"
	"base-frame/service/foo-multi/internal/user"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gowins/dionysus"
	"github.com/gowins/dionysus/cmd"
	"github.com/gowins/dionysus/grpc/server"
	"github.com/gowins/dionysus/grpc/serverinterceptors"
	"os"
	"strconv"
	"time"
)

var defaultTimeout = 10
var defaultMaxConcurrentRequests = 1000

const (
	TimeoutEnv               = "GRPC_TIMEOUT"
	MaxConcurrentRequestsEnv = "GRPC_MAX_REQUEST"
)

func main() {
	dio := dionysus.NewDio()
	cfg := server.DefaultCfg
	cfg.Address = ":8081"
	c := cmd.NewGrpcCmd(cmd.WithCfg(cfg))
	c.EnableDebug()

	timeout := os.Getenv(TimeoutEnv)
	if timeout != "" {
		if timeoutInt, err := strconv.Atoi(timeout); err != nil && timeoutInt > 0 {
			defaultTimeout = timeoutInt
		}
	}

	maxRequest := os.Getenv(MaxConcurrentRequestsEnv)
	if maxRequest != "" {
		if maxRequestInt, err := strconv.Atoi(maxRequest); err != nil && maxRequestInt > 0 {
			defaultMaxConcurrentRequests = maxRequestInt
		}
	}

	c.AddUnaryServerInterceptors(serverinterceptors.HystrixRateLimitUnary(
		hystrix.CommandConfig{MaxConcurrentRequests: defaultMaxConcurrentRequests},
	))
	c.AddUnaryServerInterceptors(serverinterceptors.TimeoutUnary(time.Duration(defaultTimeout) * time.Second))
	// recover interceptor
	c.AddUnaryServerInterceptors(serverinterceptors.RecoveryUnary(serverinterceptors.DefaultRecovery()))
	c.AddStreamServerInterceptors(serverinterceptors.RecoveryStream(serverinterceptors.DefaultRecovery()))
	// tacing interceptor
	c.AddUnaryServerInterceptors(serverinterceptors.OpenTracingUnary())
	c.AddStreamServerInterceptors(serverinterceptors.OpenTracingStream())
	// register grpc service
	c.RegisterGrpcService(user.RegisterUserServer, &rpc_server.UserServer{})
	dio.DioStart("grpc", c)
}
