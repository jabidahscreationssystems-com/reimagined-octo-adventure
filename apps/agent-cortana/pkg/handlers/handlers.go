package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// HandleRoot handles the root endpoint
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	response := map[string]string{
		"name":        "Agent Cortana",
		"version":     "1.0.0",
		"description": "An intelligent GitHub agent for automating workflows",
		"status":      "running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// HandleHealth handles the health check endpoint
func HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "healthy",
	})
}

// HandleWebhook handles GitHub webhook events
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// TODO: Validate webhook signature using WEBHOOK_SECRET
	// signature := r.Header.Get("X-Hub-Signature-256")
	// This is critical for production to ensure webhooks come from GitHub

	// Get the event type from the header
	eventType := r.Header.Get("X-GitHub-Event")
	deliveryID := r.Header.Get("X-GitHub-Delivery")

	log.Printf("Received webhook event: %s (delivery: %s)", eventType, deliveryID)

	// Parse the webhook payload
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Printf("Error decoding webhook payload: %v", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Process the webhook based on event type
	switch eventType {
	case "issues":
		log.Printf("Processing issues event")
	case "pull_request":
		log.Printf("Processing pull_request event")
	case "issue_comment":
		log.Printf("Processing issue_comment event")
	case "pull_request_review":
		log.Printf("Processing pull_request_review event")
	case "push":
		log.Printf("Processing push event")
	case "workflow_run":
		log.Printf("Processing workflow_run event")
	default:
		log.Printf("Unhandled event type: %s", eventType)
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "received",
	})
}
