package tests

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/BBennett92/cookiejar/internal/analyzer"
)

func TestAnalyzeCookie(t *testing.T) {
	testURL, _ := url.Parse("https://example.com")
	testCookie := &http.Cookie{
		Name:     "test_cookie",
		Value:    "test_value",
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteDefaultMode,
		Domain:   "example.com",
		Path:     "/",
	}

	result := analyzer.AnalyzeCookie(testCookie, testURL)

	if result.Name != "test_cookie" {
		t.Errorf("Expected cookie name 'test_cookie', got '%s'", result.Name)
	}

	if result.HttpOnly {
		t.Errorf("Expected HttpOnly to be false")
	}

	if result.Secure {
		t.Errorf("Expected Secure to be false")
	}

	if result.SameSite != "Default" && result.SameSite != "" {
		t.Errorf("Expected SameSite to be 'Default' or '', got '%s'", result.SameSite)
	}

	if len(result.Issues) == 0 {
		t.Errorf("Expected issues to be reported")
	}
}
