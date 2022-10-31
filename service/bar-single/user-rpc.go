package main

import (
	"base-frame/service/bar-single/user"
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type userServer struct {
	user.UnimplementedUserServer
}

func (u *userServer) SayHello(ctx context.Context, req *user.HelloRequest) (*user.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &user.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("listen address %s failure, error: %v", addr, err)
	}
	srv := grpc.NewServer()
	user.RegisterUserServer(srv, &userServer{})
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
