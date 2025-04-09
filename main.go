package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Issue struct {
	Name      string   `json:"name"`
	Priority string   `json:"priority"`
	Assignees []string `json:"assignees"`
}

type IssueResponse struct {
	Results []Issue `json:"results"`
}

func main() {
	url := "https://api.plane.so/api/v1/workspaces/ikhsan-workspace/projects/bd1ed58e-c52f-4cc1-b2ab-e6440b369d37/issues/"
	apiKey := "plane_api_75401edcd20343fdb616563659adcea8"
	assigneeID := "424cada8-95ca-4710-8c93-66ce6e71c7b0"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("x-api-key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var response IssueResponse
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	for _, issue := range response.Results {
		for _, assignee := range issue.Assignees {
			if assignee == assigneeID {
				fmt.Printf("Issue: %s - Priority: %s\n", issue.Name, issue.Priority)
			}
		}
	}
}
