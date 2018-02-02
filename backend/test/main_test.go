package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetRequest(t *testing.T) {
	response, err := http.Get("http://159.203.178.86:8000")

	if err != nil {
		t.Errorf("Not valid response from endpoint")
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	correctResponse := "Welcome, "

	if !equal(correctResponse, body) {
		t.Errorf("Incorrect response")
	}
}

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
