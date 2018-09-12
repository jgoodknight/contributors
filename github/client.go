package github

import (
	"encoding/json"
	"log"
	"net/http"
)

type Client struct {
	Token string
}

func (c *Client) ListContributors(user_repo string) ([]Contributor, error) {

	repo_url := "https://api.github.com/repos/" + user_repo + "/contributors"
	client := &http.Client{}

	new_request, err := http.NewRequest("GET", repo_url, nil)
	if err != nil {
		return nil, err
	}
	new_request.Header.Set("Authorization", "token "+c.Token)
	new_response, err := client.Do(new_request)
	if err != nil {
		return nil, err
	}
	defer new_response.Body.Close()
	log.Println("API Response Status: ", new_response.Status)

	var contributor_slice []Contributor
	err = json.NewDecoder(new_response.Body).Decode(&contributor_slice)
	if err != nil {
		return nil, err
	}
	return contributor_slice, nil
}
