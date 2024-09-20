// logscale/batch_processor.go
package logscale

// ProcessLogsInBatch processes logs in batches
func ProcessLogsInBatch(entries []LogEntry) {
    for _, entry := range entries {
        handleLog(entry)
    }
}
