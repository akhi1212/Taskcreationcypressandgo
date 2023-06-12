// mainloginapi_test.go

package main

import (
	"net/http"
	"net/url"
	"testing"
)

func TestRancherLogin(t *testing.T) {
	baseURL := "https://localhost:8080/v1"
	username := "admin"
	password := "newchanges!123@0"

	client := &http.Client{}

	// Prepare the login request
	loginURL, _ := url.Parse(baseURL + "/local?action=login")
	loginData := url.Values{}
	loginData.Set("username", username)
	loginData.Set("password", password)

	req, _ := http.NewRequest(http.MethodPost, loginURL.String(), nil)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.PostForm = loginData

	// Send the login request
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Login request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Login failed with status code: %d", resp.StatusCode)
	}

  t.Logf("Response Body: %s", string(body))
}
