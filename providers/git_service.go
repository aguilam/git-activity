package providers

import (
	"time"

	"github.com/aguilam/git-activity/data"
)

type ServiceScheme struct {
	Username string
	Git GitService
}

type GitService interface {
	GetCommits(username string,from time.Time, to time.Time) ([]data.Commit,error)
}