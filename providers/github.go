package providers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/aguilam/git-activity/data"
)

type GithubService struct {
	Username string
	Token *string
	Url *string
}

func (g GithubService) GetCommits(from time.Time, to time.Time) ([]data.Commit, error) {
	var result struct {
		Items []struct {
			SHA string `json:"sha"`
			Commit struct {
				Url string `json:"url"`
				Message string `json:"message"`
				Author struct {
					Date string `json:"date"`
					Name string `json:"name"`
				} `json:"author"`
			} `json:"commit"`
			Repository struct {
				Name string `json:"name"`
			} `json:"repository"`
		} `json:"items"`
	}
	
	requestString := "https://api.github.com/search/commits?q=author:%s+author-date:%s..%s&per_page=100&page=%d"
	var commits []data.Commit

	for i := 1; ;i++ {
		resp, err := http.Get(fmt.Sprintf(requestString,g.Username,from.Format("2006-01-02"),to.Format("2006-01-02"),i))
		if err != nil {
			println(err.Error())
			return nil,err
		}

		err = json.NewDecoder(resp.Body).Decode(&result)

		resp.Body.Close()
		for _, item := range result.Items {
			date, err := time.Parse(time.RFC3339,item.Commit.Author.Date)
			if err != nil {
				println(err.Error())
				continue
			}
			commits = append(commits, data.Commit{GitService: "github",Repository: item.Repository.Name, SHA: item.SHA, Date: date,Message: item.Commit.Message, CommitUrl: item.Commit.Url})
		}

		link := resp.Header.Get("Link")
		if !strings.Contains(link, `rel="next"`) {
			break
		}
	}

	return commits, nil
}

func (g GithubService) 	GetServiceInfo() GitServiceInfo {
	url := fmt.Sprintf("https://github.com/%s",g.Username)
	return GitServiceInfo{Type: "github",Username: g.Username, Url: &url}
}