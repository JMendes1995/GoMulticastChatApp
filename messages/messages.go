package messages

import (
	"GoMulticastChatApp/config"
	msg "GoMulticastChatApp/messages/proto"
	"GoMulticastChatApp/network"
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
	"time"
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
	Messages map[int][]MessageCommited
	Mutex    sync.Mutex
}

type MessageCommited struct {
	Message  MessageData
	Commit   bool
	AckNodes []string
}

type MessageCommitData struct {
	Sender  string
	Message MessageData
}

type MessageServer struct {
	msg.UnimplementedMessageServer
}

const Port = ":50001"

var (
	LocalLamportClock LamportClock
	LocalMessageQueue MessageQueue
)

// Grpc function to process the message shared from other nodes.
// When a node reciveds this message it means that a new message was generated from other nodes
// the new message is added to the local message queue queue
// the logical clock is incremented  max(recived message timestamp, current logical clock) + 1
func (m MessageServer) MessageMulticast(ctx context.Context, in *msg.MessageSharing) (*msg.MessageResponse, error) {
	reqMessage := MessageJsonUnmarshal(in.GetMessage())
	log.Printf(config.Yellow+"Recived message with Timestamp:%d Address:%s Data:%s\n"+config.Reset, reqMessage.Timestamp, reqMessage.Sender, reqMessage.Data)

	LocalLamportClock.Mutex.Lock()
	LocalLamportClock.Timestamp = max(reqMessage.Timestamp, LocalLamportClock.Timestamp) + 1
	LocalLamportClock.TimestampList = append(LocalLamportClock.TimestampList, reqMessage.Timestamp)
	LocalLamportClock.Mutex.Unlock()

	LocalMessageQueue.Mutex.Lock()

	LocalMessageQueue.Messages[reqMessage.Timestamp] = append(LocalMessageQueue.Messages[reqMessage.Timestamp], MessageCommited{
		Message: reqMessage,
		Commit:  false,
	})

	LocalMessageQueue.Mutex.Unlock()

	return &msg.MessageResponse{Result: "Message sent"}, nil
}

// Grpc function to handle the commit messages
// When a node recives this commit message it automatically means that this message was delivered to every node in the network and is ready to be process.
// This is achieved by setting the commit flag to true
func (m MessageServer) MessageCommitMulticast(ctx context.Context, in *msg.MessageCommitSharing) (*msg.MessageCommitResponse, error) {
	reqAckMessage := MessageCommitJsonUnmarshal(in.GetMessage())
	log.Printf(config.Yellow+"Recived Commit from %s with message Timestamp:%d Address:%s Data:%s\n"+config.Reset, reqAckMessage.Sender, reqAckMessage.Message.Timestamp, reqAckMessage.Message.Sender, reqAckMessage.Message.Data)

	LocalMessageQueue.Mutex.Lock()
	for i := range LocalMessageQueue.Messages[reqAckMessage.Message.Timestamp] {
		if LocalMessageQueue.Messages[reqAckMessage.Message.Timestamp][i].Message == reqAckMessage.Message {
			LocalMessageQueue.Messages[reqAckMessage.Message.Timestamp][i].Commit = true
		}
	}
	LocalMessageQueue.Mutex.Unlock()
	return &msg.MessageCommitResponse{Result: "Commit sent"}, nil
}


//After the event generator generates a new message it will be transmitter in multicast to every node in the network.
//After the messege being transmited is required in order to process the message every node recieve a commit message 
//This commit message is the confirmation tha every node recieved the message and is ok to process (print on the screen)
func (m *MessageQueue) MessageMulticastTrasmission(message MessageData) {
	LocalLamportClock.Mutex.Lock()
	//update the message timestamp with the current logical clock
	message.Timestamp = LocalLamportClock.Timestamp
	//add the new message timestamp in a list of every timestamps waiting to be process.
	LocalLamportClock.TimestampList = append(LocalLamportClock.TimestampList, message.Timestamp)
	LocalLamportClock.Mutex.Unlock()

	log.Printf(config.Cyan+"New Message generated Timestamp:%d Address:%s Data:%s\n"+config.Reset, message.Timestamp, message.Sender, message.Data)

	// add the generated message to the local queue with the commit flag set to false.
	// Once the flag commit changes to true it means that the message is ready to be process.
	m.Mutex.Lock()
	m.Messages[message.Timestamp] = append(m.Messages[message.Timestamp], MessageCommited{
		Message: message,
		Commit:  false,
	})
	m.Mutex.Unlock()

	// send message to online nodes
	for _, value := range network.LocalRouteTable.Nodes {
		messageClient(value.Addr, message)
	}
	LocalLamportClock.Timestamp = LocalLamportClock.Timestamp + 1

	// after sending all the messages successfully send an commit message to all nodes
	// confirming that the message was transmitted and is ready to process. 
	// This is achieved by sending commit messages to every node to confirm that the transaction was successfull.
	//
	ackMsg := MessageCommitData{
		Sender:  network.LocalNode.Addr,
		Message: message,
	}
	for _, value := range network.LocalRouteTable.Nodes {
		log.Printf(config.Cyan+"Sending Commit message to node %s for message {Timestamp: %d, Sender: %s, Data:%s}\n"+config.Reset, value.Addr, message.Timestamp, message.Sender, message.Data)
		messageAckClient(value.Addr, ackMsg)
	}
	m.Mutex.Lock()

	// after commiting the message set the commit flag to try to mark the message as ready to process.
	for i := range m.Messages[message.Timestamp] {
		if m.Messages[message.Timestamp][i].Message == message {
			fmt.Printf("Comit completed message ready")
			m.Messages[message.Timestamp][i].Commit = true
		}
	}
	m.Mutex.Unlock()
}

// function that is always looking at the oldest message in the queue.
// When the oldest message was the commit flag to true it is displayed and saved into a log file.
// when concurent messages occount both are added to the localqueue and ordered by the the address of the noded that generated the message.
func (m *MessageQueue) ShowResults() {
	//create log file to save messages
	file, _ := os.Create(fmt.Sprintf("messages-%s.txt", network.LocalNode.ID))
	defer file.Close()

	for {
		if len(LocalLamportClock.TimestampList) > 0 {
			sort.Ints(LocalLamportClock.TimestampList)
			m.Mutex.Lock()
			if len(m.Messages[LocalLamportClock.TimestampList[0]]) > 0 {
				m.SortMessages(LocalLamportClock.TimestampList[0])
				if m.Messages[LocalLamportClock.TimestampList[0]][0].Commit == true {

					log.Printf(config.Red+"Processing Message: Timestamp:%d SenderAddress:%s Data:%s\n"+config.Reset,
						m.Messages[LocalLamportClock.TimestampList[0]][0].Message.Timestamp,
						m.Messages[LocalLamportClock.TimestampList[0]][0].Message.Sender,
						m.Messages[LocalLamportClock.TimestampList[0]][0].Message.Data)
					//save messages on a log file
					log := fmt.Sprintf("Message: Timestamp:%d SenderAddress:%s Data:%s\n",
						m.Messages[LocalLamportClock.TimestampList[0]][0].Message.Timestamp,
						m.Messages[LocalLamportClock.TimestampList[0]][0].Message.Sender,
						m.Messages[LocalLamportClock.TimestampList[0]][0].Message.Data)
					file.WriteString(log)
					m.Messages[LocalLamportClock.TimestampList[0]] = m.Messages[LocalLamportClock.TimestampList[0]][min(len(m.Messages[LocalLamportClock.TimestampList[0]]), 1):]
				}

			} else {
				LocalLamportClock.TimestampList = LocalLamportClock.TimestampList[1:]
			}
			m.Mutex.Unlock()
		}
		time.Sleep(time.Millisecond * 50)

	}
}
