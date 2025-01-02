package network

import (
	"GoMulticastChatApp/config"
	ntw "GoMulticastChatApp/network/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func (n *Node) NetworkClient(peer Node) {
	conn, err := grpc.NewClient(peer.Addr+Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to server %s", err)
	}
	c := ntw.NewNetworkClient(conn)
	defer conn.Close()

	// Prevent exiting due to error of client not reaching the peer
	maxConnAttempts := 3
	for i := 0; i <= maxConnAttempts; i++ {
		// setup connection with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		jsonNode, _ := json.Marshal(n)
		peerResp, nodeRegisterError := c.NodeInit(ctx, &ntw.NodeInitRequest{Message: string(jsonNode)})
		if nodeRegisterError == nil {
			log.Printf(config.Green+peer.Addr+Port+" %s"+config.Reset, peerResp)
			peer.Status = "online"
			LocalRouteTable.Nodes[peer.ID] = peer
			return
		} else {
			log.Printf(config.Yellow+"Peer not ready %s Remaining connection attempts:%d"+config.Reset, peer.Addr+Port, maxConnAttempts-i)
		}
		time.Sleep(time.Second * 5)
	}

	// after trying 3 times giving up on updating status
	log.Printf(config.Red+"node:%s not responding give up status."+config.Reset, peer.Addr+Port)
}
