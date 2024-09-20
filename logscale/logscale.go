// logscale/logscale.go
package logscale

import "fmt"

// LogEntry is the structure for a log entry
type LogEntry struct {
    Level   string
    Message string
    Service string
}

// Logger is the structure that handles logging
type Logger struct {
    bufferSize int
    logs       []LogEntry
}

// NewLogger initializes a new Logger with a given buffer size
func NewLogger(bufferSize int) *Logger {
    return &Logger{bufferSize: bufferSize}
}

// Log adds a log entry to the logger
func (l *Logger) Log(level, message, service string) {
    entry := LogEntry{Level: level, Message: message, Service: service}
    l.logs = append(l.logs, entry)
    handleLog(entry)
}

// Stop finalizes the logging process
func (l *Logger) Stop() {
    fmt.Println("Logger stopped.")
}
