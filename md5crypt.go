package auth

const itoa64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var md5CryptSwaps = [16]int{12, 6, 0, 13, 7, 1, 14, 8, 2, 15, 9, 3, 5, 10, 4, 11}

// MD5Crypt is the MD5 password crypt implementation.
func MD5Crypt(password, salt, magic []byte) []byte { _ = "STUB: not implemented"; return nil }
