package main

import (
	"encoding/json"
	"fmt"
	"main/gh"
	"main/mongo"
	"os"
)

// list of recent active repos with links
// loc history for past week
// pr history for past week

func main() {
	fmt.Println("Plugin execution started...")
	rawData, err := gh.Main()
	if err != nil {
		panic(err)
	}

	var githubData gh.GitHubData
	err = json.Unmarshal(rawData, &githubData)
	if err != nil {
		panic(err)
	}

	gitHubDataCollection, issueCollection := os.Getenv("GITHUB_DATA"), os.Getenv("ISSUE_DATA")

	err = mongo.StartConnection(gitHubDataCollection, githubData)
	if err != nil {
		panic(err)
	}
	err = mongo.StartConnection(issueCollection, struct {
		Issues []string `json:"issues"`
	}{Issues: githubData.Issues})
	if err != nil {
		panic(err)
	}
}
