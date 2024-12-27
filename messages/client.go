package messages

import (
	msg "chatapp/messages/proto"
	"context"
	"encoding/json"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func messageClient(addr string, message MessageData) {
	conn, err := grpc.NewClient(addr+Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error connecting to server %s", err)
	}

	c := msg.NewMessageClient(conn)
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	messageJson, _ := json.Marshal(message)
	c.MessageMulticast(ctx, &msg.MessageSharing{Message: string(messageJson)})
}
