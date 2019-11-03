package assignment2

import "time"

// GitLabRoot = root path for GitLab API
const GitLabRoot = "https://git.gvk.idi.ntnu.no/api/"

// DatabaseRoot = rooth path for the database
const DatabaseRoot = "https://console.firebase.google.com/project/assignment2-2c6b0/database"

// Firebasecredentials is the json file with your access token to a firebase database
// Edit this your own file path/name
const Firebasecredentials = "./assignment2-2c6b0-firebase-adminsdk-9dvth-77d8aa990f.json"

// StartTime contains the timestamp when the program started
var StartTime = time.Now()
