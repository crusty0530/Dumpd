package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	handler := Status{Message: "Server Ok!"}

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status: got %d, want %d", rec.Code, http.StatusOK)
	}

	want := "Server Ok!"
	if got := rec.Body.String(); got != want {
		t.Errorf("body: got %q, want %q", got, want)
	}
}
