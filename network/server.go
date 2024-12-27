package network

import (
	ntw "chatapp/network/proto"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)


func Server() {
	flag.Parse()
	lis, err := net.Listen("tcp", LocalNode.Addr+Port)
	if err != nil {
		log.Fatalf("failed to listen from server: %v", err)
	}
	//start grpc server
	server := grpc.NewServer()
	//register server
	ntw.RegisterNetworkServer(server, &NetworkServer{})
	log.Printf("Network server P2P server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
