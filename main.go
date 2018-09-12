package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jgoodknight/contributors/github"
)

// This program will request information about all of the contributors to the Go project on GitHub

func main() {

	github_token := os.Getenv("GITHUB_TOKEN")
	target_user_repo := "golang/go"
	if github_token == "" {
		log.Print("Need Github Token stored in GITHUB_TOKEN environment variable!")
		os.Exit(1)
	}
	my_client := github.Client{github_token}
	contributor_slice, err := my_client.ListContributors(target_user_repo)
	if err != nil {
		log.Fatal(err)
	}
	for i := range contributor_slice {
		contributor := &contributor_slice[i]
		fmt.Printf("%s | %d \n", contributor.Login, contributor.Contributions)
	}

}
