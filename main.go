package main

import (
	"fmt"
	"os"
	"time"

	svg "github.com/ajstarks/svgo"

	"github.com/aguilam/git-activity/config"
	"github.com/aguilam/git-activity/data"
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

	datedCommits := map[string][]data.Commit{}
	for _, service := range services {
		result, err := service.GetCommits(from,to)
		if err != nil {
			println(err.Error())
			continue
		}

		for _, commit := range result {
			date := commit.Date.Format("2006-01-02")
			datedCommits[date] = append(datedCommits[date], commit)
		}
	}

	current := from
	weekDays := map[int]int{}

	err = os.MkdirAll("dist/images", 0755)
	if err != nil {
	    panic(err)
	}
	file, err := os.Create("dist/images/activity_750x150.svg")
	if err != nil {
	    panic(err)
	}
	defer file.Close()

	canvas := svg.New(file)
	canvas.Start(750,150)
	
	for current.Before(to){
		currentWeekDay := int(current.Weekday())
		date := current.Format("2006-01-02")
		commits := datedCommits[date]
		commitDayСount := weekDays[currentWeekDay]
		
		color := svgbuilder.GetCommitColor(len(commits))
		canvas.Rect(12 * commitDayСount,12 * currentWeekDay, 10, 10, fmt.Sprintf(`fill="%s" stroke="black" stroke-width="0.5" rx="2"`, color))

		weekDays[currentWeekDay]++
		current = current.AddDate(0,0,1)
	}
	canvas.End()

}