package bcrypt

import "golang.org/x/crypto/bcrypt"

// Compare the provided password with the stored password
// hash and return an error if they do not match.
func Compare(pwd []byte, hash []byte) (err error) {
	err = bcrypt.CompareHashAndPassword(hash, pwd)
	return
}

// Hash the provided password and return or throw an error.
// The current cost is set to 12.
func Hash(pwd []byte) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword(pwd, 12)
	return
}
