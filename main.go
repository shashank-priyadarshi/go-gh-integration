package main

import (
	"encoding/json"
	"fmt"
	"main/gh"
	"main/mongo"
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

	err = mongo.StartConnection(githubData)
	if err != nil {
		panic(err)
	}
}
