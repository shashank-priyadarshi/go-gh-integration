package gh

import (
	"encoding/json"
	"log"
	"time"
)

func getIssueData(scmActivity []SCMActivity) ([]SCMActivity, []Issue) {
	var issueList []IssueRequest
	var issues []Issue

	// api call to fetch list of all issues
	rawIssueResposne := bearerAuthAPICall("https://api.github.com/user/issues?per_page=100&page=1", authToken)
	err := json.Unmarshal(rawIssueResposne, &issueList)
	if err != nil {
		log.Fatalln("Unable to unmarshal raw issue list response: ", err)
	}
	// according to date, increment issue count
	// recording issue count & issue names for corresponding duration
	for _, issue := range issueList {
		var issueClosedDate time.Time

		issueCreatedDate, _ := time.Parse("2006-01-02", issue.CreatedAt[:10])

		if issue.ClosedAt != "" {
			issueClosedDate, _ = time.Parse("2006-01-02", issue.ClosedAt[:10])
			timeElapsed := int(issueClosedDate.Unix()/(24*60*60)) - days
			if timeElapsed < 7 && timeElapsed >= 0 {
				scmActivity[timeElapsed].ClosedIssues++
			}
			continue
		}

		issues = append(issues, Issue{
			Title: issue.Title,
		})

		timeElapsed := 7 - (int(issueCreatedDate.Unix()/(24*60*60)) - days)
		if timeElapsed < 7 && timeElapsed >= 0 {
			scmActivity[timeElapsed].Issues = append(scmActivity[timeElapsed].Issues, Issue{
				Title: issue.Title,
			})
		}

	}
	return scmActivity, issues
}
