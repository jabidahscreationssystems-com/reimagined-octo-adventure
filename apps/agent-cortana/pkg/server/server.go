package server

import (
	"net/http"

	"github.com/cli/cli/v2/apps/agent-cortana/pkg/handlers"
)

// Server represents the Agent Cortana HTTP server
type Server struct {
	mux *http.ServeMux
}

// New creates a new Server instance
func New() *Server {
	s := &Server{
		mux: http.NewServeMux(),
	}
	s.routes()
	return s
}

// ServeHTTP implements the http.Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// routes sets up the HTTP routes
func (s *Server) routes() {
	s.mux.HandleFunc("/", handlers.HandleRoot)
	s.mux.HandleFunc("/health", handlers.HandleHealth)
	s.mux.HandleFunc("/webhooks", handlers.HandleWebhook)
}
