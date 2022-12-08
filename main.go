package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func main() {
	request_response := make(chan string)

	go createRequest(request_response)
	// go createRequest2()

	fmt.Println("İşlem Bitti")
	fmt.Println(<-request_response)
}

func createRequest(channel chan string) {
	fmt.Println("İşlem İçerde")
	resp, err := http.Get("https://api.uncomsys.com/api/v1/login")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	response_content := ErrorResponse{}

	json.Unmarshal(body, &response_content)
	fmt.Printf("Message: %s\n", response_content.Message)
	fmt.Printf("Operation: %s\n", strconv.FormatBool(response_content.Success))

	channel <- string(body)
}

func createRequest2() {
	fmt.Println("İşlem İçerde")
	resp, err := http.Get("https://api.uncomsys.com/api/v1/login")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	response_content := ErrorResponse{}

	json.Unmarshal(body, &response_content)
	fmt.Printf("Message: %s\n", response_content.Message)
	fmt.Printf("Operation: %s\n", strconv.FormatBool(response_content.Success))
}
