package assignment2

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
