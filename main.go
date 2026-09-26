package main

import (
	"os"
	"sync"
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
	results := make([][]data.Commit,len(services))
	var wg sync.WaitGroup
	
	for i, service := range services {
		wg.Go(func(){
			result, err := service.GetCommits(yearAgo,to)
			if err != nil {
				println(err.Error())
				return
			}
			results[i] = result
		})
	}
	wg.Wait()
	for _, result := range results {
		for _, commit := range result {
			date := commit.Date.Format("2006-01-02")
			datedCommits[date] = append(datedCommits[date], commit)
		}
	}

	wg.Go(func(){
		svgbuilder.BuildSVG(from,to,datedCommits,services,"year_activity_750_150.svg")
	})
	wg.Go(func(){
		svgbuilder.BuildSVG(yearAgo,nowDate,datedCommits,services,"current_activity_750_150.svg")
	})
	wg.Go(func(){
		html.BuildHTML(datedCommits)
	})
	
	wg.Wait()
}