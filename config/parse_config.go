package config

import (
	"github.com/aguilam/git-activity/providers"
	"go.yaml.in/yaml/v4"
)

func ParseConfig(file []byte) ([]providers.GitService) {
	gitServices := map[string] func(username string, token *string, url *string) providers.GitService{
		"github": func(username string, token *string, url *string) providers.GitService {
			return providers.GithubService{Username: username, Token: token, Url: url}
		},
	}

	var config Config
	err := yaml.Unmarshal(file,&config)
	if err != nil {
		panic(err)
	}
	
	var Services []providers.GitService
	for _, service := range config.Services{
		gitService := gitServices[service.Service]
		if gitService == nil {
			println("Unkown git service", service.Service)
			continue
		}
		Services = append(Services, gitService(service.Username,service.Token,service.Url))
	}

	return Services
}