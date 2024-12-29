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
	"os"
	"slices"
	"sort"
	"sync"
	"time"

	"math/rand"
)

type Words struct {
	Words []string `json:"words"`
}

type MessageData struct {
	Timestamp int // lamport clock timestamp
	Sender    string
	Data      string //content of the message
}
type LamportClock struct {
	Timestamp     int
	TimestampList []int
	Mutex         sync.Mutex
}

type MessageQueue struct {
	Messages map[int][]MessageData
	Mutex    sync.Mutex
}

type MessageReady struct {
	Messages []MessageData
	Mutex    sync.Mutex
}

type MessageServer struct {
	msg.UnimplementedMessageServer
}

const Port = ":50001"

var (
	LocalLamportClock LamportClock
	LocalMessageQueue MessageQueue
	LocalMessageReady MessageReady
)

// Grpc function to process the message shared from other nodes.
func (m MessageServer) MessageMulticast(ctx context.Context, in *msg.MessageSharing) (*msg.MessageResponse, error) {
	reqMessage := MessageJsonUnmarshal(in.GetMessage())
	fmt.Printf("Recived message with index:%d sender:%s data:%s\n", reqMessage.Timestamp, reqMessage.Sender, reqMessage.Data)

	// check if the message is a bleat
	//LocalLamportClock.Mutex.Lock()

	if !slices.Contains((LocalLamportClock.TimestampList), reqMessage.Timestamp) {
		LocalLamportClock.TimestampList = append(LocalLamportClock.TimestampList, reqMessage.Timestamp)
	}
	sort.Ints(LocalLamportClock.TimestampList)

	//LocalLamportClock.Mutex.Unlock()
	LocalMessageQueue.Mutex.Lock()
	LocalMessageQueue.Messages[reqMessage.Timestamp] = append(LocalMessageQueue.Messages[reqMessage.Timestamp], reqMessage)
	LocalMessageQueue.Mutex.Unlock()

	LocalMessageQueue.CheckMessageReadinness(reqMessage)

	LocalLamportClock.Timestamp = max(reqMessage.Timestamp, LocalLamportClock.Timestamp) + 1

	return &msg.MessageResponse{Result: "Message sent"}, nil

}
func (m *MessageQueue) CheckMessageReadinness(message MessageData) {
	// all messages are stored in the queue waiting to be released to the ready map
	// Readiness Indicators
	// For a message being ready to be ready requires that the number of messages waiting in the queue is higher or equals to 3
	// The tolerance of 3 messages gives time for the out of order messages to be processed and sorted in time to be displayed.

	LocalLamportClock.Mutex.Lock()
	if len(LocalLamportClock.TimestampList) >= 3 {
		m.Mutex.Lock()
		if len(m.Messages[LocalLamportClock.TimestampList[0]]) > 1 {
			m.SetMessageOrder(LocalLamportClock.TimestampList[0])
		}
		m.Mutex.Unlock()

		// add the previous ts slice to the MessageReady map
		LocalMessageReady.Mutex.Lock()
		LocalMessageReady.Messages = m.Messages[LocalLamportClock.TimestampList[0]]
		LocalMessageReady.Mutex.Unlock()

		LocalLamportClock.TimestampList = LocalLamportClock.TimestampList[1:]
	}
	LocalLamportClock.Mutex.Unlock()

	// 2
}
func (m *MessageQueue) MessageMulticastTrasmission(message MessageData) {

	//message.LogicalCLock = network.LocalTimeVector.LogicalCLock
	message.Timestamp = LocalLamportClock.Timestamp

	fmt.Printf("\nNew Message generated ts:%d sender:%s, data:%s\n", message.Timestamp, message.Sender, message.Data)
	m.Mutex.Lock()
	m.Messages[message.Timestamp] = append(m.Messages[message.Timestamp], message)

	if !slices.Contains((LocalLamportClock.TimestampList), message.Timestamp) {
		LocalLamportClock.TimestampList = append(LocalLamportClock.TimestampList, message.Timestamp)
	}
	sort.Ints(LocalLamportClock.TimestampList)

	m.Mutex.Unlock()

	// send message to online nodes
	for _, value := range network.LocalRouteTable.Nodes {
		messageClient(value.Addr, message)
	}

	//after sending the bleat message insert the message into the ready map
	LocalMessageQueue.CheckMessageReadinness(message)
	LocalLamportClock.Timestamp = LocalLamportClock.Timestamp + 1

}

func (m *MessageQueue) MesageMulticastEventGenerator(words Words) {
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
		log.Printf(config.Green+"Minute:%d Nrequests:%d"+config.Reset, t, nRequests)
		for i := 1; i <= nRequests; i++ {
			totalRequests++

			// get the time for the next request to be executed
			interArrivalTime := poissonProcess.ExponentialRandom()
			previousTime = currentTime
			currentTime = (interArrivalTime * 60) + currentTime
			//log.Printf(config.Green+"Request %d at %f seconds\n"+config.Reset, i, currentTime)
			//log.Printf(config.Green+"Sleep %.5f seconds...\n"+config.Reset, float64(currentTime-previousTime))
			delta := time.Duration(currentTime-previousTime) * time.Second
			time.Sleep(delta)
			rword := words.Words[rand.Intn(len(words.Words))]

			newMessage := MessageData{
				Data:   rword,
				Sender: network.LocalNode.Addr + Port,
			}
			m.MessageMulticastTrasmission(newMessage)

			if i == nRequests && currentTime < 60 {
				log.Printf(config.Green+"Requests for the minute %d endend before finishing the 60s.\nWaiting %f seconds to complete the cycle of 60s....\n"+config.Reset, t, float64(60-currentTime))
				time.Sleep((time.Duration(60-currentTime) * time.Second))
			}
		}
		fmt.Println()
		log.Printf(config.Green+"Statistics: Total requests: %d Minutes spent: %d rate:%f\n"+config.Reset, totalRequests, t, float64(totalRequests)/float64(t))
		t++
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

func (m *MessageQueue) SetMessageOrder(ts int) {
	if len(m.Messages[ts]) > 1 {
		sort.Slice(m.Messages[ts], func(i, j int) bool {
			return m.Messages[ts][i].Sender < m.Messages[ts][j].Sender
		})
	}
}

func (m *MessageReady) ShowResults() {
	//create log file to save messages
	file, _ := os.Create("messages.txt")
	defer file.Close()

	for {
		m.Mutex.Lock()
		if len(m.Messages) > 0 {
			for i := range m.Messages {
				log.Printf(config.Red+"Current Message: Timestamp:%d SenderAddress:%s Data:%s\n"+config.Reset, m.Messages[i].Timestamp, m.Messages[i].Sender, m.Messages[i].Data)
				//save messages on a log file
				log := fmt.Sprintf("Message: Timestamp:%d SenderAddress:%s Data:%s\n", m.Messages[i].Timestamp, m.Messages[i].Sender, m.Messages[i].Data)
				file.WriteString(log)

			}
			m.Messages = nil
		}

		m.Mutex.Unlock()
	}
}
