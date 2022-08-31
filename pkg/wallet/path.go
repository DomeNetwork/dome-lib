package wallet

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/hdkeychain"
)

// Path is the derivation path for a key.
type Path []uint32

// ParsePath will take in a provided path string and return the Path
// type or an error.
func ParsePath(path string) (p Path, err error) {
	p = Path{0, 0, 0, 0, 0}

	var hard bool
	var u64 uint64
	path = strings.TrimPrefix(path, "m/")
	parts := strings.Split(path, "/")
	for i, part := range parts {
		hard = strings.Contains(part, "'")
		part = strings.TrimSuffix(part, "'")

		if u64, err = strconv.ParseUint(part, 10, 32); err != nil {
			return
		}

		if hard {
			u64 += hdkeychain.HardenedKeyStart
		}

		p[i] = uint32(u64)
	}
	return
}

// Equal checks to see if the provided path matches the current path.
func (p Path) Equal(t Path) bool {
	for i := range p {
		if p[i] != t[i] {
			return false
		}
	}
	return true
}

// String is the path as a string that matches BIP32 and BIP44.
func (p Path) String() string {
	var a string
	s := "m"
	for _, n := range p {
		a = ""
		if n >= hdkeychain.HardenedKeyStart {
			a = "'"
			n -= hdkeychain.HardenedKeyStart
		}
		s += fmt.Sprintf("/%d%s", n, a)
	}
	return s
}
