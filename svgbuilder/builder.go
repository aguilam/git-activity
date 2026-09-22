package svgbuilder

import (
	"fmt"
	"os"
	"time"

	"github.com/aguilam/git-activity/data"
	"github.com/aguilam/git-activity/providers"
	svg "github.com/ajstarks/svgo"
)


func BuildSVG (from time.Time, to time.Time, datedCommits map[string][]data.Commit, services []providers.GitService, filename string) {
	current := from
	weekDays := map[int]int{}

	err := os.MkdirAll("dist/images", 0755)
	if err != nil {
	    panic(err)
	}
	file, err := os.Create(fmt.Sprintf("dist/images/%s",filename))
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
		
		color := GetCommitColor(len(commits))
		canvas.Rect(12 * commitDayСount,12 * currentWeekDay, 10, 10, fmt.Sprintf(`fill="%s" stroke="black" stroke-width="0.5" rx="2"`, color))

		weekDays[currentWeekDay]++
		current = current.AddDate(0,0,1)
	}
	canvas.Text(0,95,"Based on Git services: ",`fill="white"`)
	var finalX int
	for i, service := range services {
		currentX := 13 * i + finalX 
		serviceInfo := service.GetServiceInfo()
		text := fmt.Sprintf("%s %s",serviceInfo.Type, *serviceInfo.Url)
		canvas.Text(currentX,110,text,`fill="white"`)
		finalX = currentX + len(text) * 8
	}
	canvas.End()
}