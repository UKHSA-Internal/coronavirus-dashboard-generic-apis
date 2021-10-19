package middleware

import (
	"net/http"
	"os"
	"strings"
)

var headers = map[string]string{
	"content-type":              "application/vnd.UKHSA-COVID19.v1+json; charset=utf-8",
	"server":                    "UKHSA API Service v.1 (Unix) - Go",
	"Strict-Transport-Security": "max-age=31536000; includeSubdomains; preload",
	"x-frame-options":           "deny",
	"x-content-type-options":    "nosniff",
	"x-xss-protection":          "1; mode=block",
	"referrer-policy":           "origin-when-cross-origin, strict-origin-when-cross-origin",
	"content-security-policy": "default-src 'self' coronavirus.data.gov.uk *.coronavirus.data.gov.uk; " +
		"style-src 'self' 'unsafe-inline' coronavirus.data.gov.uk *.coronavirus.data.gov.uk",
	"x-UKHSA-media-type": "UKHSA-COVID19.generic",
	"x-UKHSA-server-loc": os.Getenv("SERVER_LOCATION"),
}

func HeadersMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header()
		for key, value := range headers {
			w.Header().Set(key, value)
		}

		if strings.Contains(r.URL.Path, "rss") {
			w.Header().Set("content-type", "application/rss+xml")
		} else if strings.Contains(r.URL.Path, "atom") {
			w.Header().Set("content-type", "application/atom+xml")
		}

		next.ServeHTTP(w, r)
	})

} // LoggingMiddleware
