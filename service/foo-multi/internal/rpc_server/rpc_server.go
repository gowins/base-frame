package rpc_server

import (
	user2 "base-frame/service/foo-multi/internal/user"
	"context"
	"log"
)

type UserServer struct {
	user2.UnimplementedUserServer
}

func (u *UserServer) SayHello(ctx context.Context, req *user2.HelloRequest) (*user2.HelloReply, error) {
	log.Printf("Received: %v", req.GetName())
	return &user2.HelloReply{Message: "Hello " + req.GetName()}, nil
}
