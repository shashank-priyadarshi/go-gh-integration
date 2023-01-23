package gh

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func fetchRepoWiseData() (int, []RepoList, []SCMActivity) {
	var repoListData []RepoResponse
	var commitList []CommitResponse
	var pullRequestList []PullRequestResponse

	repoList := []RepoList{}
	scmActivity := addDate()

	repoCount := 0
	rawRepoList := bearerAuthAPICall("https://api.github.com/user/repos?per_page=100&page=1&sort=pushed", authToken)
	err := json.Unmarshal(rawRepoList, &repoListData)
	if err != nil {
		log.Fatalln("Unable to unmarshal raw repo response: ", err)
	}

	for _, repo := range repoListData {
		if repo.Forked {
			continue
		}

		repoCount++
		repoList = append(repoList, RepoList{
			Name: repo.Name,
		})

		// api call to fetch commits in repo.Name for this iteration
		rawCommitResponse := bearerAuthAPICall(fmt.Sprintf("%v?per_page=100&page=1", strings.Split(repo.CommitsURL, "{/sha}")[0]), authToken)
		err = json.Unmarshal(rawCommitResponse, &commitList)
		if err != nil {
			log.Fatalln("Unable to unmarshal raw commit response: ", err)
		}

		// api call to fetch pull requests in repo.Name for this iteration
		rawPullRequestResponse := bearerAuthAPICall(fmt.Sprintf("%v?per_page=100&page=1", strings.Split(repo.PRURL, "{/number}")[0]), authToken)
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
		commitDate, _ := time.Parse("2006-01-02", commit.Commit.Committer.Date[:10])
		timeElapsed := 7 - (int(commitDate.Unix()/(24*60*60)) - days)
		if timeElapsed < 7 && timeElapsed >= 0 {
			scmActivity[timeElapsed].Commit++
		}
	}

	// according to date, increment pr data
	// recording pr count for corresponding duration
	for _, pr := range pullRequestList {
		prDate, _ := time.Parse("2006-01-02", pr.CreatedAt[:10])
		timeElapsed := 7 - (int(prDate.Unix()/(24*60*60)) - days)
		if timeElapsed < 7 && timeElapsed >= 0 {
			scmActivity[timeElapsed].PR++
		}
	}
	return scmActivity
}
