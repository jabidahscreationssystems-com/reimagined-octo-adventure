# Development Guide for Agent Cortana

This guide helps developers contribute to and extend Agent Cortana.

## Getting Started

### Prerequisites

- Go 1.24 or later
- Git
- A GitHub account
- ngrok or similar tool for local webhook testing

### Setting Up Development Environment

1. Clone the repository:
   ```bash
   git clone https://github.com/jabidahscreationssystems-com/reimagined-octo-adventure.git
   cd reimagined-octo-adventure/apps/agent-cortana
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Copy the example environment file:
   ```bash
   cp .env.example .env
   ```

4. Create a GitHub App for development (see DEPLOYMENT.md)

5. Update `.env` with your credentials

### Running Locally

```bash
make run
```

Or with Go directly:
```bash
go run cmd/server/main.go
```

### Testing Webhooks Locally

Since GitHub webhooks require a public URL, use ngrok for local development:

1. Install ngrok:
   ```bash
   # macOS
   brew install ngrok
   
   # Linux
   wget https://bin.equinox.io/c/4VmDzA7iaHb/ngrok-stable-linux-amd64.zip
   unzip ngrok-stable-linux-amd64.zip
   ```

2. Start your local server:
   ```bash
   make run
   ```

3. In another terminal, start ngrok:
   ```bash
   ngrok http 8080
   ```

4. Update your GitHub App webhook URL with the ngrok URL (e.g., `https://abc123.ngrok.io/webhooks`)

5. Trigger events by interacting with a repository where your app is installed

## Project Structure

```
apps/agent-cortana/
├── cmd/
│   └── server/              # Main application
│       └── main.go
├── pkg/
│   ├── config/              # Configuration management
│   │   └── config.go
│   ├── github/              # GitHub API client
│   │   └── client.go
│   ├── handlers/            # HTTP handlers
│   │   └── handlers.go
│   └── server/              # HTTP server
│       └── server.go
├── docs/                    # Documentation
├── app.yml                  # GitHub App manifest
├── Dockerfile               # Docker configuration
├── docker-compose.yml       # Docker Compose configuration
├── Makefile                 # Build automation
└── README.md                # Main documentation
```

## Adding New Features

### Adding a New Webhook Handler

1. Add the event type to `app.yml` if not already present

2. Update the webhook handler in `pkg/handlers/handlers.go`:

```go
switch eventType {
case "new_event":
    log.Printf("Processing new_event")
    handleNewEvent(payload)
// ... existing cases
}
```

3. Implement the handler function:

```go
func handleNewEvent(payload map[string]interface{}) error {
    // Process the event
    return nil
}
```

### Adding New API Endpoints

1. Add the route in `pkg/server/server.go`:

```go
func (s *Server) routes() {
    // ... existing routes
    s.mux.HandleFunc("/api/new-endpoint", handlers.HandleNewEndpoint)
}
```

2. Implement the handler in `pkg/handlers/handlers.go`:

```go
func HandleNewEndpoint(w http.ResponseWriter, r *http.Request) {
    // Implementation
}
```

### Extending the GitHub Client

Add new methods to `pkg/github/client.go`:

```go
func (c *Client) NewMethod(ctx context.Context, params string) error {
    // Implementation
    return nil
}
```

## Testing

### Unit Tests

Create test files alongside your code:

```go
// pkg/handlers/handlers_test.go
package handlers

import (
    "testing"
)

func TestHandleRoot(t *testing.T) {
    // Test implementation
}
```

Run tests:
```bash
make test
```

### Integration Tests

For integration testing with GitHub:

1. Create a test repository
2. Install your development GitHub App
3. Trigger events and verify behavior

## Code Style

Follow Go best practices:

- Use `gofmt` to format code
- Follow effective Go guidelines
- Write meaningful comments
- Keep functions small and focused
- Use meaningful variable names

Format your code:
```bash
go fmt ./...
```

## Debugging

### Enable Debug Logging

Add debug logging to your code:

```go
import "log"

log.Printf("Debug: %+v", data)
```

### Using Delve Debugger

Install Delve:
```bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Debug your application:
```bash
dlv debug ./cmd/server
```

## Common Development Tasks

### Adding Environment Variables

1. Add to `.env.example`
2. Update `pkg/config/config.go`
3. Document in README.md

### Updating Dependencies

```bash
go get -u ./...
go mod tidy
```

### Building for Production

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o agent-cortana ./cmd/server
```

## Contributing

1. Fork the repository
2. Create a feature branch: `git checkout -b feature/my-feature`
3. Make your changes
4. Write tests
5. Ensure tests pass: `make test`
6. Commit your changes: `git commit -am 'Add new feature'`
7. Push to the branch: `git push origin feature/my-feature`
8. Submit a pull request

## Resources

- [GitHub Apps Documentation](https://docs.github.com/en/apps)
- [GitHub Webhooks](https://docs.github.com/en/webhooks)
- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)

## Getting Help

- Open an issue in the repository
- Check existing issues and pull requests
- Review the documentation
