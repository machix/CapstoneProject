package main

import (
	"testing"
)

//Test basic get request to API
func TestGetRequest(t *testing.T) {
	//Implemented again once api has changed.
}

//Test get requestion to /position endpoint
func TestGetPositionRequest(t *testing.T) {
	// response, err := http.Get("http://159.203.178.86:8000/position")

	// if err != nil {
	// 	t.Errorf("Error. Not a valid response from endpoint.")
	// }

	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// correctResponse := "{\"Text\":\"Soon you will get some really cool info herer! It will be very cool!\"}"

	// if !equal(correctResponse, body) {
	// 	t.Errorf("Incorrect response, %s", string(body[:len(body)]))
	// }
}

//Compares string and byte array to determine equality
func equal(s string, b []byte) bool {
	if len(s) != len(b) {
		return false
	}
	for i, x := range b {
		if x != s[i] {
			return false
		}
	}
	return true
}
