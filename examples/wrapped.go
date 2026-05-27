//go:build ignore
// +build ignore

/*
 Example demonstrating how to wrap an application which is unaware of
 authenticated requests with a "pass-through" authentication

 Build with:

 go build wrapped.go
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

func regularHandler(w http.ResponseWriter, r *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	authenticator := auth.NewBasicAuthenticator("example.com", secret)
	http.HandleFunc("/", auth.JustCheck(authenticator, regularHandler))
	http.ListenAndServe(":8080", nil)
}
