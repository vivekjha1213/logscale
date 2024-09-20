package logscale

import (
    "bytes"
    "fmt"
    "testing"
)

// MockLogger is a mock implementation for testing
type MockLogger struct {
    buffer bytes.Buffer
}

// Log captures log messages in the buffer
func (m *MockLogger) Log(level, message, service string) {
    fmt.Fprintf(&m.buffer, "[%s] %s: %s\n", level, service, message)
}

// Stop does nothing for the mock logger
func (m *MockLogger) Stop() {
    // No operation for mock
}

// TestLogger tests the logging functionality
func TestLogger(t *testing.T) {
    logger := &MockLogger{}

    // Simulate logging
    logger.Log("INFO", "Test service started", "TestService")
    logger.Log("ERROR", "Test error occurred", "TestService")

    // Check if log messages are present in the buffer
    expected := "[INFO] TestService: Test service started\n[ERROR] TestService: Test error occurred\n"
    if logger.buffer.String() != expected {
        t.Errorf("Expected log output: %s, got: %s", expected, logger.buffer.String())
    }
}

// TestExampleHandler tests the ExampleHandler implementation
func TestExampleHandler(t *testing.T) {
    handler := &ExampleHandler{}
    entry := LogEntry{
        Level:   "INFO",
        Message: "Test log entry",
        Service: "TestService",
    }

    err := handler.HandleLog(entry)
    if err != nil {
        t.Errorf("Expected no error, got: %v", err)
    }
}
