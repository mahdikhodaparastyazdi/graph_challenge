package main

import (
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
	dataChan := make(chan net.Conn)
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8088") // Replace with the desired listening address and port
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer listener.Close()

	log.Println("Server started, listening on", listener.Addr())

	// Handle connection in a separate goroutine
	//go handleConnection(conn)
	var num = 0
	go func() {
		buffer := make([]byte, 1000000)
		for conn := range dataChan {

			// Create a buffer to store the received data
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Failed to read data:", err)
				conn.Close()
				continue
			}

			payload := buffer[:n]
			go sendToLog(payload)
			num++
			//// Log the received payload
			//
			////Open a TCP socket connection to another service
			destConn, err := net.Dial("tcp", "127.0.0.1:8888")
			if err != nil {
				log.Printf("Failed to open TCP connection to destination: %v", err)
				return
			}

			// Send the payload to the destination service
			_, err = destConn.Write(payload)
			if err != nil {
				log.Printf("Failed to send payload to destination: %v", err)
				return
			}
			destConn.Close()
			conn.Close()
		}

	}()
	var nu = 0
	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		dataChan <- conn

		fmt.Println(nu)
	}
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	log.Print("New connection from", conn.RemoteAddr())

	// Read incoming data from the connection
	//reader := bufio.NewReader(conn)
	//payload, err := reader.ReadBytes('\n')
	//if err != nil {
	//	log.Printf("Failed to read data from connection: %v", err)
	//	return
	//}

	buffer := make([]byte, 4096) // Create a buffer to store the received data
	n, err := conn.Read(buffer)
	if err != nil {
		//fmt.Println("Failed to read data:", err)
		return
	}
	payload := buffer[:n]
	// Log the received payload

	//Open a TCP socket connection to another service
	destConn, err := net.Dial("tcp", "127.0.0.1:8888")
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

	//	log.Println("Payload sent to destination successfully")
}
func sendToLog(payload []byte) {
	log.Print(payload)
}
