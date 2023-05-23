package main

import (
	"fmt"
	"github.com/sender/pkg/service"
	"math/rand"
	"sync"
	"time"
)

func main() {

	numRequests := 10000
	concurrency := 400 // Number of goroutines to run concurrently
	//requestsPerSecond := 10000

	// Create a wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numRequests)

	// Create a channel to limit the concurrency
	semaphore := make(chan struct{}, concurrency)

	startTime := time.Now()

	// Start the goroutines
	rand.Seed(time.Now().UnixNano())
	sizeInBytes := rand.Intn(7950) + 50

	// Generate the JSON payload
	jsonPayload, err := service.GenerateJSONPayload(sizeInBytes)
	for i := 0; i < numRequests; i++ {
		go func() {
			defer wg.Done()
			// Acquire a semaphore to limit concurrency
			semaphore <- struct{}{}

			// Make the HTTP request
			////////////////////////////////////////////////////////////////
			//err := makeRequest()
			//if err != nil {
			//	fmt.Println("Error making request:", err)
			//}
			////////////////////////////////////////////////////////////////
			// Make the HTTP request

			if err != nil {
				fmt.Println("Error generating JSON payload:", err)
			}
			// Specify the URL of the POST endpoint
			url := "http://127.0.0.1:8085/endpoint"
			// Send the payload to the POST endpoint
			err = service.MakeRequest(jsonPayload, url)
			if err != nil {
				fmt.Println("Error sending payload:", err)
			}
			//fmt.Println("Payload sent successfully.")
			////////////////////////////////////////////////////////////////
			// Release the semaphore
			<-semaphore
		}()
	}
	// Wait for all goroutines to finish
	wg.Wait()
	elapsedTime := time.Since(startTime)
	requestsPerSecondActual := float64(numRequests) / elapsedTime.Seconds()
	fmt.Printf("Elapsed Time: %s\n", elapsedTime)
	fmt.Printf("Requests Per Second: %.2f\n", requestsPerSecondActual)
}
