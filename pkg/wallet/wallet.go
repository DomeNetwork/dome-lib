package wallet

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/domenetwork/dome-lib/pkg/log"
)

// Wallet provides a BIP32 HD wallet for usage with coins.
type Wallet struct {
	coins     map[string]Coin
	masterKey *hdkeychain.ExtendedKey
	seed      []byte
}

// New will return an instance of a wallet that must be loaded
// before usage because the seed is not properly set at this step.
func New() (w *Wallet) {
	log.D("wallet", "new")
	w = &Wallet{
		coins:     make(map[string]Coin),
		masterKey: nil,
		seed:      make([]byte, 0),
	}
	return
}

// AddCoins will register coins with the wallet for usage.
func (w *Wallet) AddCoins(coins ...Coin) {
	log.D("wallet", "add coins", coins)
	for _, coin := range coins {
		w.coins[coin.GetName()] = coin
	}
}

// GetCoin returns a coin for the matching provided name.
func (w *Wallet) GetCoin(name string) (coin Coin) {
	log.D("wallet", "get coin", name)
	coin = w.coins[name]
	return
}

// GetCoins returns a list of all supported coins in the wallet.
func (w *Wallet) GetCoins() (coins []Coin) {
	log.D("wallet", "get coins")
	coins = make([]Coin, len(w.coins))
	i := 0
	for _, coin := range w.coins {
		coins[i] = coin
		i++
	}
	return
}

// Derive an account for the provided derivation path or error.
func (w *Wallet) Derive(path Path) (acct *Account, err error) {
	log.D("wallet", "derive", path)
	key := w.masterKey
	for _, p := range path {
		if key, err = key.Derive(p); err != nil {
			return
		}
	}

	acct = &Account{
		Key:  key,
		Path: path,
	}
	return
}

// Load the wallet with the provides seed.  This MUST be called before wallet usage.
func (w *Wallet) Load(seed []byte) (err error) {
	log.D("wallet", "load", seed)
	if w.masterKey, err = hdkeychain.NewMaster(seed, &chaincfg.MainNetParams); err != nil {
		return
	}

	w.seed = seed

	// Load the coins as well.
	for _, coin := range w.coins {
		if err = coin.Load(w); err != nil {
			return
		}
	}
	return
}

// Seed returns the seed for the wallet that is used for encrypted storage and loading.
func (w *Wallet) Seed() []byte {
	log.D("wallet", "seed")
	return w.seed
}
