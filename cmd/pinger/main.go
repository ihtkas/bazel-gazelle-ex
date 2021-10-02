package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pingerpb "github.com/ihtkas/bazel-gazelle-ex/api/pinger"

	"github.com/ihtkas/bazel-gazelle-ex/pinger"
	grpc "google.golang.org/grpc"
)

var port = flag.Int("port", 10000, "The server port")

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pingerpb.RegisterPingerServiceServer(grpcServer, &pinger.Pinger{})
	grpcServer.Serve(lis)
}
