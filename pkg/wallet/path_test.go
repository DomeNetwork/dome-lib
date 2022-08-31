package wallet

import (
	"fmt"
	"testing"

	"github.com/btcsuite/btcutil/hdkeychain"
)

func TestParsePath(t *testing.T) {
	p := Path{44, 60, 0, 0, 0}
	tp := "m/44/60/0/0/0"

	ts, err := ParsePath(tp)
	if err != nil {
		t.Error(err)
	}

	if !p.Equal(ts) {
		t.Errorf("parsed path does not match expected: %v != %v", p, t)
	}

	p[0] += hdkeychain.HardenedKeyStart
	p[1] += hdkeychain.HardenedKeyStart
	p[2] += hdkeychain.HardenedKeyStart
	tp = "m/44'/60'/0'/0/0"

	ts, err = ParsePath(tp)
	if err != nil {
		t.Error(err)
	}

	if !p.Equal(ts) {
		t.Errorf("hardened parsed path does not match expected: %v != %v", p, t)
	}
}

func TestPathToString(t *testing.T) {
	p := Path{44, 60, 0, 0, 0}

	s := p.String()
	tp := "m/44/60/0/0/0"
	if s != tp {
		t.Error(fmt.Errorf("path should match `%s` but got `%s`", p.String(), tp))
	}

	for i := 0; i < 3; i++ {
		p[i] += hdkeychain.HardenedKeyStart
	}

	s = p.String()
	tp = "m/44'/60'/0'/0/0"
	if s != tp {
		t.Error(fmt.Errorf("path should match `%s` but got `%s`", p.String(), tp))
	}
}
