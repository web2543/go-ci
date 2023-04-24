package main

import (
	"argo/api/hope"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestVersionRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()
	w := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/version", nil)
	if err != nil {
		t.Fatalf("Err:%v", err)
	}
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Fatalf("Return code: %d expect: 200", w.Code)
	}
	var version hope.Version
	json.NewDecoder(w.Body).Decode(&version)
	if version.Version != "0.0.1" {
		t.Fatalf("Return code: %s expect: 0.0.1", version.Version)
	}
}
