package messages

import (
	"chatapp/config"
	msg "chatapp/messages/proto"
	"chatapp/network"
	"chatapp/poisson"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"math/rand"
)

type Words struct {
	Words []string `json:"words"`
}

type MessageData struct {
	LogicalCLock int // lampert clock timestamp
	Sender       string
	Data         string //content of the message
}
type Messages struct {
	MessageQueue []MessageData // message is a map where the key is the message id
	Mutex        sync.Mutex
}

type MessageServer struct {
	msg.UnimplementedMessageServer
}

const Port = ":50001"

var (
	LocalMessages Messages
)

// Grpc function to process the message shared from other nodes.
func (m MessageServer) MessageMulticast(ctx context.Context, in *msg.MessageSharing) (*msg.MessageResponse, error) {
	reqMessage := MessageJsonUnmarshal(in.GetMessage())
	fmt.Printf("Recived message with index:%d sender:%s data:%s", reqMessage.LogicalCLock, reqMessage.Sender, reqMessage.Data)
	LocalMessages.Mutex.Lock()
	network.LocalTimeVector.Mutex.Lock()
	LocalMessages.MessageQueue = append(LocalMessages.MessageQueue, reqMessage)
	network.LocalTimeVector.LogicalCLock = max(reqMessage.LogicalCLock, network.LocalTimeVector.LogicalCLock) + 1
	fmt.Printf("Logical clock:%d\n", network.LocalTimeVector.LogicalCLock)

	// check which of the times is higher updates the local if it is lower
	//update the time vector
	network.LocalTimeVector.Vector[reqMessage.Sender] = reqMessage.LogicalCLock

	LocalMessages.Mutex.Unlock()
	network.LocalTimeVector.Mutex.Unlock()

	//fmt.Printf("New message arriverd! => %d %s, %s\n", reqMessage.Timestamp, reqMessage.Sender, reqMessage.Data)

	// return status ok after processing the message shared
	return &msg.MessageResponse{Result: "Message sent"}, nil
}

func (mm *Messages) MesageMulticastEventGenerator(words Words) {
	// sleep to wait that all the nodes are sync
	time.Sleep(time.Second * 20)
	seed := time.Now().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	lambda := 60.0 // rate of 2 updates per minute
	poissonProcess := poisson.PoissonProcess{Lambda: lambda, Rng: rng}
	totalRequests := 0
	t := 1 // current minute
	for {
		currentTime := 0.0
		previousTime := 0.0

		nRequests := poissonProcess.PoissonRandom()
		for i := 1; i <= nRequests; i++ {
			totalRequests++
			log.Printf(config.Green+"Minute:%d Nrequests:%d"+config.Reset, t, nRequests)

			// get the time for the next request to be executed
			interArrivalTime := poissonProcess.ExponentialRandom()
			previousTime = currentTime
			currentTime = (interArrivalTime * 60) + currentTime
			log.Printf(config.Green+"Request %d at %f seconds\n"+config.Reset, i, currentTime)
			log.Printf(config.Green+"Sleep %.5f seconds...\n"+config.Reset, float64(currentTime-previousTime))
			delta := time.Duration(currentTime-previousTime) * time.Second
			time.Sleep(delta)
			rword := words.Words[rand.Intn(len(words.Words))]

			newMessage := MessageData{
				Data:   rword,
				Sender: network.LocalNode.Addr + Port,
			}

			mm.Mutex.Lock()
			network.LocalTimeVector.Mutex.Lock()
			fmt.Printf("Logical clock:%d\n", network.LocalTimeVector.LogicalCLock)
			newMessage.LogicalCLock = network.LocalTimeVector.LogicalCLock
			network.LocalTimeVector.LogicalCLock++

			//To ensure that other treads trey to claim data from message queue before being sync with the current local time is under the lock mutex
			mm.MessageQueue = append(mm.MessageQueue, newMessage)
			log.Printf(config.Cyan+"New Local Message Generated: Timestamp:%d LocalAddress:%s Data:%s\n"+config.Reset, newMessage.LogicalCLock, newMessage.Sender, newMessage.Data)
			mm.Mutex.Unlock()

			MessageMulticastTrasmission(newMessage)

			network.LocalTimeVector.Mutex.Unlock()
			if i == nRequests && currentTime < 60 {
				log.Printf(config.Green+"Requests for the minute %d endend before finishing the 60s.\nWaiting %f seconds to complete the cycle of 60s....\n"+config.Reset, t, float64(60-currentTime))
				time.Sleep((time.Duration(60-currentTime) * time.Second))
			}

		}
		fmt.Println()
		log.Printf(config.Green+"Statistics: Total requests: %d Minutes spent: %d rate:%f\n"+config.Reset, totalRequests, t, float64(totalRequests)/float64(t))
		t++
		//time.Sleep(time.Second * 5)
	}
}
func MessageMulticastTrasmission(message MessageData) {
	for _, value := range network.LocalRouteTable.Nodes {
		if value.Status == "online" {
			messageClient(value.Addr, message)
		}
	}
}

func MessageJsonUnmarshal(data string) MessageData {
	var n MessageData
	err := json.Unmarshal([]byte(data), &n)
	if err != nil {
		log.Fatalf("Error in Unmarshalling Message: %s", err)
	}
	return n
}

//func (m *Messages) ShowResults() {
//	ephemeralCounter := 0
//	modCounter := 0
//	for {
//		if len(m.MessageQueue) > 0 && network.LocalTimeVector.LogicalCLock != ephemeralCounter {
//			network.LocalTimeVector.Mutex.Lock()
//			if len(m.MessageQueue) == network.LocalTimeVector.LogicalCLock {
//				log.Printf(config.Red+"Current Message: Timestamp:%d SenderAddress:%s Data:%s\n"+config.Reset, m.MessageQueue[ephemeralCounter].LogicalCLock, m.MessageQueue[ephemeralCounter].Sender, m.MessageQueue[ephemeralCounter].Data)
//
//			} else if len(m.MessageQueue) < network.LocalTimeVector.LogicalCLock {
//				// in a scenario where the node lost its connection to the network an rejoins the message queue will not match with the local time.
//				// therefore, the node requires to match the local time with the network. since the local time > len message queue then is calculated the module of that number with the ephemeral counter.
//				modCounter = modCounter % network.LocalTimeVector.LogicalCLock
//				fmt.Printf("new run!!!!!!!!!!!!")
//				log.Printf(config.Red+"Current Message: Timestamp:%d SenderAddress:%s Data:%s\n"+config.Reset, m.MessageQueue[modCounter].LogicalCLock, m.MessageQueue[modCounter].Sender, m.MessageQueue[modCounter].Data)
//				modCounter++
//			}
//			ephemeralCounter = network.LocalTimeVector.LogicalCLock
//
//			network.LocalTimeVector.Mutex.Unlock()
//		}
//	}
//}

func (m *Messages) ShowResults() {
	for {
		if len(m.MessageQueue) > 0 {
			log.Printf(config.Red+"Current Message: Timestamp:%d SenderAddress:%s Data:%s\n"+config.Reset, m.MessageQueue[network.LocalTimeVector.LogicalCLock].LogicalCLock, m.MessageQueue[network.LocalTimeVector.LogicalCLock].Sender, m.MessageQueue[network.LocalTimeVector.LogicalCLock].Data)
		}
	}
}

// send messages over udp multicast group messages with data or bleit messages
// check if the message for a certain id reach timeout validade which nodes sent ack messages.
// retransmit for nodes whom did´t reply.
// max retry is 3 if no ack node is removed from route table.
// handle recived messages
// display messages under certain rules\
