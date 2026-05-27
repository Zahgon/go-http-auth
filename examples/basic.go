//go:build ignore
// +build ignore

/*
 Example application using Basic auth

 Build with:

 go build basic.go
*/

package main

import (
	"net/http"

	auth ".."
)

func secret(user, realm string) string {
	_ = "STUB: not implemented"

	// password is "hello"
	return ""
}

func handle(w http.ResponseWriter, r *auth.AuthenticatedRequest) { _ = "STUB: not implemented"; return }

func main() {
	authenticator := auth.NewBasicAuthenticator("example.com", secret)
	http.HandleFunc("/", authenticator.Wrap(handle))
	http.ListenAndServe(":8080", nil)
}
