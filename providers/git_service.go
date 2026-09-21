package providers

import (
	"time"

	"github.com/aguilam/git-activity/data"
)

type GitServiceInfo struct {
	Type string
	Username string
	Url *string
}

type GitService interface {
	GetCommits(from time.Time, to time.Time) ([]data.Commit,error)
	GetServiceInfo() GitServiceInfo
}