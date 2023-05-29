package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

func main() {
	httpClient := &http.Client{
		// Use a custom transport that overrides the Dial function
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/tmp/unix.sock")
			},
		},
	}

	reqURL := &url.URL{
		Scheme: "http",
		Host:   "unix",
		Path:   "/",
	}

	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		os.Exit(1)
	}

	// Send the request
	res, err := httpClient.Do(req)
	if err != nil {
		log.Fatalf("Request error: %v", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	// Print the response status and headers
	fmt.Println("Status:", res.Status)
	fmt.Println("Headers:", res.Header)

	// Read and print the body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading body: %v", err)
	}
	fmt.Println("Body:", string(body))
}
