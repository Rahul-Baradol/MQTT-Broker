package main

import (
	"encoding/json"
	"log"
	"net"
	"sync"
)

type ClientData struct {
	ClientType string
	Topic      string
	Message    string
}

type Client struct {
	ID         string
	Conn       net.Conn
	TopicIndex int32 // index -> indices for Topics array
	Mutex      sync.Mutex
}

type Broker struct {
	Topics      map[string][]string  // map[topic] = [data1, data2, ...]
	Subscribers map[string][]*Client // map[topic] = [client1, client2, ...]
	Mutex       sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		// Clients:    make(map[string]*Client),
		// Subscripts: make(map[string]map[*Client]bool),

		Topics:      make(map[string][]string),
		Subscribers: make(map[string][]*Client),
	}
}

func (b *Broker) Run() {
	listener, err := net.Listen("tcp", ":1883")
	if err != nil {
		log.Fatalf("Failed to start broker: %v", err)
	}

	defer listener.Close()
	log.Println("MQTT broker started, listening on :1883")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go b.handleConnection(conn)
	}
}

func (b *Broker) handleConnection(conn net.Conn) {
	// Implement MQTT protocol here
	// ...

	// Example: Just print the received data
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("Failed to read from connection: %v", err)
			return
		}

		var data ClientData
		err = json.Unmarshal(buf[:n], &data)

		if err != nil {
			log.Printf("Failed to unmarshal data: %v", err)
			return
		}

		log.Printf("Received data: %v", data)
		conn.Write([]byte("Data received"))
	}
}

func main() {
	broker := NewBroker()
	broker.Run()
}
