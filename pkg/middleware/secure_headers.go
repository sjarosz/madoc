package middleware

import (
	"net/http"

	"github.com/sqoopdata/madoc/pkg/application"
)

// SecureHeaders add required secure headers to a response
func SecureHeaders(next http.HandlerFunc, a *application.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Below headers are from the OWASP Secure Headers Project
		// HTTP Strict Transport Security
		//	- Prevents against protocol downgrade attack
		// X-Frame-Options
		// X-Content-Type-Options
		// Content-Security-Policy
		// X-Permitted-Cross-Domain-Policies
		//	- XML document to provide permission to handle data across domains
		// Referrer-Policy
		//	- Governs which referrer information sent should be included with requests
		// Clear-Site-Data
		//	- Clears browsing data associated with requesting website
		// Cross-Origin-Embedder-Policy
		//	- Prevents a document from loading cross-origin resources that dont
		//	  explicitly grant document persmission
		// Cross-Origin-Opener-Policy
		//	- Allows you to ensure a top-level document does not share a browsing
		//	  context group with cross-origin documents.
		// Cross-Origin-Resource-Policy
		//	- Allows to define a policy that lets web sites and applications opt in
		//	  to protection against certain requests from other origins (such as
		//	  those issued with elements like <script> and <img>), to mitigate
		//	  speculative side-channel attacks, like Spectre, as well as Cross-Site
		// 	  Script Inclusion (XSSI) attacks (source Mozilla MDN).
		rHeaders := []struct {
			Name  string
			Value string
		}{
			{"Content-Type", "application/json;charset=utf-8"},
			{"X-Content-Type-Options", "nosniff"}, // prevents browser guess correct content-type
			{"X-Frame-Options", "DENY"},           // prevents api responses loaded in frame or iframe
			{"X-XSS-Protection", "0"},             // tells browser whether to block/ignore suspected xss attacks. current guidance is to use 0
			{"Cache-Control", "no-store"},         // controls whether browser/proxies can cache content in the response and for how long

			// Reduces scope of xss attacks by restricting where scripts can be loaded from
			// and what they can do. CSP is valuable defence against xss attacks.
			// * default-src: prevents response from loading any scripts or resources
			// * frame-ancestors: replacement of x-frame-options and prevents response from loadingin iframe
			// * sandbox: disables scripts and other potentially dangerous content from being executed
			{"Content-Security-Policy", "default-src 'none'; frame-ancestors 'none'; sandbox"},
			{"Server", "Monitored"}, // hides server details in the response
		}

		for _, rH := range rHeaders {
			w.Header().Set(rH.Name, rH.Value)
		}

		next(w, r)
	}
}
