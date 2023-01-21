package main

import (
	"fmt"
	"main/gh"
)

// list of starred repos, with links
// list of active repos in last week, with links
// commit history last week
// pr history past week

func main() {
	fmt.Println("Plugin execution started...")
	gh.Main()
}
