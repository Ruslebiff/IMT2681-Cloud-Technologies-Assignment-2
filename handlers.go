package assignment2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
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

	for i := 0; i < limitint; i++ {
		c.Repos = append(c.Repos, results[i])
	}

	http.Header.Add(w.Header(), "content-type", "application/json") // json header
	json.NewEncoder(w).Encode(c)                                    // Encode struct to JSON
}

// HandlerLanguages returns the languages used in given projects by distribution in descending order
func HandlerLanguages(w http.ResponseWriter, r *http.Request) {
	var p = &Projectinfo{}     // languages[] + auth bool
	var repos []Repo           // temp storage for all repos
	var repolanguages []string // list of all languages used in each project

	limit := r.URL.Query().Get("limit")
	auth := r.URL.Query().Get("auth")

	if limit == "" { // if no limit in url
		limit = "5" // set default limit
	}

	if auth != "" {
		p.Auth = true
	} else {
		p.Auth = false
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
		langmap := make(map[string]float64) // current repo's languages and percentage value
		resp, err := http.Get(url1 + strconv.Itoa(j.ID) + "/languages?private_token=" + auth)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			fmt.Println(http.StatusInternalServerError)
			return
		}

		if resp.StatusCode == http.StatusOK { // if repo has /languages/
			jsonresponse, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				fmt.Println(http.StatusInternalServerError)
				return
			}

			json.Unmarshal([]byte(jsonresponse), &langmap)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			for lang := range langmap { // add langmap for this repo to repolanguages
				repolanguages = append(repolanguages, lang)
			}
		}
	}

	langcounter := countduplicates(repolanguages) // Languages and their counted number
	var langcounterarray []LangCount              // Turn langcounter into an array
	for l, c := range langcounter {
		langcounterarray = append(langcounterarray, LangCount{l, c})
	}

	sort.Slice(langcounterarray, func(i, j int) bool { // sort languages by counted value
		return langcounterarray[i].Count > langcounterarray[j].Count
	})

	limitint, err := strconv.Atoi(limit) // limit from url query to int
	if err != nil {
		return
	}

	if limitint > len(langcounterarray) { // make sure limit doesn't extend number of languages
		limitint = len(langcounterarray)
	}
	for i := 0; i < limitint; i++ {
		p.Languages = append(p.Languages, langcounterarray[i].LanguageName)
	}
	http.Header.Add(w.Header(), "content-type", "application/json") // json header
	json.NewEncoder(w).Encode(p)                                    // Encode struct to JSON
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

	dbget, err := http.Get(DatabaseRoot)
	if err != nil {
		http.Error(w, "Database lookup failed", http.StatusServiceUnavailable)
		fmt.Println(http.StatusServiceUnavailable)

	}

	// Close connection, prevent resource leak if get-request fails
	defer gitlabget.Body.Close()

	// Assign values to struct
	s.GitLab = gitlabget.StatusCode
	s.Database = dbget.StatusCode
	//s.Database = 501 // dummy, to be removed when database is implemented
	s.Version = "V1"
	elapsed := time.Since(StartTime)
	s.Uptime = elapsed.Seconds()

	http.Header.Add(w.Header(), "content-type", "application/json")
	json.NewEncoder(w).Encode(s) // Encode struct to JSON
}

// HandlerWebhooks handles webhooks
func HandlerWebhooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		webhook := Webhookreg{}
		err := json.NewDecoder(r.Body).Decode(&webhook)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		webhook.Time = time.Now()

		_, err = DBSave(&webhook)
		if err != nil {
			http.Error(w, "Error occurred when saving", http.StatusInternalServerError)
			return
		}

		fmt.Println("Webhook " + webhook.URL + " registered")

	case http.MethodGet:
		parts := strings.Split(r.URL.String(), "/")
		if parts[4] != "" {
			// print only the specified webhook by id

			webhookid := parts[4]               // Find webhook id
			webhook, err := DBReadid(webhookid) // Read the webhook from database
			if err != nil {
				http.Error(w, "Error: Not found!", http.StatusBadRequest)
				return
			}
			http.Header.Add(w.Header(), "content-type", "application/json") // json header
			json.NewEncoder(w).Encode(webhook)                              // Encode it
		} else {
			// create list of all webhooks
			var webhooks []Webhookreg    // temp storage for all webhooks from db
			webhooks, err := DBReadall() // read from db into webhooks array
			if err != nil {
				http.Error(w, "Error: Failed to read webhooks from DB", http.StatusInternalServerError)
				return
			}
			http.Header.Add(w.Header(), "content-type", "application/json") // json header
			json.NewEncoder(w).Encode(webhooks)                             // Encode them all

		}

	case http.MethodDelete:
		parts := strings.Split(r.URL.String(), "/")
		if parts[4] != "" {
			webhookid := parts[4]      // find id of webhook
			err := DBDelete(webhookid) // delete webhook with this id from database
			if err != nil {
				fmt.Println("Failed to delete webhook with id " + webhookid)
			}

			fmt.Println("Webhook with id" + webhookid + " deleted from database.")

		} else {
			fmt.Fprintln(w, "No webhook ID was specified", http.StatusBadRequest)
		}

	default:
		fmt.Println("DEFAULT used")
		http.Error(w, "invalid http method used", http.StatusBadRequest)
	}
}
