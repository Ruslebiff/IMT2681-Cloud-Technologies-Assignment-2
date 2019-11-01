package main

import (
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
	//http.HandleFunc("/", handlerNil)

	// print to console
	fmt.Println("Program started: ", StartTime)
	fmt.Println("Listening on port " + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
