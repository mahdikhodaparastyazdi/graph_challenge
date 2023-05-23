package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	dataChan := make(chan []byte, 1000000)
	/*

		fmt.Println("main:started")

		app := fiber.New()

		app.Post("/endpoint", handleRequest)

		log.Fatal(app.Listen(":8085"))
	*/

	app := fiber.New()

	// Create a channel to receive data from the POST endpoint

	// Use a wait group to ensure all goroutines finish before shutting down
	//var wg sync.WaitGroup

	// Start a goroutine to send data from the channel to the socket
	//wg.Add(1)
	app.Post("/endpoint", func(c *fiber.Ctx) error {
		body := c.Body()
		go sendValue(body, dataChan)
		//dataChan <- body
		return c.Status(fiber.StatusOK).JSON("")
	})

	var num = 0
	//address := "127.0.0.1:8088"
	//conn, err := net.Dial("tcp", address)
	//if err != nil {
	//	log.Fatal("Failed to connect to the socket:", err)
	//}

	go func() {
		fmt.Println("start go func")
		for payload := range dataChan {
			//if num == 190 {
			//	close(dataChan)
			//}
			if payload != nil {
				fmt.Println(num)
				num++
			}
			go sendToSocket(payload)
			//_, err = conn.Write(payload)
			//if err != nil {
			//	log.Printf("Failed to send payload: %v", err)
			//	return
			//}
		}
	}()
	log.Fatal(app.Listen(":8085"))
	// Define the endpoint to receive POST requests
	//app.Post("/endpoint", func(c *fiber.Ctx) error {
	//	// Read the body of the POST request
	//	body := c.Body()
	//	// Send the received data to the channel
	//	dataChan <- body
	//	// Respond with a success status
	//	return c.SendString("Data received")
	//})
	//go func() {
	//	if err := app.Listen(":8085"); err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	//address := "127.0.0.1:8088"
	//conn, err := net.Dial("tcp", address)
	//if err != nil {
	//	log.Fatal("Failed to connect to the socket:", err)
	//}
	//defer conn.Close()

	//go func() {
	//	defer wg.Done()
	//
	//	for data := range dataChan {
	//		// Send the received data to the socket
	//		_, err := conn.Write([]byte(data))
	//		if err != nil {
	//			log.Println("Failed to send data to socket:", err)
	//		} else {
	//			log.Println("Data sent to socket successfully")
	//		}
	//	}
	//
	//}()
	// Wait for all goroutines to finish before shutting down
	//wg.Wait()

}
func sendValue(body []byte, c chan []byte) {
	c <- body
}
func sendToSocket(payload []byte) {
	address := "127.0.0.1:8088"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Failed to connect to the socket:", err)
	}
	_, err1 := conn.Write(payload)
	if err1 != nil {
		log.Printf("Failed to send payload: %v", err)

	}
}

func handleRequest(c *fiber.Ctx) error {

	// Parse the request JSON payload
	//var req string
	//body := c.Body()
	//dataChan <- string(body)
	//fmt.Println("handleRequest:body", string(body))
	//var wg sync.WaitGroup
	//numConnections := 10 // Number of concurrent connections to establish
	//wg.Add(numConnections)

	//go func() {
	//
	//	err := sendData("127.0.0.1:8088", body)
	//	if err != nil {
	//		fmt.Println("Failed to send data:", err)
	//	}
	//}()

	//for i := 0; i < numConnections; i++ {
	/*
		go func(payload []byte) {

			//defer wg.Done()
			// Open a TCP socket connection to the other service
			conn, err := net.Dial("tcp", "127.0.0.1:8088")
			if err != nil {
				log.Printf("Failed to open TCP connection: %v", err)
				return
			}
			defer conn.Close()
			// Send the payload to the other service
			_, err = conn.Write(payload)
			if err != nil {
				log.Printf("Failed to send payload: %v", err)
				return
			}
			//dataChan <- string(payload)

			//log.Println("Payload sent successfully")
		}(body)
		//}
		//wg.Wait()

	*/

	return c.Status(fiber.StatusOK).JSON("ok")
}

func sendRequestToService(data string) error {
	// Create a TCP socket connection to the other service
	conn, err := net.Dial("tcp", "service-address:service-port")
	if err != nil {
		return err
	}
	defer conn.Close()

	// Send the data to the other service
	_, err = conn.Write([]byte(data))
	if err != nil {
		return err
	}

	return nil
}
