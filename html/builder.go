package html

import (
	"os"

	"github.com/aguilam/git-activity/data"
)

func BuildHTML(commits map[string][]data.Commit) {
	err := os.MkdirAll("dist", 0755)
	if err != nil {
		panic(err)
	}
	file, err := os.ReadFile("html/template.html")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("dist/index.html",file,0644)
	if err != nil {
		panic(err)
	}

	file, err = os.ReadFile("html/styles.css")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("dist/styles.css",file,0644)
	if err != nil {
		panic(err)
	}
}