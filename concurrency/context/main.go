package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	timeoutContext, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	url := "https://placeholder.com/wp-content/uploads/2018/10/placeholder.com-logo1.png"

	// Create HTTP request with timeout
	req, err := http.NewRequestWithContext(timeoutContext, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}

	// Perform HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Get data from HTTP response
	imageData, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Image size is %d bytes\n", len(imageData))
}
