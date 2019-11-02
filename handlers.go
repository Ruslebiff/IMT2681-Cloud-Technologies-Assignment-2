package assignment2

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// HandlerNil is the default http handler
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default Handler: Invalid request received.") // error to console
	http.Error(w, "Invalid request", http.StatusBadRequest)   // error to http
}

// HandlerCommits returns the repos with highest numbers of commits, ?limit parameter
func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

// HandlerLanguages returns the languages used in given projects by dist. in descending order
func HandlerLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

// HandlerIssues returns the name of the users or labels (see parameters) for a given project
func HandlerIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

// HandlerStatus returns information about availability of invoked service and database connectivity
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "HandlerStatus is not finished yet.") // TO BE REMOVED
	var S = &Status{}

	gitlabget, err := http.Get(GitLabRoot)
	if err != nil {
		http.Error(w, "API lookup failed", http.StatusServiceUnavailable)
		fmt.Println(http.StatusServiceUnavailable)
	}
	// close connection, prevent resource leak if get-request fails
	defer gitlabget.Body.Close()

	S.GitLab = gitlabget.StatusCode
	S.Version = "V1"
	elapsed := time.Since(StartTime)
	S.Uptime = elapsed.Seconds()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(S)
}

// HandlerWebhooks handles webhooks
func HandlerWebhooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}
