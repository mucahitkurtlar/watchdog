package esp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type status struct {
	Status int `json:"status"`
}

type action struct {
	Action string `json:"action"`
}

// Checks pin status and returns bool value
// AskStatus exported
func AskStatus(url string) bool {
	// Create client object
	spaceClient := http.Client{
		// Timeout after 5 seconds
		Timeout: time.Second * 5,
	}

	// Create new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panic(err)
	}

	// Set request header
	req.Header.Set("User-Agent", "spacecount-tutorial")

	// Do the request and take the response
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Println("Can't reach the target! Target link: ", url)
		return false
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Read the response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Panic(readErr)
	}

	// Unmarshal the response
	status1 := status{}
	jsonErr := json.Unmarshal(body, &status1)
	if jsonErr != nil {
		log.Panic(jsonErr)
	}

	// Print the response
	fmt.Println(status1.Status)

	// Return the pin status
	return status1.Status == 1

}

// Sets pin status
// PinAction exported
func PinAction(url string) bool {
	spaceClient := http.Client{
		// Timeout after 5 seconds
		Timeout: time.Second * 5,
	}

	// Timeout after 2 seconds
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Panic(err)
	}

	// Set request header
	req.Header.Set("User-Agent", "spacecount-tutorial")

	// Do the request and take the response
	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Println("Can't reach the target! Target link: ", url)
		return false
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Read the response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Panic(readErr)
	}

	// Unmarshal the response
	action1 := action{}
	jsonErr := json.Unmarshal(body, &action1)
	if jsonErr != nil {
		log.Panic(jsonErr)
	}

	// Print action response
	fmt.Println(action1.Action)

	// Return `true` if the action was successful
	return action1.Action == "ok"
}
