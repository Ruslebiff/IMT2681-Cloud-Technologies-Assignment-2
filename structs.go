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

// Webhookreg struct
type Webhookreg struct {
	ID    string
	Event string `json:"event"`
	URL   string `json:"url"`
	Time  time.Time
}

// WebhookPayload struct
type WebhookPayload struct {
	Event      string
	Parameters []string
	Time       time.Time
}

// Firebase struct
type Firebase struct {
	Ctx    context.Context
	Client *firestore.Client
}
