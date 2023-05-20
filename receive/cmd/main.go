package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
	"sync"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	app := fiber.New()

	app.Post("/endpoint", handleRequest)

	log.Fatal(app.Listen(":8088"))
}

func handleRequest(c *fiber.Ctx) error {
	// Parse the request JSON payload
	//var req string
	body := c.Body()
	var wg sync.WaitGroup
	numConnections := 10 // Number of concurrent connections to establish
	wg.Add(numConnections)

	for i := 0; i < numConnections; i++ {
		go func(payload []byte) {
			defer wg.Done()

			// Open a TCP socket connection to the other service
			conn, err := net.Dial("tcp", "service-address:service-port") // Replace with the actual address and port of the other service
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

			log.Println("Payload sent successfully")
		}(body)
	}

	wg.Wait()
	//if err != nil {
	//	return fiber.NewError(fiber.StatusBadRequest, err.Error())
	//}

	// Create the response
	//resp := Response{
	//	Status:  "success",
	//	Message: "Request processed successfully",
	//}
	//
	//// Convert the response to JSON
	//respJSON, err := json.Marshal(resp)
	//if err != nil {
	//	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	//}
	//
	//// Send the request to another service using a socket connection
	//err = sendRequestToService(string(respJSON))
	//if err != nil {
	//	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	//}

	// Send the response back
	//fmt.Println(string(req))
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

	// Read the response from the other service (if applicable)
	// response := make([]byte, bufferSize)
	// _, err = conn.Read(response)
	// if err != nil {
	// 	return err
	// }

	return nil
}
