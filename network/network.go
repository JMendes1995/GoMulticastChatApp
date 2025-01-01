package network

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"sync"

	ntw "GoMulticastChatApp/network/proto"
)

// 1. Setup the routing table with the node list passed via input
// 2. Init fuction that will fload the network with init messages.
// 2.1 for each init message comming from the neighbours will change the status (STATUS AVALIABLE: offline, online)
// 2.2 when a node sends a message to its peer and gots no response, will change the connection status to offine.
// 2.3 when a node is offline will not be shared any messages. This will avoid floading the network with messages.
// 3. Each node must have a status connected to transmit messages

type Node struct {
	ID     string `json:"ID"`
	Addr   string `json:"Addr"`
	Status string `json:"Status"`
}

type RoutingTable struct {
	Nodes map[string]Node
	Mutex sync.Mutex
}

type NetworkServer struct {
	ntw.UnimplementedNetworkServer
}

const Port = ":50000"

var (
	LocalID         = os.Getenv("ID")
	LocalNode       Node
	LocalRouteTable RoutingTable
)

func (rt *RoutingTable) BuildRoutingTable(nJson []byte) {
	rt.Mutex.Lock()
	err := json.Unmarshal([]byte(nJson), &rt)
	if err != nil {
		log.Fatalf("Error in Unmarshalling Node: %s", err)
	}
	// check if the localnode is in the routing table if is remove it
	_, OK := rt.Nodes[LocalNode.ID]
	if OK {
		delete(rt.Nodes, LocalNode.ID)
	}
	rt.Mutex.Unlock()
	LocalNode.Status = "online"
}

func (n NetworkServer) NodeInit(ctx context.Context, in *ntw.NodeInitRequest) (*ntw.NodeInitResponse, error) {
	// convert the string comming from request to Node format
	reqNode := NodeJsonUnmarshal(in.GetMessage())

	LocalRouteTable.Mutex.Lock()
	// update the node status to online
	LocalRouteTable.Nodes[reqNode.ID] = reqNode
	LocalRouteTable.Mutex.Unlock()

	return &ntw.NodeInitResponse{Result: "Node status changed"}, nil

}

func NodeJsonUnmarshal(data string) Node {
	var n Node
	err := json.Unmarshal([]byte(data), &n)
	if err != nil {
		log.Fatalf("Error in Unmarshalling Node: %s", err)
	}
	return n
}
