package svgbuilder

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aguilam/git-activity/data"
	"github.com/aguilam/git-activity/providers"
	svg "github.com/ajstarks/svgo"
)

func addPattern(canvas *svg.SVG, id string, colors []string) {
	canvas.Def()
	width := len(colors) * 3

	canvas.Pattern(
		id,
		0, 0,
		width, width,
		"user",
		`patternTransform="rotate(45)"`,
	)

	for i, color := range colors {
		canvas.Rect(
			i*4,
			0,
			4,
			width,
			fmt.Sprintf(`fill="%s"`, color),
		)
	}

	canvas.PatternEnd()
	canvas.DefEnd()
}

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
		commitDayCount := weekDays[currentWeekDay]
		
		colors := GetCommitColor(commits)
		if len(colors) > 1 {
			patternID := fmt.Sprintf("stripe-%s", strings.ReplaceAll(strings.Join(colors, "-"),"#","-"))

			addPattern(canvas, patternID, colors)
	
			canvas.Rect(
				12*commitDayCount,
				12*currentWeekDay,
				10,
				10,
				fmt.Sprintf(
					`fill="url(#%s)" stroke="black" stroke-width="0.5" rx="2"`,
					patternID,
				),
			)
		} else {
			canvas.Rect(12 * commitDayCount,12 * currentWeekDay, 10, 10, fmt.Sprintf(`fill="%s" stroke="black" stroke-width="0.5" rx="2"`, colors[0]))
		}

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