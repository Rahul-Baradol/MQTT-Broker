package main

import (
	"log"
	"net"
	"sync"
)

type Client struct {
	ID     string
	Conn   net.Conn
	Topics map[string]bool
	Mutex  sync.Mutex
}

type Broker struct {
	Clients    map[string]*Client
	Subscripts map[string]map[*Client]bool
	Mutex      sync.Mutex
}

func NewBroker() *Broker {
	return &Broker{
		Clients:    make(map[string]*Client),
		Subscripts: make(map[string]map[*Client]bool),
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
	defer conn.Close()

	// Implement MQTT protocol here
	// ...

	// Example: Just print the received data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Printf("Failed to read from connection: %v", err)
		return
	}

	_, err = conn.Write([]byte("Data sent successfully!"))

	if err != nil {
		log.Printf("Failed to write data: %v", err)
		return
	}

	log.Printf("Received data: %s", buf[:n])
}

func main() {
	broker := NewBroker()
	broker.Run()
}
