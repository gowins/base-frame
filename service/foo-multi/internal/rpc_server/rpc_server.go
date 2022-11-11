package rpc_server

import (
	user2 "base-frame/service/foo-multi/internal/user"
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"log"
	"strconv"
	"time"
)

type UserServer struct {
	user2.UnimplementedUserServer
}

func (u *UserServer) SayHello(ctx context.Context, req *user2.HelloRequest) (*user2.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok && len(md.Get("timesleep")) == 1 {
		sleepInt, _ := strconv.Atoi(md.Get("timesleep")[0])
		fmt.Printf("time sleep %v second\n", sleepInt)
		time.Sleep(time.Duration(sleepInt) * time.Second)
	}
	log.Printf("Received: %v", req.GetName())
	return &user2.HelloReply{Message: "Baseframe" + req.GetName()}, nil
}
