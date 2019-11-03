package assignment2

import (
	"time"

	"cloud.google.com/go/firestore"
	"golang.org/x/net/context"
)

// Status struct for /status/ endpoint
type Status struct {
	GitLab   int
	Database int
	Uptime   float64
	Version  string
}

// Result struct
type Result struct {
	Repository string
	Commits    int
}

// Repo struct for json
type Repo struct {
	ID       int    `json:"id"`
	Reponame string `json:"path_with_namespace"`
}

// Commit struct
type Commit struct {
	Repos []Result
	Auth  bool
}

// Projectinfo struct for info about specific project
type Projectinfo struct {
	Languages []string
	Auth      bool
}

// LangCount struct for name of language and its counted number of occurrences in projects
type LangCount struct {
	LanguageName string
	Count        int
}

// Webhookreg struct for registering a webhook and saving it to database
type Webhookreg struct {
	ID    string
	Event string `json:"event"`
	URL   string `json:"url"`
	Time  time.Time
}

// WebhookPayload struct for whats sent when we invoke a webhook
type WebhookPayload struct {
	Event      string `json:"event"`
	Parameters string `json:"parameters"`
	Time       string `json:"time"`
}

// Firebase struct for database
type Firebase struct {
	Ctx    context.Context
	Client *firestore.Client
}

// Issuerepo struct for /issues/ handler
type Issuerepo struct {
	Projectname string `json:"project"`
}
