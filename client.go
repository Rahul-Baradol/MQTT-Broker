package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type ProducerData struct {
	Topic   string
	Message string
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1883")
	if err != nil {
		log.Fatalf("Failed to connect to broker: %v", err)
	}
	defer conn.Close()

	var producer_consumer int32

	fmt.Print("If you are a producer, type 1 and if you are a consumer, type 2 else type any number to exit: ")
	fmt.Scanln(&producer_consumer)

	if producer_consumer == 1 {
		// Producer
		for {
			fmt.Print("Enter the topic you want to publish to: ")
			var topic string
			fmt.Scanln(&topic)
			fmt.Print("Enter the message you want to publish: ")
			var message string
			fmt.Scanln(&message)

			data := &ProducerData{ Topic: topic, Message: message }

			marshaledData, err := json.Marshal(data)
			if err != nil {
				log.Fatalf("Failed to marshal data: %v\n", err)
				return
			}
			_, err = conn.Write(marshaledData)

			if err != nil {
				log.Printf("Failed to write data: %v\n", err)
				return
			}

			log.Println("Data sent successfully!")
		}
	} else if producer_consumer == 2 {
		// Consumer
		fmt.Println("Please enter the topic you want to subscribe to: ")
		var topic string
		fmt.Scanln(&topic)
		
		_, err = conn.Write([]byte(topic))
		if err != nil {
			log.Printf("Failed to write data: %v\n", err)
			return
		}

		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			
			if err != nil {
				log.Printf("Failed to read from connection: %v\n", err)
				return
			}
			
			fmt.Printf("Received message from broker: %s\n", buf[:n])
		}
	} else {
		return
	}

	// Implement MQTT protocol here
	// ...

	// Example: Send some data to the broker
	// data := []byte("Hello, MQTT broker!")
	// _, err = conn.Write(data)
	// if err != nil {
	// 	log.Printf("Failed to write data: %v", err)
	// 	return
	// }

	// Read the success message from the broker
	// buf := make([]byte, 1024)
	// n, err := conn.Read(buf)
	// if err != nil {
	// 	log.Printf("Failed to read from connection: %v", err)
	// 	return
	// }

	// fmt.Printf("Received message from broker: %s\n", buf[:n])
}
