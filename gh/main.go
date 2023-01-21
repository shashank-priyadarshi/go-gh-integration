package gh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

func Main() {
	fmt.Println("Inside Main function")

	repoCount, repoList, scmActivity := fetchRepoWiseData()
	scmActivity = getIssueData(scmActivity)

	gitHubData := GitHubData{
		Repo: Repo{
			List:  repoList,
			Count: repoCount,
		},
		WeekData: scmActivity,
	}
	fmt.Println(gitHubData)
	fmt.Println("Plugin executed successfully!")
}

func fetchRepoWiseData() (int, []RepoList, []SCMActivity) {
	var repoListData []RepoResponse
	var commitList []CommitResponse
	var pullRequestList []PullRequestResponse

	repoList := []RepoList{}
	scmActivity := addDate()

	repoCount := 0
	rawRepoList := bearerAuthAPICall("https://api.github.com/user/repos?per_page=100&page=1&type=private&sort=pushed", authToken)
	err := json.Unmarshal(rawRepoList, &repoListData)
	if err != nil {
		log.Fatalln("Unable to unmarshal raw repo response: ", err)
	}

	for _, repo := range repoListData {
		if !strings.EqualFold(repo.Owner.Login, "shashank-priyadarshi") {
			continue
		}

		repoCount++
		repoList = append(repoList, RepoList{
			Name: repo.Name,
		})

		// api call to fetch commits in repo.Name for this iteration
		rawCommitResponse := bearerAuthAPICall(fmt.Sprintf("https://api.github.com/repos/shashank-priyadarshi/%v/commits?per_page=100&page=1", repo.Name), authToken)
		err = json.Unmarshal(rawCommitResponse, &commitList)
		if err != nil {
			log.Fatalln("Unable to unmarshal raw commit response: ", err)
		}

		// api call to fetch pull requests in repo.Name for this iteration
		rawPullRequestResponse := bearerAuthAPICall(fmt.Sprintf("https://api.github.com/repos/shashank-priyadarshi/%v/pulls?per_page=100&page=1&state=all", repo.Name), authToken)
		err = json.Unmarshal(rawPullRequestResponse, &pullRequestList)
		if err != nil {
			log.Fatalln("Unable to unmarshal raw pull request response: ", err)
		}

		scmActivity = appendRepoWiseData(scmActivity, commitList, pullRequestList)
	}
	return repoCount, repoList, scmActivity
}

func appendRepoWiseData(scmActivity []SCMActivity, commitList []CommitResponse, pullRequestList []PullRequestResponse) []SCMActivity {
	// according to date, increment commit data
	// recording commit count for corresponding duration
	for _, commit := range commitList {
		commitDate, _ := time.Parse(time.RFC3339, commit.Commit.Committer.Date)
		timeElapsed := int(commitDate.Unix()/(24*60*60)) - days
		if timeElapsed < 7 && timeElapsed >= 0 {
			scmActivity[timeElapsed].Commit++
		}
	}

	// according to date, increment pr data
	// recording pr count for corresponding duration
	for _, pr := range pullRequestList {
		prDate, _ := time.Parse(time.RFC3339, pr.CreatedAt)
		timeElapsed := int(prDate.Unix()/(24*60*60)) - days
		if timeElapsed < 7 && timeElapsed >= 0 {
			scmActivity[timeElapsed].PR++
		}
	}
	return scmActivity
}

func getIssueData(scmActivity []SCMActivity) []SCMActivity {
	var issueList []IssueRequest

	// api call to fetch list of all issues
	rawIssueResposne := bearerAuthAPICall("https://api.github.com/user/issues?per_page=100&page=1", authToken)
	err := json.Unmarshal(rawIssueResposne, &issueList)
	if err != nil {
		log.Fatalln("Unable to unmarshal raw issue list response: ", err)
	}
	// according to date, increment issue count
	// recording issue count & issue names for corresponding duration
	for _, issue := range issueList {
		issueClosedDate, _ := time.Parse(time.RFC3339, issue.ClosedAt)
		issueCreatedDate, _ := time.Parse(time.RFC3339, issue.CreatedAt)

		if strings.EqualFold(issue.State, "Open") {
			timeElapsed := int(issueCreatedDate.Unix()/(24*60*60)) - days
			if timeElapsed < 7 && timeElapsed >= 0 {
				scmActivity[timeElapsed].Issues = append(scmActivity[timeElapsed].Issues, Issues{
					Title: issue.Title,
				})
			}
			continue
		}

		timeElapsed := int(issueClosedDate.Unix()/(24*60*60)) - days
		if timeElapsed < 7 && timeElapsed >= 0 {
			scmActivity[timeElapsed].ClosedIssues++
		}
	}
	return scmActivity
}

func addDate() []SCMActivity {
	scmActivity := []SCMActivity{}

	num := 0
	for num < 7 {
		scmActivity = append(scmActivity, SCMActivity{
			Date: time.Now().AddDate(0, 0, num).String(),
		})
		num++
	}

	return scmActivity
}

func bearerAuthAPICall(reqURL, authToken string) []byte {
	timeOut := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeOut,
	}

	request, err := http.NewRequest("GET", reqURL, bytes.NewBuffer([]byte("")))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %v", authToken))
	if err != nil {
		log.Fatalln("err in creating new request: ", err)
	}

	resp, err := client.Do(request)

	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("error '%v' while making request to %v: %v", err, reqURL, resp.StatusCode))
	}

	if err != nil {
		log.Fatalln("err in making bearerAuth req: ", err)
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln("err in reading req response: ", err)
	}

	return respBody
}
