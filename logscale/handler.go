// logscale/handler.go
package logscale

import "fmt"

// LogHandler is an interface for custom log processing
type LogHandler interface {
    HandleLog(entry LogEntry) error
}

// ExampleHandler is a simple example of a custom log handler
type ExampleHandler struct{}

func (h *ExampleHandler) HandleLog(entry LogEntry) error {
    // Here you can define how you want to process logs (e.g., send to a cloud service)
    fmt.Printf("Log Entry: %v\n", entry)
    return nil
}

// handleLog is a placeholder for the user-defined log handler
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
