package main

/*
Level 2: API Automation using Go lang (standard test framework or ginkgo)
Implement e2e API test for single node install of Rancher UI.
Test to cover: Login into Rancher
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {

	// Task is to login into Rancher UI. 2 Steps
	// First create token with user credentials [ POST Call]
	// Second login into UI thro the token [ GET Call]

	rancher_url := "http://192.168.1.11:8080"
	token_endpoint_url := rancher_url + "/v2-beta/token"
	// Rancher UI Login credentials
	var userCreds = []byte(`{
		"code": "krishna:krishna123", 
		"authProvider": "localauthconfig"
	}`)
	// Post call for Token. Payload is User Credentials
	request, error := http.NewRequest("POST", token_endpoint_url, bytes.NewBuffer(userCreds))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	fmt.Println("Token POST Call Response Status:", response.Status)
	body, _ := io.ReadAll(response.Body)
	// Extract token from POST Response
	type Result struct {
		JWT string `json:"jwt"`
	}
	var result Result
	json.Unmarshal(body, &result)

	// GET Call with token for login into UI
	loginurl := rancher_url + "?token=" + result.JWT
	response, err := http.Get(loginurl)
	if err != nil {
		fmt.Printf("Login HTTP GET request failed with error %s\n", err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}
	fmt.Println("response Status:", response.Status)
}
