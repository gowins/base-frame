package rpc_client

import (
	user2 "base-frame/service/foo-multi/internal/user"
	"context"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/gowins/dionysus/grpc/client"
	"github.com/gowins/dionysus/grpc/clientinterceptors"
	"google.golang.org/grpc/metadata"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	TimeoutEnv               = "GRPC_TIMEOUT"
	MaxConcurrentRequestsEnv = "GRPC_MAX_REQUEST"
)

var defaultMaxConcurrentRequests = 1000
var defaultTimeout = 10

func GetRPCHello() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user2.NewUserClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &user2.HelloRequest{Name: "nameing"})
	if err != nil {
		log.Printf("could not sayhello: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func GetRPCHelloV2() {
	maxRequest := os.Getenv(MaxConcurrentRequestsEnv)
	if maxRequest != "" {
		if maxRequestInt, err := strconv.Atoi(maxRequest); err != nil && maxRequestInt > 0 {
			defaultMaxConcurrentRequests = maxRequestInt
		}
	}
	timeout := os.Getenv(TimeoutEnv)
	if timeout != "" {
		if timeoutInt, err := strconv.Atoi(timeout); err != nil && timeoutInt > 0 {
			defaultTimeout = timeoutInt
		}
	}

	breaker := grpc.WithChainUnaryInterceptor(clientinterceptors.BreakerUnary(
		hystrix.CommandConfig{MaxConcurrentRequests: defaultMaxConcurrentRequests},
	))
	conn, err := client.Get("foo-multi-grpc:8081",
		breaker,
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := user2.NewUserClient(conn)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// Contact the server and print out its response.
			md := metadata.Pairs("timesleep", "2")
			ctx := metadata.NewOutgoingContext(context.Background(), md)
			r, err := c.SayHello(ctx, &user2.HelloRequest{Name: "nameing"})
			if err != nil {
				log.Printf("could not greet: %v", err)
				return
			}
			log.Printf("Greeting: %s", r.GetMessage())
		}()
	}
	wg.Wait()
}
