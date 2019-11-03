package assignment2

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// countduplicates takes an array of strings as input,
// returns a map of respective strings and their
// number of occurrences in the list
func countduplicates(arr []string) map[string]int {
	arritem := make(map[string]int) // map of items in list, and their number

	for _, i := range arr { // loop through whole array
		arritem[i]++ // add occurrences to the counter map
	}

	return arritem
}

// CallWebhooks calls the webhooks for the specified event and parameters
func CallWebhooks(event string, parameters string, timestamp time.Time) {
	var webhooks []Webhookreg
	webhooks, err := DBReadall()
	if err != nil {
		fmt.Println("Error: ", err)
	}

	for i := range webhooks {
		if webhooks[i].Event == event {
			var request = WebhookPayload{Event: event, Parameters: parameters, Time: timestamp.String()}

			requestBody, err := json.Marshal(request)
			if err != nil {
				fmt.Println("Can not encode: " + err.Error())
			}

			fmt.Println("Attempting invoation of URL " + webhooks[i].URL + "...")

			resp, err := http.Post(webhooks[i].URL, "json", bytes.NewReader([]byte(requestBody)))
			if err != nil {
				fmt.Println("Error in HTTP request: " + err.Error())
			}

			jsonresponse, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error when reading response")
				return
			}

			json.Unmarshal([]byte(jsonresponse), &request)
			if err != nil {
				fmt.Println("Error when unmarshalling jsonresponse")
				return
			}
		}
	}
}
