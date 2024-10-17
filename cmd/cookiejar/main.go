package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BBennett92/cookiejar/internal/analyzer"
	"github.com/BBennett92/cookiejar/internal/utils"
)

func main() {
	// Command-line flags
	urlPtr := flag.String("url", "", "URL of the website to scan")
	outputPtr := flag.String("output", "text", "Output format: text, json, csv")
	configPtr := flag.String("config", "", "Path to configuration file")
	flag.Parse()

	if *urlPtr == "" {
		fmt.Println("Please provide a URL to scan using the -url flag.")
		os.Exit(1)
	}

	// Load configurations
	config, err := utils.LoadConfig(*configPtr)
	if err != nil {
		fmt.Printf("Error loading configuration: %v\n", err)
		os.Exit(1)
	}

	// Scan cookies
	results := analyzer.ScanCookies(*urlPtr, config)

	// Output results
	err = analyzer.OutputResults(results, *outputPtr)
	if err != nil {
		fmt.Printf("Error outputting results: %v\n", err)
		os.Exit(1)
	}
}
