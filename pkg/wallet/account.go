package wallet

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// Account represents an internal account for a Coin. It contains the
// private key, derivation path, and balance information.
type Account struct {
	Balance *big.Int
	Key     *hdkeychain.ExtendedKey
	Path    Path
}

// PrivateKey will return the ECDSA version of the internal account key.
func (acct *Account) PrivateKey() (prv *ecdsa.PrivateKey, err error) {
	var ecprv *btcec.PrivateKey
	if ecprv, err = acct.Key.ECPrivKey(); err != nil {
		return
	}

	prv = ecprv.ToECDSA()
	return
}
