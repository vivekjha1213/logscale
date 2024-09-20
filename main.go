package main

import (
    "github.com/vivekjha1213/logscale" // Import the `logscale` package.
    "fmt"
)

func main() {
    logger := logscale.NewLogger(100) // Example logger instantiation.
    logger.Log("INFO", "Application started", "MainService")
    logger.Log("ERROR", "Something went wrong", "MainService")
    logger.Stop()
    fmt.Println("Logging done.")
}
