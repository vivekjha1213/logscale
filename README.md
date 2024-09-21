# LogScale

<div align="center">
  <img src="https://github.com/user-attachments/assets/cec439cc-4974-4cd7-b3d0-512459dd3896" alt="LogScale Logo" width="1000" height="200"/>
</div>

  
  [![Go Report Card](https://goreportcard.com/badge/github.com/vivekjha1213/logscale)](https://goreportcard.com/report/github.com/vivekjha1213/logscale)
  [![GoDoc](https://godoc.org/github.com/vivekjha1213/logscale?status.svg)](https://godoc.org/github.com/vivekjha1213/logscale)
  [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
</div>

**LogScale** is a scalable logging library written in Go, designed to handle logs efficiently with customizable log handlers, batch processing, and flexible logging levels. This package helps you implement scalable and structured logging in your Go applications.

## Features

- **Customizable Handlers**: Support for custom log handlers that can be implemented according to your needs.
- **Batch Processing**: Process logs in batches to improve performance and reduce I/O operations.
- **Scalable Architecture**: Efficient logging even with high loads due to its buffer-based design.
- **Multiple Logging Levels**: Support for different log levels like `INFO`, `DEBUG`, `WARN`, `ERROR`, etc.
- **Asynchronous Logging**: Non-blocking log calls for improved application performance.
- **Structured Logging**: Support for key-value pairs for better log analysis.

## Installation

```bash
go get github.com/vivekjha1213/logscale
```

## Usage

### Initializing the Logger

To start using LogScale, create a logger with a specified buffer size and log messages with different levels.

```go
package main

import (
    "github.com/vivekjha1213/logscale"
)

func main() {
    // Create a new logger with a buffer size of 100 entries
    logger := logscale.NewLogger(100)

    // Log an info message
    logger.Log("INFO", "Application started", "MyService")

    // Log an error message
    logger.Log("ERROR", "Failed to load configuration", "MyService")

    // Stop the logger when done
    logger.Stop()
}
```

### Custom Log Handlers

You can create custom log handlers by implementing the `LogHandler` interface. For example:

```go
package main

import (
    "fmt"
    "github.com/vivekjha1213/logscale"
)

// CustomHandler processes logs by printing them in a custom format
type CustomHandler struct{}

func (h *CustomHandler) HandleLog(entry logscale.LogEntry) error {
    fmt.Printf("Custom Log [%s]: %s - %s\n", entry.Level, entry.Service, entry.Message)
    return nil
}

func main() {
    logger := logscale.NewLogger(100)

    // Create and set a custom handler
    customHandler := &CustomHandler{}
    logscale.SetLogHandler(customHandler)

    logger.Log("INFO", "Custom logging handler set", "MyService")
    logger.Stop()
}
```

### Batch Processing Logs

LogScale allows logs to be processed in batches to increase efficiency under heavy loads. You can define batch sizes and intervals based on your system's needs.

```go
package main

import (
    "github.com/vivekjha1213/logscale"
    "time"
)

func main() {
    logger := logscale.NewLogger(1000)
    
    // Configure batch processing
    logger.ConfigureBatchProcessing(100, 5 * time.Second)

    // Log messages will be processed in batches of 100 or every 5 seconds, whichever comes first
    for i := 0; i < 1000; i++ {
        logger.Log("INFO", fmt.Sprintf("Log entry %d", i), "BatchService")
    }

    logger.Stop()
}
```

### Structured Logging

LogScale supports structured logging with key-value pairs for better log analysis:

```go
logger.LogStructured("INFO", "User logged in", "AuthService", map[string]interface{}{
    "user_id": 12345,
    "ip_address": "192.168.1.1",
    "login_time": time.Now(),
})
```

## Testing

To test your logging functionality, use Go's standard testing tools. For example:

```bash
go test ./logscale
```

### Example Test Case

```go
package logscale_test

import (
    "testing"
    "github.com/vivekjha1213/logscale"
)

func TestLogging(t *testing.T) {
    logger := logscale.NewLogger(10)
    logger.Log("INFO", "Test log", "TestService")
    logger.Stop()
}
```

## Benchmarking

We've included benchmarks to measure the performance of LogScale. Run them using:

```bash
go test -bench=. ./logscale
```

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request if you'd like to add features or improve the code.

1. Fork it (https://github.com/vivekjha1213/logscale/fork)
2. Create your feature branch (`git checkout -b feature/my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin feature/my-new-feature`)
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by scalable logging solutions for large-scale applications.
- Thanks to the open-source Go community for guidance and support.
- Logo created with [LogoMakr](https://logomakr.com/)
- Badges provided by [Shields.io](https://shields.io/)

## Open Source

LogScale is proud to be an open-source project. We believe in the power of community-driven development and welcome contributions from developers around the world.

<div align="center">
  <img src="https://github.com/user-attachments/assets/9085782f-90e9-4709-8d61-1a98d266e221" alt="Go Logo" width="1000" height="200"/>
</div>

---

Made with ❤️ by [Vivek Jha](https://github.com/vivekjha1213)
