package svgbuilder

import "github.com/aguilam/git-activity/data"


func GetGitlabColors(commitsCount int) string {
	switch {
		case commitsCount >= 4:
			return "#303470"
		case commitsCount == 3:
			return "#4e65cd"
		case commitsCount == 2:
			return "#7992f5"
		case commitsCount == 1:
			return "#d2dcff"
		default:
			return "#151b23"
	}
}

func GetGithubColors(commitsCount int) string {
	switch {
		case commitsCount >= 4:
			return "#56d364"
		case commitsCount == 3:
			return "#2ea043"
		case commitsCount == 2:
			return "#196c2e"
		case commitsCount == 1:
			return "#033a16"
		default:
			return "#151b23"
	}
}

func GetCommitColor(commits []data.Commit) []string {
	if len(commits) == 0 {
		return []string{"#151b23"}
	}
	var colors []string
	services := map[string]int{}
	for _,commit := range commits {
		services[commit.GitService]++
	}
	for name , count := range services {
		switch {
			case name == "gitlab":
				colors = append(colors, GetGitlabColors(count))
			case name == "github":
				colors = append(colors, GetGithubColors(count))
		}
	}
	return colors
}

