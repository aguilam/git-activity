package main

import (
	"os"
	"time"

	"github.com/aguilam/git-activity/config"
	"github.com/aguilam/git-activity/data"
	"github.com/aguilam/git-activity/html"
	"github.com/aguilam/git-activity/svgbuilder"
)



func main() {
	configPath := os.Args[1]
	configBytes, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	services := config.ParseConfig(configBytes)

	from := time.Date(2026,time.January,1,0,0,0,0,time.Local)
	to := time.Date(2026,time.December,30,23,59,59,0,time.Local)
	nowDate := time.Now()

	yearAgo := nowDate.AddDate(-1,0,0)
	datedCommits := map[string][]data.Commit{}

	for _, service := range services {
		result, err := service.GetCommits(yearAgo,to)
		if err != nil {
			println(err.Error())
			continue
		}

		for _, commit := range result {
			date := commit.Date.Format("2006-01-02")
			datedCommits[date] = append(datedCommits[date], commit)
		}
	}

	svgbuilder.BuildSVG(from,to,datedCommits,services,"year_activity_750_150.svg")
	svgbuilder.BuildSVG(yearAgo,nowDate,datedCommits,services,"current_activity_750_150.svg")

	html.BuildHTML(datedCommits)
}