package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func api(w http.ResponseWriter, r *http.Request) {
	slackName := r.URL.Query().Get("slack_name")
	track := r.URL.Query().Get("track")

	currentTime := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	currentDay := time.Now().UTC().Weekday().String()

	// GitHub file URL and repo URL
	githubFileURL := "https://github.com/Pro-282/hngx_tasks/blob/main/hngstage1task/main.go"
	githubRepoURL := "https://github.com/Pro-282/hngx_tasks/tree/main/hngstage1task"

	response := map[string]interface{}{
		"slack_name":      slackName,
		"current_day":     currentDay,
		"utc_time":        currentTime,
		"track":           track,
		"github_file_url": githubFileURL,
		"github_repo_url": githubRepoURL,
		"status_code":     200,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleRequest() {
	http.HandleFunc("/api", api)
	address := "0.0.0.0:8080"
	log.Fatal(http.ListenAndServe(address, nil))
}

func main() {
	handleRequest()
}
