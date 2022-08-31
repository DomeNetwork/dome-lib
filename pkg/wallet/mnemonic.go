package wallet

import (
	"github.com/tyler-smith/go-bip39"
)

// NewMnemonic will generate a list of words or return an error.
// The current entropy is 256 bits in size.
func NewMnemonic() (words string, err error) {
	var entropy []byte
	if entropy, err = bip39.NewEntropy(256); err != nil {
		return
	}

	words, err = bip39.NewMnemonic(entropy)
	return
}

// MnemonicFromSeed uses a provided seed to generate a mnemonic word
// phrase or an error.
func MnemonicFromSeed(seed []byte) (words string, err error) {
	words, err = bip39.NewMnemonic(seed)
	return
}

// SeedFromMnemonic will use the provided mnemonic words to generate
// a seed or return an error.
func SeedFromMnemonic(words string) (seed []byte, err error) {
	seed, err = bip39.EntropyFromMnemonic(words)
	return
}
