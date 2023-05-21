package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
)

var mutex sync.Mutex

func MakeRequest(jsonPayload []byte, url string) error {
	mutex.Lock()
	defer mutex.Unlock()
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
	_, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {

	}
	//fmt.Println(string(re))
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
