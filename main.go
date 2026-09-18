package main

import (
	"time"

	"github.com/aguilam/git-activity/data"
	"github.com/aguilam/git-activity/providers"
)
func main() {
	services := []providers.ServiceScheme{
		{
			Username: "username",
			Git: &providers.GithubService{},
		},
	}

	commits := []data.Commit{}

	for _, service := range services {
		result, err := service.Git.GetCommits(service.Username,time.Date(2026,time.January,1,0,0,0,0,time.Local),time.Date(2026,time.December,30,23,59,59,0,time.Local))
		if err != nil {
			println(err.Error())
			continue
		}
		
		commits = append(commits, result...)
	}
	print(len(commits))
}