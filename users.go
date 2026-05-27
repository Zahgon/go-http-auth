package auth

import (
	"os"
	"sync"
)

// SecretProvider is used by authenticators. Takes user name and realm
// as an argument, returns secret required for authentication (HA1 for
// digest authentication, properly encrypted password for basic).
//
// Returning an empty string means failing the authentication.
type SecretProvider func(user, realm string) string

// File handles automatic file reloading on changes.
type File struct {
	Path string
	Info os.FileInfo
	/* must be set in inherited types during initialization */
	Reload func()
	mu     sync.Mutex
}

// ReloadIfNeeded checks file Stat and calls Reload() if any changes
// were detected. File mutex is Locked for the duration of Reload()
// call.
//
// This function will panic() if Stat fails.
func (f *File) ReloadIfNeeded() { _ = "STUB: not implemented"; return }

// HtdigestFile is a File holding htdigest authentication data.
type HtdigestFile struct {
	// File is used for automatic reloading of the authentication data.
	File
	// Users is a map of realms to users to HA1 digests.
	Users map[string]map[string]string
	mu    sync.RWMutex
}

func reloadHTDigest(hf *HtdigestFile) { _ = "STUB: not implemented"; return }

// HtdigestFileProvider is a SecretProvider implementation based on
// htdigest-formated files. It will automatically reload htdigest file
// on changes. It panics on syntax errors in htdigest files.
func HtdigestFileProvider(filename string) SecretProvider {
	_ = "STUB: not implemented"
	return *new(SecretProvider)
}

// HtpasswdFile is a File holding basic authentication data.
type HtpasswdFile struct {
	// File is used for automatic reloading of the authentication data.
	File
	// Users is a map of users to their secrets (salted encrypted
	// passwords).
	Users map[string]string
	mu    sync.RWMutex
}

func reloadHTPasswd(h *HtpasswdFile) { _ = "STUB: not implemented"; return }

// HtpasswdFileProvider is a SecretProvider implementation based on
// htpasswd-formated files. It will automatically reload htpasswd file
// on changes. It panics on syntax errors in htpasswd files. Realm
// argument of the SecretProvider is ignored.
func HtpasswdFileProvider(filename string) SecretProvider {
	_ = "STUB: not implemented"
	return *new(SecretProvider)
}
