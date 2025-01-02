package messages

import (
	"encoding/json"
	"log"
	"sort"
)

func MessageCommitJsonUnmarshal(data string) MessageCommitData {
	var n MessageCommitData
	err := json.Unmarshal([]byte(data), &n)
	if err != nil {
		log.Fatalf("Error in Unmarshalling Message: %s", err)
	}
	return n
}

func MessageJsonUnmarshal(data string) MessageData {
	var n MessageData
	err := json.Unmarshal([]byte(data), &n)
	if err != nil {
		log.Fatalf("Error in Unmarshalling Message: %s", err)
	}
	return n
}

func (m *MessageQueue) SortMessages(ts int) {
	if len(m.Messages[ts]) > 1 {
		sort.Slice(m.Messages[ts], func(i, j int) bool {
			return m.Messages[ts][i].Message.Sender < m.Messages[ts][j].Message.Sender
		})
	}
}
