package main

import (
	"assignment_2"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// StartTime contains the timestamp when the program started
var StartTime = time.Now()

func main() {

	port := os.Getenv("PORT") // auto assign port, needed for heroku support
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", assignment_2.HandlerNil)
	http.HandleFunc("/repocheck/v1/commits", assignment_2.HandlerCommits)
	http.HandleFunc("/repocheck/v1/languages", assignment_2.HandlerLanguages)
	http.HandleFunc("/repocheck/v1/issues", assignment_2.HandlerIssues)
	http.HandleFunc("/repocheck/v1/status", assignment_2.HandlerStatus)
	http.HandleFunc("/repocheck/v1/webhooks", assignment_2.HandlerWebhooks)

	// print to console
	fmt.Println("Program started: ", StartTime)
	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
