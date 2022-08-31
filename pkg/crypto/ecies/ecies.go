package ecies

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/ethereum/go-ethereum/crypto/ecies"
)

// Decrypt the cipher with the provided ECDSA private key using ECIES or return error.
func Decrypt(prv *ecdsa.PrivateKey, cipher []byte) (data []byte, err error) {
	pk := ecies.ImportECDSA(prv)
	data, err = pk.Decrypt(cipher[:], nil, nil)
	return
}

// Encrypt the plain data using the provided ECDSA public key and return the cipher
// using ECIES or return an error.
func Encrypt(pub *ecdsa.PublicKey, data []byte) (cipher []byte, err error) {
	pk := ecies.ImportECDSAPublic(pub)
	cipher, err = ecies.Encrypt(rand.Reader, pk, data[:], nil, nil)
	return
}
