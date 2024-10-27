package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	owner             = os.Getenv("TARGET_OWNER")
	repository        = os.Getenv("TARGET_REPOSITORY")
	GitHubAccessToken = "Bearer" + os.Getenv("GITHUB_ACCESS_TOKEN")
)

var now = time.Now().AddDate(0, 0, -7).Unix()

func Do() (string, error) {
	return "", nil
}

func GetCommits() (string, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.github.com/repos/%s/%s/commits?since=2024-10-20T00:00:00Z&until=2024-10-27T00:00:00Z", owner, repository),
		nil,
	)
	if err != nil {
		return "", err
	}

	// ref: https://docs.github.com/en/rest/commits/commits?apiVersion=2022-11-28#list-commits
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", GitHubAccessToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return "", nil
}

func GetPulls() (string, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.github.com/repos/%s/%s/pulls", owner, repository),
		nil,
	)
	if err != nil {
		return "", err
	}

	// ref: https://docs.github.com/en/rest/pulls/pulls?apiVersion=2022-11-28#list-pull-requests
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", GitHubAccessToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	var pull []pulls
	json.NewDecoder(resp.Body).Decode(&pull)

	return "", nil
}

func GetIssues() (string, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repository),
		nil,
	)
	if err != nil {
		return "", err
	}

	// ref: https://docs.github.com/en/rest/issues/issues?apiVersion=2022-11-28#list-repository-issues
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", GitHubAccessToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	return "", nil
}
