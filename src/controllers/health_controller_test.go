package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var resource = "/health"

func TestGetHealth(t *testing.T) {
	req, err := http.NewRequest("GET", resource, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	var handler = http.Handler(http.HandlerFunc(healthController.GetHealth))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("failed to request, return code: %v, payload: %v", rr.Code, rr.Body)
	}

	expected := `{"message":"health check ok"}`
	if strings.Contains(expected, rr.Body.String()) {
		t.Errorf("payload returned is invalid, returned: %s, expected: %s]", rr.Body.String(), expected)
	}
}