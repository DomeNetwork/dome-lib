package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// IClient provides an interface for internal Ethereum client.
type IClient interface {
	Balance(common.Address) (*big.Int, error)
	ChainID() (*big.Int, error)
	GasLimit(common.Address, []byte) (uint64, error)
	GasPrice() (*big.Int, error)
	Nonce(common.Address) (uint64, error)
	SendTX(*types.Transaction) error
	Subscribe(chan interface{}) error
	Unsubscribe() error
}
