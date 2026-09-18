package data

import (
	"time"
)
type Commit struct {
	GitService string
	Repository string
	SHA        string
	Date       time.Time
	Message    string
	CommitUrl string
}