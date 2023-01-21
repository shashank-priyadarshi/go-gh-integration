package gh

type RepoResponse struct {
	Name  string `json:"name"`
	URL   string `json:"html_url"`
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
}

type CommitResponse struct {
	Commit struct {
		Committer struct {
			Name  string `json:"name"`
			Email string `json:"email"`
			Date  string `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}

type PullRequestResponse struct {
	State string `json:"state"`
	User  struct {
		Login string `json:"login"`
	} `json:"user"`
	CreatedAt string `json:"created_at"`
}

type IssueRequest struct {
	Title     string `json:"title"`
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
	ClosedAt  string `json:"closed_at"`
	Assignee  struct {
		Login string `json:"login"`
	} `json:"assignee"`
}

type GitHubData struct {
	Repo     Repo          `json:"repo"`
	WeekData []SCMActivity `json:"weekdata"`
}

type Repo struct {
	Count int        `json:"count"`
	List  []RepoList `json:"list"`
}

type RepoList struct {
	Name string `json:"name"`
}

type SCMActivity struct {
	PR           int      `json:"pr"`
	LOC          int      `json:"loc"`
	Date         string   `json:"date"`
	Commit       int      `json:"commit"`
	Issues       []Issues `json:"issues"`
	ClosedIssues int      `json:"closed_issues"`
}

type Issues struct {
	Title string `json:"title"`
}
