# GoMulticastChatApp

# Table Of Contents

- [GoMulticastChatApp](#gomulticastchatapp)
  - [Requirements](#requirements)

This project is a Basic Chat Application Using Totally-Ordered Multicast usig Go programming language for the scope of Distributed systems Class final project. This project 

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
In the developent phase of the project it was used a docker environment where the nodes and networks setup where virtualize to create a more seamlessly envorinment as possible.

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


