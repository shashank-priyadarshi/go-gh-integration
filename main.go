package main

import (
	"encoding/json"
	"fmt"
	"main/gh"
	"main/mongo"
)

// list of starred repos, with links
// commit history for past week
// loc history for past week
// pr history for past week
// closed issues count for past week
// open issues list for past week
// all data for user as well as org repos

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

	err = mongo.StartConnection(githubData)
	if err != nil {
		panic(err)
	}
}
