package auth

import (
	"context"
	"net/http"
	"sync"
)

type digestClient struct {
	nc       uint64
	lastSeen int64
}

// DigestAuth is an authenticator implementation for 'Digest' HTTP Authentication scheme (RFC 7616).
//
// Note: this implementation was written following now deprecated RFC
// 2617, and supports only MD5 algorithm.
//
// TODO: Add support for SHA-256 and SHA-512/256 algorithms.
type DigestAuth struct {
	Realm            string
	Opaque           string
	Secrets          SecretProvider
	PlainTextSecrets bool
	IgnoreNonceCount bool
	// Headers used by authenticator. Set to ProxyHeaders to use with
	// proxy server. When nil, NormalHeaders are used.
	Headers *Headers

	/*
	   Approximate size of Client's Cache. When actual number of
	   tracked client nonces exceeds
	   ClientCacheSize+ClientCacheTolerance, ClientCacheTolerance*2
	   older entries are purged.
	*/
	ClientCacheSize      int
	ClientCacheTolerance int

	clients map[string]*digestClient
	mutex   sync.RWMutex
}

// check that DigestAuth implements AuthenticatorInterface
var _ = (AuthenticatorInterface)((*DigestAuth)(nil))

type digestCacheEntry struct {
	nonce    string
	lastSeen int64
}

type digestCache []digestCacheEntry

func (c digestCache) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (c digestCache) Len() int { _ = "STUB: not implemented"; return 0 }

func (c digestCache) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Purge removes count oldest entries from DigestAuth.clients
func (da *DigestAuth) Purge(count int) { _ = "STUB: not implemented"; return }

func (da *DigestAuth) purgeLocked(count int) { _ = "STUB: not implemented"; return }

// RequireAuth is an http.HandlerFunc which initiates the
// authentication process (or requires reauthentication).
func (da *DigestAuth) RequireAuth(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// DigestAuthParams parses Authorization header from the
// http.Request. Returns a map of auth parameters or nil if the header
// is not a valid parsable Digest auth header.
func DigestAuthParams(authorization string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// CheckAuth checks whether the request contains valid authentication
// data. Returns a pair of username, authinfo, where username is the
// name of the authenticated user or an empty string and authinfo is
// the contents for the optional Authentication-Info response header.
func (da *DigestAuth) CheckAuth(r *http.Request) (username string, authinfo *string) {
	_ = "STUB: not implemented"
	return "", nil
}

// RFC2617 Section 3.2.1 specifies that unset value of algorithm in
// WWW-Authenticate Response header should be treated as
// "MD5". According to section 3.2.2 the "algorithm" value in
// subsequent Request Authorization header must be set to whatever
// was supplied in the WWW-Authenticate Response header. This
// implementation always returns an algorithm in WWW-Authenticate
// header, however there seems to be broken clients in the wild
// which do not set the algorithm. Assume the unset algorithm in
// Authorization header to be equal to MD5.

// Check if the requested URI matches auth header

// We allow auth["uri"] to be a full path prefix of request-uri
// for some reason lost in history, which is probably wrong, but
// used to be like that for quite some time
// (https://tools.ietf.org/html/rfc2617#section-3.2.2 explicitly
// says that auth["uri"] is the request-uri).
//
// TODO: make an option to allow only strict checking.

// At this point crypto checks are completed and validated.
// Now check if the session is valid.

// Default values for ClientCacheSize and ClientCacheTolerance for DigestAuth
const (
	DefaultClientCacheSize      = 1000
	DefaultClientCacheTolerance = 100
)

// Wrap returns an http.HandlerFunc wraps AuthenticatedHandlerFunc
// with this DigestAuth authentication checks. Once the request
// contains valid credentials, it calls wrapped
// AuthenticatedHandlerFunc.
//
// Deprecated: new code should use NewContext instead.
func (da *DigestAuth) Wrap(wrapped AuthenticatedHandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// JustCheck returns a new http.HandlerFunc, which requires
// DigestAuth to successfully authenticate a user before calling
// wrapped http.HandlerFunc.
//
// Authenticated Username is passed as an extra
// X-Authenticated-Username header to the wrapped HandlerFunc.
func (da *DigestAuth) JustCheck(wrapped http.HandlerFunc) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// NewContext returns a context carrying authentication information for the request.
func (da *DigestAuth) NewContext(ctx context.Context, r *http.Request) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// return back digest WWW-Authenticate header

// NewDigestAuthenticator generates a new DigestAuth object
func NewDigestAuthenticator(realm string, secrets SecretProvider) *DigestAuth {
	_ = "STUB: not implemented"
	return nil
}
