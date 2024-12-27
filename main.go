package main

import (
	"chatapp/messages"
	"chatapp/network"
	"encoding/json"
	"log"
	"os"
)

func main() {
	localAddr, _ := network.GetLocalAddr()
	network.LocalNode = network.Node{
		ID:     network.LocalID,
		Addr:   localAddr,
		Status: "offline",
	}

	network.LocalTimeVector.Vector = make(map[string]int)
	network.LocalTimeVector.LogicalCLock = 0

	nodesJson, _ := os.ReadFile("nodes.json")

	network.LocalRouteTable.BuildRoutingTable(nodesJson)

	var words messages.Words

	rb, _ := os.ReadFile("words.json")

	err := json.Unmarshal([]byte(rb), &words)
	if err != nil {
		log.Fatalf("Error in Unmarshalling words list: %s", err)
	}

	// //fmt.Printf("%v", words.Words)

	// if err != nil {
	// 	log.Fatal("Error getting local addresses:", err)
	// 	os.Exit(1)
	// }

	//go network.ServerMulticastUDP(localAddr)
	go network.Server()
	go messages.Server()

	for id := range network.LocalRouteTable.Nodes {
		go network.LocalNode.NetworkClient(network.LocalRouteTable.Nodes[id])
	}

	go messages.LocalMessages.MesageMulticastEventGenerator(words)
	go messages.LocalMessages.ShowResults()
	select {}
}
