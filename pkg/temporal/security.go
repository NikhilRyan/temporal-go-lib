package temporal

import (
    "temporal-go-lib/internal/security"
)

// Encrypt encrypts plain text string into cipher text string
func Encrypt(plainText, key string) (string, error) {
    return security.Encrypt(plainText, key)
}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(cipherText, key string) (string, error) {
    return security.Decrypt(cipherText, key)
}
