// logscale.go
package logscale


import (
    "github.com/vivekjha1213/logscale"
    "fmt"
)

func main() {
    // Initialize logger with a buffer size of 100
    logger := logscale.NewLogger(100)
    
    // Example of logging messages
    logger.Log("INFO", "Service started", "MyService")
    logger.Log("ERROR", "Failed to connect to database", "MyService")

    // Custom handler example (this could be sending logs to a cloud service)
    customHandler := &logscale.ExampleHandler{}
    logscale.SetLogHandler(customHandler)

    fmt.Println("Logging completed.")
    logger.Stop()  // Stop the logger when done
}
