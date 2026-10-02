package html

import (
	"html/template"
	"os"

	"github.com/aguilam/git-activity/data"
)

func BuildHTML(commits map[string][]data.Commit) {
	err := os.MkdirAll("dist", 0755)
	if err != nil {
		panic(err)
	}
	templ, err := template.ParseFiles("html/template.html")
	if err != nil {
		panic(err)
	}

	distFile, err := os.Create("dist/index.html")
	if err != nil {
		panic(err)
	}
	defer distFile.Close()

	err = templ.Execute(distFile,commits)
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