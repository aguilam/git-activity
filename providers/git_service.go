package providers

import (
	"time"

	"github.com/aguilam/git-activity/data"
)

type GitService interface {
	GetCommits(from time.Time, to time.Time) ([]data.Commit,error)
}