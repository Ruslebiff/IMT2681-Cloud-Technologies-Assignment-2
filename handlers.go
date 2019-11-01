package assignment_2

import (
	"fmt"
	"net/http"
)

func HandlerNil(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Default Handler: Invalid request received.") // error to console
	http.Error(w, "Invalid request", http.StatusBadRequest)   // error to http
}

func HandlerCommits(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

func HandlerLanguages(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

func HandlerIssues(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}

func HandlerWebhooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Unimplemented handler used")                                        // error to console
	http.Error(w, "This endpoint is not implemented yet", http.StatusNotImplemented) // error to http
}
