package main

import (
	"GoMulticastChatApp/messages"
	"GoMulticastChatApp/network"
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
	messages.LocalLamportClock.Timestamp = 0
	messages.LocalMessageQueue.Messages = make(map[int][]messages.MessageData)

	nodesJson, _ := os.ReadFile("nodes.json")

	network.LocalRouteTable.BuildRoutingTable(nodesJson)

	var words messages.Words

	rb, _ := os.ReadFile("words.json")

	err := json.Unmarshal([]byte(rb), &words)
	if err != nil {
		log.Fatalf("Error in Unmarshalling words list: %s", err)
	}

	go network.Server()
	go messages.Server()

	for id := range network.LocalRouteTable.Nodes {
		go network.LocalNode.NetworkClient(network.LocalRouteTable.Nodes[id])
	}

	go messages.LocalMessageQueue.MesageMulticastEventGenerator(words)
	go messages.LocalMessageReady.ShowResults()
	select {}
}
