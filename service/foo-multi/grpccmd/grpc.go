package main

import (
	"base-frame/service/foo-multi/internal/rpc_server"
	"base-frame/service/foo-multi/internal/user"
	"github.com/gowins/dionysus"
	"github.com/gowins/dionysus/cmd"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"log"
	"net"
)

type grpcCommand struct {
	cmd    *cobra.Command
	server *grpc.Server
	addr   string
}

func NewGrpcCommand() *grpcCommand {
	return &grpcCommand{
		cmd:  &cobra.Command{Use: "grpc", Short: "Run as grpc server"},
		addr: ":50051",
	}
}

func (g *grpcCommand) GetCmd() *cobra.Command {
	g.cmd.RunE = func(cmd *cobra.Command, args []string) error {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Printf("listen error %v\n", err)
			return err
		}
		return g.server.Serve(lis)
	}
	return g.cmd
}

func (g *grpcCommand) RegShutdownFunc(stopSteps ...cmd.StopStep) {
	return
}

func (g *grpcCommand) GetShutdownFunc() cmd.StopFunc {
	return func() {
		g.server.GracefulStop()
	}
}

func main() {
	grpcServer := grpc.NewServer()
	user.RegisterUserServer(grpcServer, &rpc_server.UserServer{})
	grpcCmd := NewGrpcCommand()
	grpcCmd.server = grpcServer
	d := dionysus.NewDio()
	if err := d.DioStart("grpcCmd", grpcCmd); err != nil {
		log.Fatalf("dio start error %v", err)
	}
}
