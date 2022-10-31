package main

import (
	"base-frame/service/foo-multi/internal/rpc_server"
	user2 "base-frame/service/foo-multi/internal/user"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen address %s failure, error: %v", addr, err)
	}
	srv := grpc.NewServer()
	user2.RegisterUserServer(srv, &rpc_server.UserServer{})
	go func() {
		log.Printf("listening on: %s", lis.Addr())
		if err := srv.Serve(lis); err != nil {
			log.Fatalf("grpc server serve failure: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// graceful stop
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	srv.GracefulStop()
}
