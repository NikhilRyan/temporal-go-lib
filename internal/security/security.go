package security

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
    "io"
)

// Encrypt encrypts plain text string into cipher text string
func Encrypt(plainText, key string) (string, error) {
    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonce := make([]byte, aesGCM.NonceSize())
    if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

    cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
    return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts cipher text string into plain text string
func Decrypt(cipherText, key string) (string, error) {
    data, err := base64.URLEncoding.DecodeString(cipherText)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher([]byte(key))
    if err != nil {
        return "", err
    }

    aesGCM, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonceSize := aesGCM.NonceSize()
    if len(data) < nonceSize {
        return "", errors.New("ciphertext too short")
    }

    nonce, cipherText := data[:nonceSize], data[nonceSize:]
    plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
    if err != nil {
        return "", err
    }

    return string(plainText), nil
}
