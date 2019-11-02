package main

import (
	"assignment2"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT") // auto assign port, needed for heroku support
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", assignment2.HandlerNil)
	http.HandleFunc("/repocheck/v1/commits/", assignment2.HandlerCommits)     // ?limit=4&auth=<access-token>
	http.HandleFunc("/repocheck/v1/languages/", assignment2.HandlerLanguages) // ?limit=4&auth=<access-token>
	http.HandleFunc("/repocheck/v1/issues/", assignment2.HandlerIssues)       // ?type=users|labels&auth=<access-token>
	http.HandleFunc("/repocheck/v1/status/", assignment2.HandlerStatus)
	http.HandleFunc("/repocheck/v1/webhooks/", assignment2.HandlerWebhooks)

	// print to console
	fmt.Println("Program started: ", assignment2.StartTime)
	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
