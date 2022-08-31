package wallet

import (
	"strings"
	"testing"

	"github.com/tyler-smith/go-bip39"
)

var (
	seed  []byte
	words string
)

func TestNewMnemonic(t *testing.T) {
	var err error
	if words, err = NewMnemonic(); err != nil {
		t.Error(err)
	}

	if !bip39.IsMnemonicValid(words) {
		t.Errorf("mnemonic is not valid: %s", words)
	}
}

func TestSeedFromMnemonic(t *testing.T) {
	var err error
	if seed, err = SeedFromMnemonic(words); err != nil {
		t.Error(err)
	}
}

func TestMnemonicFromSeed(t *testing.T) {
	testWords, err := MnemonicFromSeed(seed)
	if err != nil {
		t.Error(err)
	}

	if strings.Compare(words, testWords) != 0 {
		t.Errorf("mnemonic from seed does not match: `%s` != `%s`", words, testWords)
	}
}
