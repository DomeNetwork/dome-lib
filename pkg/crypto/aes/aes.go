package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// Set up AES with Galois/Counter Mode (AES-GCM).
func aesGCM(data []byte) (gcm cipher.AEAD, err error) {
	key := sha256.Sum256(data)

	var block cipher.Block
	block, err = aes.NewCipher(key[:])
	if err != nil {
		return
	}

	gcm, err = cipher.NewGCM(block)
	return
}

// Decrypt uses the given secret string to decrypt the cipher byte slice provided
// and returns the plain text bytes back or an error.
func Decrypt(secret string, encrypted []byte) (decrypted []byte, err error) {
	var gcm cipher.AEAD
	password := []byte(secret)
	if gcm, err = aesGCM(password); err != nil {
		return
	}

	nonce, encrypted := encrypted[:gcm.NonceSize()], encrypted[gcm.NonceSize():]

	decrypted, err = gcm.Open(nil, nonce, encrypted, nil)
	return
}

// Encrypt uses the provided secret string to encrypt the provided byte slice and return
// the cipher back as a byte slice or an error.
func Encrypt(secret string, decrypted []byte) (encrypted []byte, err error) {
	var gcm cipher.AEAD
	password := []byte(secret)
	if gcm, err = aesGCM(password); err != nil {
		return
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}

	encrypted = gcm.Seal(nonce, nonce, decrypted, nil)
	return
}
