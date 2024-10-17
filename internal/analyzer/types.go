package analyzer

// CookieResult holds the analysis result for a single cookie.
type CookieResult struct {
	Name            string   `json:"name"`
	HttpOnly        bool     `json:"httpOnly"`
	Secure          bool     `json:"secure"`
	SameSite        string   `json:"sameSite"`
	DomainScope     bool     `json:"domainScope"`
	PathScope       bool     `json:"pathScope"`
	Expiration      string   `json:"expiration"`
	Issues          []string `json:"issues"`
	Recommendations []string `json:"recommendations"`
}
