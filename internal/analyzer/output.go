package analyzer

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
)

// OutputResults outputs the analysis results in the specified format.
func OutputResults(results []CookieResult, format string) error {
	switch format {
	case "text":
		OutputText(results)
	case "json":
		return OutputJSON(results)
	case "csv":
		return OutputCSV(results)
	default:
		return fmt.Errorf("unsupported output format: %s", format)
	}
	return nil
}

// OutputText outputs the results in plain text format.
func OutputText(results []CookieResult) {
	for _, result := range results {
		fmt.Printf("Cookie Name: %s\n", result.Name)
		fmt.Printf("- HttpOnly: %v\n", result.HttpOnly)
		fmt.Printf("- Secure: %v\n", result.Secure)
		fmt.Printf("- SameSite: %s\n", result.SameSite)
		fmt.Printf("- Expiration: %s\n", result.Expiration)
		if len(result.Issues) > 0 {
			fmt.Println("Issues:")
			for _, issue := range result.Issues {
				fmt.Printf("  - %s\n", issue)
			}
			fmt.Println("Recommendations:")
			for _, rec := range result.Recommendations {
				fmt.Printf("  - %s\n", rec)
			}
		} else {
			fmt.Println("No issues found.")
		}
		fmt.Println()
	}
}

// OutputJSON outputs the results in JSON format.
func OutputJSON(results []CookieResult) error {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(jsonData))
	return nil
}

// OutputCSV outputs the results in CSV format.
func OutputCSV(results []CookieResult) error {
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()

	// Write header
	header := []string{"Name", "HttpOnly", "Secure", "SameSite", "Expiration", "Issues", "Recommendations"}
	if err := writer.Write(header); err != nil {
		return err
	}

	for _, result := range results {
		record := []string{
			result.Name,
			fmt.Sprintf("%v", result.HttpOnly),
			fmt.Sprintf("%v", result.Secure),
			result.SameSite,
			result.Expiration,
			fmt.Sprintf("%v", result.Issues),
			fmt.Sprintf("%v", result.Recommendations),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
