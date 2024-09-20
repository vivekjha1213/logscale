package logscale

import (
    "encoding/json"
    "os"
    "sync"
)

type Logger struct {
    mu      sync.Mutex
    logFile *os.File
    batch   []LogEntry
    maxSize int
}

type LogEntry struct {
    Level   string `json:"level"`
    Message string `json:"message"`
    Source  string `json:"source"`
}

func NewLogger(maxSize int) *Logger {
    file, _ := os.OpenFile("logs.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    return &Logger{
        logFile: file,
        maxSize: maxSize,
    }
}

func (l *Logger) Log(level, message, source string) {
    l.mu.Lock()
    defer l.mu.Unlock()

    entry := LogEntry{Level: level, Message: message, Source: source}
    l.batch = append(l.batch, entry)

    if len(l.batch) >= l.maxSize {
        l.flush()
    }
}

func (l *Logger) flush() {
    data, _ := json.Marshal(l.batch)
    l.logFile.Write(data)
    l.logFile.WriteString("\n")
    l.batch = nil
}

func (l *Logger) Stop() {
    l.flush()
    l.logFile.Close()
}
