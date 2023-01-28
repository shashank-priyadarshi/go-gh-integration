package gh

import (
	"encoding/json"
	"fmt"
	"time"
)

func Main() ([]byte, error) {
	fmt.Println("Inside Main function")

	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println(time.LoadLocation("Asia/Calcutta"))
	}
	fmt.Println("Location: ", loc)
	fmt.Println("Time: ", time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)))

	repoCount, repoList, scmActivity := fetchRepoWiseData()
	scmActivity, issues := getIssueData(scmActivity)

	gitHubData := GitHubData{
		Repos: Repo{
			List:  repoList,
			Count: repoCount,
		},
		WeekData:     scmActivity,
		Issues:       issues,
		StarredRepos: fetchStarredRepos(),
		Time:         time.Now().In(time.FixedZone("Asia/Kolkata", 5*60*60+30*60)),
	}
	fmt.Println("Plugin execution completed!")
	return json.Marshal(gitHubData)
}

func addDate() []SCMActivity {
	scmActivity := []SCMActivity{}

	num := 0
	for num < 7 {
		scmActivity = append(scmActivity, SCMActivity{
			Date: time.Now().AddDate(0, 0, -num).String(),
		})
		num++
	}

	return scmActivity
}
