package main

import (
	"encoding/json"
	"fmt"
	"main/gh"
)

// list of starred repos, with links
// list of active repos in last week, with links
// commit history last week
// pr history past week

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

	prettyJSON, err := (json.MarshalIndent(githubData, "", "  "))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(prettyJSON))
	fmt.Println("Data fetched successfully!")
}
