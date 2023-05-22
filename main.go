package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	key := os.Getenv("PLUGIN_KEY")

	RepoName := os.Getenv("DRONE_REPO_NAME")
	BuildStatus := os.Getenv("DRONE_BUILD_STATUS")

	body := fmt.Sprintf("%s:%s", RepoName, BuildStatus)
	reqUrl := fmt.Sprintf("%s/%s", key, body)

	resp, err := http.Get(reqUrl)
	if err != nil {
		log.Fatalln("request error:", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalln("status code:", resp.StatusCode)
	}
}
