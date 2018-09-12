// This program will request information about all of the contributors to the Go project on GitHub
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Holds the data we actually care about from the Contributor request
type Contributor struct {
	Login         string
	Contributions int
}

func main() {
	github_url := "https://api.github.com/repos/golang/go/contributors"
	// mock_url := "http://127.0.0.1:54270/repos/golang/go/contributors"

	target_url := github_url

	github_token := os.Getenv("GITHUB_TOKEN")
	if github_token == "" {
		log.Print("Need Github Token!")
		os.Exit(1)
	}
	client := &http.Client{}

	new_request, err := http.NewRequest("GET", target_url, nil)
	if err != nil {
		log.Fatal(err)
	}
	new_request.Header.Set("Authorization", "token "+github_token)
	new_response, err := client.Do(new_request)
	if err != nil {
		log.Fatal(err)
	}
	defer new_response.Body.Close()
	log.Println("API Response Status: ", new_response.Status)

	var contributor_slice []Contributor
	err = json.NewDecoder(new_response.Body).Decode(&contributor_slice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(contributor_slice)
}
