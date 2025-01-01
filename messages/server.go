package messages

import (
	msg "GoMulticastChatApp/messages/proto"
	"GoMulticastChatApp/network"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
)

func Server() {
	flag.Parse()
	lis, err := net.Listen("tcp", network.LocalNode.Addr+Port)
	if err != nil {
		log.Fatalf("failed to listen from server: %v", err)
	}
	//start grpc server
	server := grpc.NewServer()
	//register server
	msg.RegisterMessageServer(server, &MessageServer{})
	log.Printf("Message server listening at %v", lis.Addr())

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
