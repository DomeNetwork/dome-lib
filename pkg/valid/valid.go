package valid

import (
	"fmt"

	"github.com/domenetwork/dome-lib/pkg/log"
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/tyler-smith/go-bip39"
)

// Address return an error if it is unable to validate the
// provided input against the expected type.
func Address(s string) error {
	log.D("valid", "address", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(32, 32),
		is.Hexadecimal,
	)
}

// Email return an error if it is unable to validate the
// provided input against the expected type.
func Email(s string) error {
	log.D("valid", "email", s)
	return v.Validate(
		s,
		v.Required,
		is.Email,
	)
}

// Entropy return an error if it is unable to validate the
// provided input against the expected type.
func Entropy(s string) error {
	log.D("valid", "entropy", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(32, 32),
		is.Hexadecimal,
	)
}

// Mnemonic return an error if it is unable to validate the
// provided input against the expected type.
func Mnemonic(s string) error {
	log.D("valid", "mnemonic", s)
	if !bip39.IsMnemonicValid(s) {
		return fmt.Errorf("mnemonic is not valid")
	}
	// words := strings.Split(s, " ")

	// return v.Validate(
	// 	words,
	// 	v.Length(12, 12),
	// 	v.Each(v.Required, is.Alpha),
	// )
	return nil
}

// Password return an error if it is unable to validate the
// provided input against the expected type.
func Password(s string) error {
	log.D("valid", "password", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(6, 20),
		is.ASCII,
	)
}

// PrivateKey return an error if it is unable to validate the
// provided input against the expected type.
func PrivateKey(s string) error {
	log.D("valid", "privateKey", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(64, 64),
		is.Hexadecimal,
	)
}

// PublicKey return an error if it is unable to validate the
// provided input against the expected type.
func PublicKey(s string) error {
	log.D("valid", "publicKey", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(32, 32),
		is.Hexadecimal,
	)
}

// Secret return an error if it is unable to validate the
// provided input against the expected type.
func Secret(s string) error {
	log.D("valid", "secret", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(4, 4),
		is.Int,
	)
}

// Username return an error if it is unable to validate the
// provided input against the expected type.
func Username(s string) error {
	log.D("valid", "username", s)
	return v.Validate(
		s,
		v.Required,
		v.Length(3, 20),
		is.Alphanumeric,
	)
}

// UUIDv4 return an error if it is unable to validate the
// provided input against the expected type.
func UUIDv4(s string) error {
	log.D("valid", "uuidv4", s)
	return v.Validate(
		s,
		v.Required,
		is.UUIDv4,
	)
}
