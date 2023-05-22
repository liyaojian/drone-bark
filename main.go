package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	key := os.Getenv("PLUGIN_KEY")

	RepoName := os.Getenv("DRONE_REPO_NAME")
	//BuildStatus := os.Getenv("DRONE_BUILD_STATUS")
	BuildLink := os.Getenv("DRONE_BUILD_LINK")
	CommitMessage := os.Getenv("CI_COMMIT_MESSAGE")
	CommitMessage = strings.TrimSuffix(CommitMessage, "\n")
	CommitBranch := os.Getenv("CI_COMMIT_BRANCH")

	body := fmt.Sprintf("[Drone] 仓库:%s,分支:%s,修改:%s?url=%s&group=drone&level=timeSensitive", RepoName, CommitBranch, CommitMessage, BuildLink)
	reqUrl := fmt.Sprintf("%s/%s", key, body)

	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatalln("request error:", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("status code:", resp.StatusCode)
	}
}
