package providers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/aguilam/git-activity/data"
)

type GitlabService struct {
	Username string
	Token *string
	Url *string
}

func (g GitlabService) GetCommits(from time.Time, to time.Time) ([]data.Commit, error) {
	resp, err := http.Get(fmt.Sprintf("https://gitlab.com/api/v4/users/%s/contributed_projects?per_page=100&page=1",g.Username))
	if err != nil {
		return nil, err
	}
	
	var projects []struct {
		Id int  `json:"id"`
		Name string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&projects)
	resp.Body.Close()

	if err != nil {
		return nil, err
	}

	var gitlabCommits []struct {
		Id        string `json:"id"`
		Date       time.Time `json:"authored_date"`
		Message    string `json:"message"`
		CommitUrl  string `json:"web_url"`
	}

	var commits []data.Commit
	getCommitsString := "https://gitlab.com/api/v4/projects/%d/repository/commits?author=%s&since=%s&until=%s&per_page=70&page=%d"
	for _, project := range projects {
		for i := 1; ;i++ {
			resp, err := http.Get(fmt.Sprintf(getCommitsString,project.Id,g.Username,from.UTC().Format(time.RFC3339),to.UTC().Format(time.RFC3339),i))
			if err != nil {
				println(err.Error())
				continue
			}
			if resp.StatusCode < 200 || resp.StatusCode >= 300 {
				body, _ := io.ReadAll(resp.Body)
				resp.Body.Close()
			
				return nil, fmt.Errorf(
					"gitlab: status=%d body=%s",
					resp.StatusCode,
					body,
				)
			}
			err = json.NewDecoder(resp.Body).Decode(&gitlabCommits)
			resp.Body.Close()
			
			if err != nil {
				println(err.Error())
				continue
			}

			if len(gitlabCommits) == 0 {
				break
			}

			for _, commit := range gitlabCommits {
				commits = append(commits, data.Commit{
					GitService: "gitlab",
					Repository: project.Name,
					SHA:       commit.Id,
					Date:      commit.Date,
					Message:   commit.Message,
					CommitUrl: commit.CommitUrl,
				})
			}
		}
	}
	return commits, nil
}