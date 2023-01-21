package gh

import (
	"os"
	"time"
)

var (
	authToken = os.Getenv("GITHUB_TOKEN")
	days      = int(time.Now().AddDate(0, 0, -7).Unix() / (24 * 60 * 60))
)
