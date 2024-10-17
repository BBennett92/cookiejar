package analyzer

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"

	"github.com/BBennett92/cookiejar/internal/utils"
)

// ScanCookies scans the cookies from the provided URL.
func ScanCookies(targetURL string, config utils.Config) []CookieResult {
	parsedURL, err := url.Parse(targetURL)
	if err != nil {
		fmt.Printf("Invalid URL: %v\n", err)
		os.Exit(1)
	}

	timeout := config.Scan.TimeoutDuration
	if timeout == 0 {
		timeout = 10 * time.Second
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if config.Scan.FollowRedirects {
				return nil
			}
			return http.ErrUseLastResponse
		},
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", parsedURL.String(), nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	cookies := resp.Cookies()
	if len(cookies) == 0 {
		fmt.Println("No cookies found.")
		os.Exit(0)
	}

	var wg sync.WaitGroup
	results := make([]CookieResult, len(cookies))

	for i, cookie := range cookies {
		wg.Add(1)
		go func(i int, cookie *http.Cookie) {
			defer wg.Done()
			results[i] = AnalyzeCookie(cookie, parsedURL)
		}(i, cookie)
	}

	wg.Wait()
	return results
}

// AnalyzeCookie analyzes a single cookie and returns the result.
func AnalyzeCookie(cookie *http.Cookie, parsedURL *url.URL) CookieResult {
	result := CookieResult{
		Name:     cookie.Name,
		HttpOnly: cookie.HttpOnly,
		Secure:   cookie.Secure,
		SameSite: sameSiteToString(cookie.SameSite),
	}

	// Check Domain Scope
	result.DomainScope = (cookie.Domain == "" || cookie.Domain == parsedURL.Hostname())

	if !result.DomainScope {
		result.Issues = append(result.Issues, "Cookie domain scope is too broad.")
		result.Recommendations = append(result.Recommendations, "Set cookie domain to a more specific scope.")
	}

	// Check Path Scope
	result.PathScope = (cookie.Path == "" || cookie.Path == "/")

	if !result.PathScope {
		result.Issues = append(result.Issues, "Cookie path scope is too broad.")
		result.Recommendations = append(result.Recommendations, "Set cookie path to a more specific scope.")
	}

	// Check Expiration
	if cookie.Expires.IsZero() {
		result.Expiration = "Session"
	} else {
		result.Expiration = cookie.Expires.Format(time.RFC1123)
	}

	// Analyze flags and attributes
	if !cookie.HttpOnly {
		result.Issues = append(result.Issues, "HttpOnly flag is not set.")
		result.Recommendations = append(result.Recommendations, "Add HttpOnly flag to prevent client-side scripts from accessing the cookie.")
	}

	if !cookie.Secure {
		result.Issues = append(result.Issues, "Secure flag is not set.")
		result.Recommendations = append(result.Recommendations, "Add Secure flag to ensure cookie is sent over HTTPS only.")
	}

	if cookie.SameSite == http.SameSiteDefaultMode || cookie.SameSite == 0 {
		result.Issues = append(result.Issues, "SameSite attribute is not set.")
		result.Recommendations = append(result.Recommendations, "Set SameSite attribute to Lax or Strict to mitigate CSRF attacks.")
	}

	// Additional checks can be implemented here (e.g., size, encryption)

	return result
}

// sameSiteToString converts the http.SameSite value to its string representation.
func sameSiteToString(sameSite http.SameSite) string {
	switch sameSite {
	case http.SameSiteDefaultMode:
		return "Default"
	case http.SameSiteLaxMode:
		return "Lax"
	case http.SameSiteStrictMode:
		return "Strict"
	case http.SameSiteNoneMode:
		return "None"
	default:
		return "Unknown"
	}
}
