package main

import (
	"bytes"
	"encoding/json"
	_ "fmt"
	"github.com/stretchr/testify/assert"
	_ "github.com/stretchr/testify/assert"
	"net/http"
	_ "net/http"
	"net/http/httptest"
	_ "net/http/httptest"
	"testing"
	_ "testing"
)

func TestGetRequest(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetRequestId(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/tasks/1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPostRequest(t *testing.T) {

	router := setupRouter()
	w := httptest.NewRecorder()

	postBody, _ := json.Marshal(map[string]string{
		"title": "some title",
	})
	responseBody := bytes.NewBuffer(postBody)

	req, _ := http.NewRequest("POST", "/api/v1/tasks", responseBody)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 201, w.Code)
}


func TestPutRequest(t *testing.T) {

	router := setupRouter()
	w := httptest.NewRecorder()

	postBody, _ := json.Marshal(map[string]string{
		"title": "some modification",
	})
	responseBody := bytes.NewBuffer(postBody)

	req, _ := http.NewRequest("PUT", "/api/v1/tasks/4", responseBody)
	req.Header.Add("Content-Type", "application/json")

	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}


