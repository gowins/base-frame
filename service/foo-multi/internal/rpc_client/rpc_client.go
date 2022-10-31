package rpc_client

import (
	user2 "base-frame/service/foo-multi/internal/user"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

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
