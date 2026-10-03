package html

import (
	"html/template"
	"os"
	"time"

	"github.com/aguilam/git-activity/data"
)

func BuildHTML(commits map[string][]data.Commit) {
	err := os.MkdirAll("dist", 0755)
	if err != nil {
		panic(err)
	}

	distFile, err := os.Create("dist/index.html")
	if err != nil {
		panic(err)
	}
	defer distFile.Close()

	nowDate := time.Now()
	yearAgo := nowDate.AddDate(-1,0,0)
	weekDay := int(yearAgo.Weekday())
	if (weekDay > 0) {
		yearAgo = yearAgo.AddDate(0,0, -weekDay)
	}
	var activity Activities

	activity.CurrentActivity = diagramBuilder(yearAgo,nowDate,commits)

	from := time.Date(2026,time.January,1,0,0,0,0,time.Local)
	to := time.Date(2026,time.December,31,23,59,59,0,time.Local)
	lastActivity := diagramBuilder(from,to,commits)
	activity.YearActivities = append(activity.YearActivities,lastActivity)
	
	templ := template.Must(
		template.New("template.html").
			Funcs(template.FuncMap{
				"dayBackground": dayBackground,
			}).
			ParseFiles("html/template.html"),
	)
	err = templ.Execute(distFile,activity)

	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile("html/styles.css")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("dist/styles.css",file,0644)
	if err != nil {
		panic(err)
	}
}