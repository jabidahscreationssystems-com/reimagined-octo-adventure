package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	HandleRoot(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s", contentType)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Agent Cortana") {
		t.Error("Expected response to contain 'Agent Cortana'")
	}
}

func TestHandleHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HandleHealth(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	body := w.Body.String()
	if !strings.Contains(body, "healthy") {
		t.Error("Expected response to contain 'healthy'")
	}
}

func TestHandleWebhook(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		eventType      string
		body           string
		expectedStatus int
	}{
		{
			name:           "Valid POST request",
			method:         http.MethodPost,
			eventType:      "issues",
			body:           `{"action":"opened"}`,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "Invalid method",
			method:         http.MethodGet,
			eventType:      "issues",
			body:           "",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "Invalid JSON",
			method:         http.MethodPost,
			eventType:      "issues",
			body:           `invalid json`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, "/webhooks", strings.NewReader(tt.body))
			req.Header.Set("X-GitHub-Event", tt.eventType)
			req.Header.Set("X-GitHub-Delivery", "test-delivery-id")
			w := httptest.NewRecorder()

			HandleWebhook(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}
