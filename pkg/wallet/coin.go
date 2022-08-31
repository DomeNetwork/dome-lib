package wallet

import (
	"math/big"
)

// Coin provides an interface for all supported Coins in the
// DOME platform.  Each coin must provide these methods in order
// to be compatible for usage with the wallet and DOME platform.
type Coin interface {
	GetAccount() *Account
	GetAccounts() []*Account
	GetAddress(*Account) string
	GetBalance() (*big.Int, error)
	GetGas() (*big.Int, error)
	GetKey(*Wallet) (*Account, error)
	GetPath() string
	GetName() string
	GetSymbol() string
	Load(*Wallet) error
	SendTX(string, *big.Int, []byte) (string, error)
	String() string
	Subscribe(chan interface{}) error
	Unsubscribe() error
}
