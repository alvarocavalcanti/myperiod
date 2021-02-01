package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestHomeHandler(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
			t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Act
	a.Router.ServeHTTP(rr, req)

	// Assert
	assert := assert.New(t)
	assert.Equal(http.StatusOK, rr.Code)
	assert.Equal("{}", rr.Body.String())
}

func TestEmptyEntryList(t *testing.T) {
	// Arrange
	req, err := http.NewRequest("GET", "/entries", nil)
	if err != nil {
			t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	// Act
	a.Router.ServeHTTP(rr, req)

	// Assert
	assert := assert.New(t)
	assert.Equal(http.StatusOK, rr.Code)
	var entries []Entry
	json.Unmarshal([]byte(rr.Body.String()), &entries)
	assert.Equal(len(entries), 0)
}
