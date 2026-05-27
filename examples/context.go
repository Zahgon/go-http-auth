//go:build ignore
// +build ignore

/*
 Example application using NewContext/FromContext

 Build with:

 go build context.go
*/

package main

import (
	"context"
	"net/http"

	auth ".."
)

func secret(user, realm string) string {
	_ = "STUB: not implemented"

	// password is "hello"
	return ""
}

type contextHandler interface {
	ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

type contextHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

func (f contextHandlerFunc) ServeHTTP(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func handle(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func authenticatedHandler(a auth.AuthenticatorInterface, h contextHandler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func main() {
	authenticator := auth.NewDigestAuthenticator("example.com", secret)
	http.Handle("/", authenticatedHandler(authenticator, contextHandlerFunc(handle)))
	http.ListenAndServe(":8080", nil)
}
