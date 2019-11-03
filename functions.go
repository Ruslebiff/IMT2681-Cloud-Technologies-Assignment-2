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

/*
// CallURL calls given URL with given content and awaits response (status and body).
func CallURL(url string, content string) {
	fmt.Println("Attempting invocation of url " + url + " ...")
	res, err := http.Post(url, "string", bytes.NewReader([]byte(content)))
	if err != nil {
		fmt.Println("Error in HTTP request: " + err.Error())
	}
	response, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Something is wrong with invocation response: " + err.Error())
	}

	fmt.Println("Webhook invoked. Received status code " + strconv.Itoa(res.StatusCode) +
		" and body: " + string(response))
}
*/

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
