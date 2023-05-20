package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
)

func MakeRequest(jsonPayload []byte, url string) error {
	//_, err := http.Get(url) // Replace with your endpoint URL
	//if err != nil {
	//	return err
	//}
	////responseBody, err := ioutil.ReadAll(resp.Body)
	////if err != nil {
	////	fmt.Println("Error reading response body:", err)
	////	return err
	////}
	////
	////// Print the response status code and body
	////fmt.Println("Response Status:", resp.Status)
	////defer resp.Body.Close()
	////fmt.Println("Response Body:", string(responseBody))
	//// Do something with the response if needed
	//return nil

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	re, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(re))
	return nil

}
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
func GenerateJSONPayload(sizeInBytes int) ([]byte, error) {
	payload := Payload{
		Data: GenerateRandomString(sizeInBytes),
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	return jsonPayload, nil
}
