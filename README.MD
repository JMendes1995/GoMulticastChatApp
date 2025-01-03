# GoMulticastChatApp

# Table Of Contents

- [GoMulticastChatApp](#gomulticastchatapp)
  - [Requirements](#requirements)
  - [How to build the project](#how-to-build-the-project)
  - [How to execute the project](#how-to-execute-the-project)
    - [Execution in physical nodes](#execution-in-physical-nodes)
    - [Docker](#docker)
      - [Clean up docker environment](#clean-up-docker-environment)

This project is a Basic Chat Application Using Totally-Ordered Multicast usig Go programming language for the scope of Distributed systems Class final project.

### Requirements
* Create a network of 6 peers (p1 to p6) each in a different machine (m1 to m6) with the
topology shown in the figure above. Each peer has a table with the IPs of all the other
peers/machines.
* Implement Lamport clocks in each peer so that messages can be adequately times-
tamped.
* Each peer sends a random word according to a Poisson distribution with a frequency
of 1 per second. You can use a text format dictionary from the Internet and extract its
keywords into a set. Then choose a random keyword from that set. Each word is sent
in a message to all other peers (they are all in the IP table).
* The goal is that all peers print the same sequence of words. To do this, the peers must
agree on a global order for the messages before they process them (print the words
therein). This is not trivial, as factors such as communication latency and varying
network topology affect message delivery even in small networks.
* To achieve this you must implement the Totally-Ordered Multicast (TOM) Algorithm
using Lamport clocks to timestamp the messages. Check Chapter 6 of van Steen and
Tanenbaum for the details of the algorithm and here for another, detailed, description).
* Note that in TOM there is a difference between receiving and processing a message.
Processes always receive incoming messages but they are processed only when specific
conditions are met. In this application, peers process messages by printing the words
therein. If you correctly implemented the TOM algorithm the printed list of words must
be the same for all peers.

### How to build the project
In the root directory of the project are present 2 binaries from **Linux (GoMulticastChatApp-linux)** and **MacOS (GoMulticastChatApp-macos)** arm architecture. However is possible to build the project manually by executing the following commands.

```bash
$ cd GoMulticastChatApp
$ go build
```

### How to execute the project

To execute the project there is 2 possible ways.
1. Execute the code in phsical nodes
2. Build docker environment using docker compose.

#### Execution in physical nodes
Before executing the project, is required to pull the source code into the nodes that will be part of the network.
Furthermore is require to update the **nodes.json** file located at the root of the project with the node **address** and **id** of each node. 
Finnaly, to execute the project is also required to execute the following command in each node.

```bash
$ cd GoMulticastChatApp
$ ./GoMulticastChatApp-linux 
```

#### Docker
To execute the project using docker is required to have the docker and docker compose prior installed  Then is required to execute the following commands.

1. Setup the network and the nodes.
```bash
$ cd GoMulticastChatApp
$ docker-compose -d
```
2. In different windows, enter in each container.
```bash
$ docker exec -it gomulticastchatapp-peer1-1 bash 
```
```bash
$ docker exec -it gomulticastchatapp-peer2-1 bash 
```
```bash
$ docker exec -it gomulticastchatapp-peer3-1 bash 
```
```bash
$ docker exec -it gomulticastchatapp-peer4-1 bash 
```
```bash
$ docker exec -it gomulticastchatapp-peer5-1 bash 
```
```bash
$ docker exec -it gomulticastchatapp-peer6-1 bash 
```

3. Execute binary in every container
```bash
$ ./GoMulticastChatApp-linux 
```
##### Clean up docker environment
```bash
$ docker compose down && docker compose rm && docker rmi -f $(docker images -aq)
```


### Result intrepertation
Note: The list of generated messages is at the node in a file named messages-<nodeID> in the root directory of the project.

##### Server startup and Peer Registration
Nodes initially are in a offline state. Therefore, before sending messages between them, nodes start by sending a message to anounce themselves to the network resulting in state change from **offline** to **Online** state.
Nodes which are online state are placed in a routing table and if a node stops responding it will result in a removal from the routing table.

```bash
2025/01/03 09:42:01 Network server P2P server listening at 192.168.0.142:50000
2025/01/03 09:42:01 Message server listening at 192.168.0.142:50001
2025/01/03 09:42:01 Peer not ready 192.168.0.144:50000 Remaining connection attempts:3
2025/01/03 09:42:01 Peer not ready 192.168.0.143:50000 Remaining connection attempts:3
2025/01/03 09:42:01 Peer not ready 192.168.0.147:50000 Remaining connection attempts:3
2025/01/03 09:42:01 Peer not ready 192.168.0.146:50000 Remaining connection attempts:3
2025/01/03 09:42:01 Peer not ready 192.168.0.145:50000 Remaining connection attempts:3
2025/01/03 09:42:06 192.168.0.146:50000 result:"Node status changed"
2025/01/03 09:42:06 192.168.0.144:50000 result:"Node status changed"
2025/01/03 09:42:06 192.168.0.145:50000 result:"Node status changed"
2025/01/03 09:42:06 192.168.0.147:50000 result:"Node status changed"
2025/01/03 09:42:06 192.168.0.143:50000 result:"Node status changed"
```

##### Event generation
After update the peer status in the local routing table, The event generator starts producing messages at a rate of 1 message per second following a Poisson distribution.
Each new message generated will increment the logical clock (Lamport clock) by 1, stores in a local queue and transmit the message to other peers in multicast using gRPC.
With the new message generated the node will distribute the message in a multicasst 
```bash
2025/01/03 09:42:11 Minute:1 Nrequests:58
2025/01/03 09:42:11 New Message generated Timestamp:0 Address:192.168.0.142:50001 Data:Knapsack
```

##### Commit message by the message sender.
For every node know that a message is ready to be processed (display in the screen), Is required to send confirmation messages confirming that a message was successfully transmitted. Therefore, The owner of the generated message 
is responsible to generate and distribute commit messages peer peer for each message that generates. This mechanism ensures that only messages that are successfully sent and confirmed by the message ower can be process.
Using commit messages instead of ACK messages by the peers broacasted to the entire network reduces the number of messages per message.
In this approach will result per message (1+nNodes) messages. On the other hand if we used ACK messages broadcasted to the internet we would have (1 + nNodes²) which is less scalable in larger networks and in netowrks with high throuput.

```bash
2025/01/03 09:42:11 Sending Commit message to node 192.168.0.146 for message {Timestamp: 0, Sender: 192.168.0.142:50001, Data:Knapsack}
2025/01/03 09:42:11 Sending Commit message to node 192.168.0.147 for message {Timestamp: 0, Sender: 192.168.0.142:50001, Data:Knapsack}
2025/01/03 09:42:11 Sending Commit message to node 192.168.0.143 for message {Timestamp: 0, Sender: 192.168.0.142:50001, Data:Knapsack}
2025/01/03 09:42:11 Sending Commit message to node 192.168.0.144 for message {Timestamp: 0, Sender: 192.168.0.142:50001, Data:Knapsack}
2025/01/03 09:42:11 Sending Commit message to node 192.168.0.145 for message {Timestamp: 0, Sender: 192.168.0.142:50001, Data:Knapsack}
2025/01/03 09:42:11 Processing Message: Timestamp:0 SenderAddress:192.168.0.142:50001 Data:Knapsack
```

##### Commit message by the message sender.
When a node recives a new message, it is inserted in the local queue and awaits for the commit message from the message owner. Once the commit message arrives, the node marks the message as commited and pushes the message to a message ready map which is regurarly checked to be process.
```bash
2025/01/03 09:42:11 Recived message with Timestamp:1 Address:192.168.0.143:50001 Data:Jelly
2025/01/03 09:42:11 Recived Commit from 192.168.0.143 with message Timestamp:1 Address:192.168.0.143:50001 Data:Jelly
2025/01/03 09:42:11 Processing Message: Timestamp:1 SenderAddress:192.168.0.143:50001 Data:Jelly
```

##### Concurrent messages.
In Network where every node is competing for resouces which means sending messages to each other, concurent messages might occour. Therefore, In a Tottaly order multicast algorithm is required that all messages have the same order in every node.
Although exists mechanisms that allow clock syncronization, namely the lamport clocks, this does not prevent entirely the generation of concurrent messages. Thus, was implemented a tie-break scenario which will sort concurrent messages by the address of the sender before the message being process by the node. Thereby, all nodes will process the same message.

```bash
2025/01/03 09:42:11 Recived message with Timestamp:1 Address:192.168.0.142:50001 Data:Jelly
2025/01/03 09:42:11 Recived message with Timestamp:1 Address:192.168.0.143:50001 Data:Jelly
```