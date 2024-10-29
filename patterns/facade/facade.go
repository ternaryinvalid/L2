package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClinet() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

func (c *HttpClient) Get(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *HttpClient) Post(url string, data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	resp, err := c.client.Post(url, "application/json", bytes.NewBuffer(jsonData))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

type ErrorLogger struct{}

func NewErrorLogger() *ErrorLogger {
	return &ErrorLogger{}
}

func (e *ErrorLogger) Log(err error) {
	if err != nil {
		log.Println("Error", err)
	}
}

type ApiFacade struct {
	client      *HttpClient
	errorlogger *ErrorLogger
}

func NewFacade() *ApiFacade {
	return &ApiFacade{
		client:      NewHttpClinet(),
		errorlogger: NewErrorLogger(),
	}
}

func (f *ApiFacade) FetchData(url string) (string, error) {
	data, err := f.client.Get(url)

	if err != nil {
		f.errorlogger.Log(err)
		return "", err
	}

	return string(data), err
}

func (f *ApiFacade) PostData(url string, payload interface{}) (string, error) {
	data, err := f.client.Post(url, payload)

	if err != nil {
		f.errorlogger.Log(err)
		return "", err
	}

	return string(data), nil
}

func main() {
	api := NewFacade()
	url := "https://jsonplaceholder.typicode.com/posts/1"
	data, err := api.FetchData(url)

	if err != nil {
		return
	}

	fmt.Println("Fetched data: ", data)

	payload := map[string]string{"title": "title1", "body": "body1", "userId": "1"}
	response, err := api.PostData("https://jsonplaceholder.typicode.com/posts", payload)

	if err != nil {
		return
	}

	fmt.Println("Post Response:", response)
}
