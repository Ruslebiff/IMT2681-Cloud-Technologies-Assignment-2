package assignment2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"time"
)

// HandlerNil is the default http handler
func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default Handler: Invalid request received.") // error to console
	http.Error(w, "Invalid request", http.StatusBadRequest)   // error to http
}

// HandlerCommits returns the repos with highest numbers of commits, ?limit parameter
func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	var c = &Commit{}

	var repos []Repo
	var results []Result

	limit := r.URL.Query().Get("limit")
	auth := r.URL.Query().Get("auth")

	if limit == "" { // if no limit in url
		limit = "5" // set default limit
	}

	if auth != "" {
		c.Auth = true
	} else {
		c.Auth = false
	}

	url1 := GitLabRoot + "v4/projects/"
	url2 := GitLabRoot + "v4/projects?per_page=1"

	resp, err := http.Get(url2 + "&private_token=" + auth)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	totalpages, err := strconv.Atoi(resp.Header.Get("X-Total-Pages"))
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}

	//	totalRepositories, err := strconv.Atoi(resp.Header.Get("X-Total"))
	//	if err != nil {
	//		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	//	}

	for i := 1; i <= totalpages; i++ {
		var temp []Repo

		pagenumber := strconv.Itoa(i)
		resp, err := http.Get(url2 + "&page=" + pagenumber + "&private_token=" + auth)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		jsonresponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(http.StatusInternalServerError)
			return
		}

		json.Unmarshal([]byte(jsonresponse), &temp)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		repos = append(repos, temp...) // add each page to repos array
	}

	for _, j := range repos {
		resp, err := http.Get(url1 + strconv.Itoa(j.ID) + "/repository/commits?private_token=" + auth)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(http.StatusInternalServerError)
			return
		}

		totalcommits, err := strconv.Atoi(resp.Header.Get("X-total"))
		if err != nil {
			totalcommits = 0
		}

		resultsappend := Result{Repository: j.Reponame, Commits: totalcommits}
		results = append(results, resultsappend)
	}

	sort.Slice(results, func(i, k int) bool { // Sorting by # of commits
		return results[i].Commits > results[k].Commits
	})

	limitint, err := strconv.Atoi(limit)
	if err != nil {
		return
	}

	for i := 1; i <= limitint; i++ {
		c.Repos = append(c.Repos, results[i])
	}

	http.Header.Add(w.Header(), "content-type", "application/json") // json header
	json.NewEncoder(w).Encode(c)                                    // Encode struct to JSON
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

	//fmt.Fprintln(w, "HandlerStatus is not finished yet.") // TO BE REMOVED
	var s = &Status{}

	// GET-request to gGitLab API
	gitlabget, err := http.Get(GitLabRoot)
	if err != nil {
		http.Error(w, "API lookup failed", http.StatusServiceUnavailable)
		fmt.Println(http.StatusServiceUnavailable)
	}
	/*
		dbget, err := http.Get(DatabaseRoot)
		if err != nil {
			http.Error(w, "Database lookup failed", http.StatusServiceUnavailable)
			fmt.Println(http.StatusServiceUnavailable)

		}
	*/
	// Close connection, prevent resource leak if get-request fails
	defer gitlabget.Body.Close()

	// Assign values to struct
	s.GitLab = gitlabget.StatusCode
	//s.Database = dbget.StatusCode
	s.Database = 501 // dummy, to be removed when database is implemented
	s.Version = "V1"
	elapsed := time.Since(StartTime)
	s.Uptime = elapsed.Seconds()

	http.Header.Add(w.Header(), "content-type", "application/json")
	json.NewEncoder(w).Encode(s) // Encode struct to JSON
}

// HandlerWebhooks handles webhooks
func HandlerWebhooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}
