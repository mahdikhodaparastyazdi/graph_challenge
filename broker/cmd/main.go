package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Start the server
	go startServer()

	// Wait for user input to exit
	fmt.Println("Press enter to exit.")
	fmt.Scanln()
}

func startServer() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8888") // Replace with the desired listening address and port
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("Server started, listening on", listener.Addr())

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		// Handle connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	log.Printf("New connection from %s", conn.RemoteAddr())

	// Read incoming data from the connection
	reader := bufio.NewReader(conn)
	payload, err := reader.ReadBytes('\n')
	if err != nil {
		log.Printf("Failed to read data from connection: %v", err)
		return
	}

	// Log the received payload
	log.Printf("Received payload: %s", payload)

	// Open a TCP socket connection to another service
	destConn, err := net.Dial("tcp", "destination-address:destination-port") // Replace with the actual address and port of the destination service
	if err != nil {
		log.Printf("Failed to open TCP connection to destination: %v", err)
		return
	}
	defer destConn.Close()

	// Send the payload to the destination service
	_, err = destConn.Write(payload)
	if err != nil {
		log.Printf("Failed to send payload to destination: %v", err)
		return
	}

	log.Println("Payload sent to destination successfully")
}
