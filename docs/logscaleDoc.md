### Doc Example for `logscale`

```go
/*
Package logscale provides a lightweight logging framework that supports 
custom log handling and batch processing.

This package allows users to create a logger with configurable buffer sizes, 
log messages with different severity levels, and utilize custom handlers for 
log processing.

Usage:

    import "github.com/vivekjha1213/logscale"

    // Create a new logger with a buffer size of 100
    logger := logscale.NewLogger(100)

    // Log messages
    logger.Log("INFO", "Service started", "MyService")
    logger.Log("ERROR", "Failed to connect to database", "MyService")

    // Set a custom log handler
    customHandler := &logscale.ExampleHandler{}
    logscale.SetLogHandler(customHandler)

    logger.Stop()  // Stop the logger when done

Logging Structure:
- Log entries consist of a level (e.g., INFO, ERROR), a message, and the associated service.
- Custom handlers can be implemented by creating types that satisfy the LogHandler interface.

ExampleHandler provides a simple implementation of the LogHandler interface.
*/

package logscale

import (
    "fmt"
)

// LogEntry represents a single log entry with its details.
type LogEntry struct {
    Level   string // The severity level of the log (e.g., INFO, ERROR)
    Message string // The log message
    Service string // The name of the service that generated the log
}

// LogHandler is an interface for custom log processing.
type LogHandler interface {
    HandleLog(entry LogEntry) error
}

// ExampleHandler is a simple example of a custom log handler.
type ExampleHandler struct{}

// HandleLog processes the log entry, for example, by printing it.
func (h *ExampleHandler) HandleLog(entry LogEntry) error {
    fmt.Printf("Log Entry: %v\n", entry)
    return nil
}

// Logger manages the log entries and their processing.
type Logger struct {
    bufferSize int
}

// NewLogger creates a new logger with the specified buffer size.
func NewLogger(bufferSize int) *Logger {
    return &Logger{bufferSize: bufferSize}
}

// Log records a log entry.
func (l *Logger) Log(level, message, service string) {
    fmt.Printf("[%s] %s: %s\n", level, service, message)
}

// Stop stops the logger.
func (l *Logger) Stop() {
    fmt.Println("Logger stopped.")
}

// SetLogHandler allows users to set a custom log handler.
func SetLogHandler(handler LogHandler) {
    // Implementation here...
}
```

### Key Points to Include
- **Package Overview**: A brief description of what the package does.
- **Usage Examples**: How to use the package, including creating a logger and logging messages.
- **Types and Functions**: Description of key types (like `LogEntry`, `Logger`, etc.) and their methods.
- **Custom Handlers**: Explain how to create and use custom log handlers.

