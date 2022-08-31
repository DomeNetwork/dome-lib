//go:build !js
// +build !js

package eth

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client is a basic wrapper around the Ethereum client.
type Client struct {
	eth *ethclient.Client
}

// NewClient returns a new wrapped client.
func NewClient(url string) (cli *Client, err error) {
	cli = &Client{}

	cli.eth, err = ethclient.Dial(url)
	return
}

// Balance for the given address or error.
func (cli *Client) Balance(addr common.Address) (balance *big.Int, err error) {
	balance, err = cli.eth.BalanceAt(context.Background(), addr, nil)
	return
}

// ChainID for the client for error.
func (cli *Client) ChainID() (chainID *big.Int, err error) {
	chainID, err = cli.eth.ChainID(context.Background())
	return
}

// GasLimit of the transaction and data or error.
func (cli *Client) GasLimit(addr common.Address, data []byte) (gasLimit uint64, err error) {
	msg := ethereum.CallMsg{
		Data: data,
		To:   &addr,
	}
	gasLimit, err = cli.eth.EstimateGas(context.Background(), msg)
	return
}

// GasPrice suggested for current network conditions or error.
func (cli *Client) GasPrice() (gasPrice *big.Int, err error) {
	gasPrice, err = cli.eth.SuggestGasPrice(context.Background())
	return
}

// Nonce of given address or error.
func (cli *Client) Nonce(addr common.Address) (nonce uint64, err error) {
	nonce, err = cli.eth.NonceAt(context.Background(), addr, nil)
	return
}

// SendTX to client network or error.
func (cli *Client) SendTX(signedTX *types.Transaction) (err error) {
	err = cli.eth.SendTransaction(context.Background(), signedTX)
	return
}

// Subscribe not supported.
func (cli *Client) Subscribe(ch chan interface{}) (err error) {
	return
}

// Unsubscribe not supported.
func (cli *Client) Unsubscribe() (err error) {
	return
}
