package main

import (
	"GoMulticastChatApp/messages"
	"GoMulticastChatApp/network"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func cli() string {
	// Define flags
	id := flag.String("id", "", "Specify the Local Node ID")

	help := flag.Bool("help", false, "help")

	// Parse flags
	flag.Parse()

	if *help {
		fmt.Println("Usage: cli [options]")
		flag.PrintDefaults()
		os.Exit(0)
	}

	return *id
}

func main() {
	id := cli()
	network.LocalID = id
	localAddr, _ := network.GetLocalAddr()
	network.LocalNode = network.Node{
		ID:     id,
		Addr:   localAddr,
		Status: "offline",
	}
	messages.LocalLamportClock.Timestamp = 0
	messages.LocalMessageQueue.Messages = make(map[int][]messages.MessageCommited)
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
	go messages.LocalMessageQueue.ShowResults()
	select {}
}
