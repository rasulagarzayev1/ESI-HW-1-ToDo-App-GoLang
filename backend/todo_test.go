package main

import (
	"bytes"
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
	body:= []byte(`{{
		"Title": "adeu",
			"Status": true,
			"Level": 3
	}`)
	req, _ := http.NewRequest("GET", "/api/v1/tasks", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestPutRequest(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	body:= []byte(`{{
		"Title": "adeu",
			"Status": true,
			"Level": 3
	}`)
	req, _ := http.NewRequest("GET", "/api/v1/tasks", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}


