package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Fatal("Expected server instance, got nil")
	}
	if s.mux == nil {
		t.Fatal("Expected mux to be initialized")
	}
}

func TestServerRoutes(t *testing.T) {
	s := New()

	tests := []struct {
		name           string
		path           string
		method         string
		expectedStatus int
	}{
		{
			name:           "Root endpoint",
			path:           "/",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Health endpoint",
			path:           "/health",
			method:         http.MethodGet,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Webhook endpoint with GET (should fail)",
			path:           "/webhooks",
			method:         http.MethodGet,
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			s.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
