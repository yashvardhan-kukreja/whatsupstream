package status

import (
	"fmt"
	"whatsupstream/pkg/apis/config"
)

func getReposFromConfigFile(configPath string) ([]string, error) {
	repos := []string{}

	configFile, err := config.YamlConfigToInputConfig(configPath)
	if err != nil {
		return nil, err
	}

	for _, issueConfig := range configFile.IssueConfigs {
		repo := fmt.Sprintf(issueConfig.RepositoryURL)
		for _, issueLabel := range issueConfig.Labels {
			repo = fmt.Sprintf("%s [%s]", repo, issueLabel)
		}
		repos = append(repos, repo)
	}
	return repos, nil
}
