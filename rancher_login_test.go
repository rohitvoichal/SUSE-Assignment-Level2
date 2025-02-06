package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const rancherURL = "https://localhost/dashboard/auth/login"

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func TestRancherLogin(t *testing.T) {
	payload := LoginPayload{
		Username: "admin",
		Password: "QG9cb9mNRS8H9WpS",
	}
	jsonData, _ := json.Marshal(payload)

	// ✅ Create a custom HTTP client that ignores SSL errors
	client := &http.Client{
		Timeout: 10 * time.Second, // Set a timeout
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // 🔥 Disable SSL verification
		},
	}

	// ✅ Send request with custom client
	resp, err := client.Post(rancherURL, "application/json", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Login should return 200 OK")
}
