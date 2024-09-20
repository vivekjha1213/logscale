// logscale/handler.go
package logscale

import "fmt"

// LogHandler is an interface for custom log processing
type LogHandler interface {
    HandleLog(entry LogEntry) error
}

// ExampleHandler is a simple example of a custom log handler
type ExampleHandler struct{}

// HandleLog processes a log entry with the ExampleHandler
func (h *ExampleHandler) HandleLog(entry LogEntry) error {
    fmt.Printf("Log Entry: %v\n", entry)
    return nil
}

// handleLog processes the log entry using the current handler
var logHandler LogHandler = &ExampleHandler{}

// SetLogHandler allows users to set a custom log handler
func SetLogHandler(handler LogHandler) {
    logHandler = handler
}

// handleLog processes the log entry using the configured handler
func handleLog(entry LogEntry) {
    if err := logHandler.HandleLog(entry); err != nil {
        fmt.Printf("Failed to handle log: %v\n", err)
    }
}
