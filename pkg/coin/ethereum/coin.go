package ethereum

import (
	"crypto/ecdsa"
	"errors"
	"math/big"

	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/eth"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/domenetwork/dome-lib/pkg/wallet"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

var _ wallet.Coin = &Ethereum{}

type Ethereum struct {
	accts    map[string]*wallet.Account
	baseAcct *wallet.Account
	cli      eth.IClient
	path     wallet.Path

	// Listen to incoming information.
	headers chan *types.Header
	sub     ethereum.Subscription
}

func New() (coin *Ethereum, err error) {
	coin = &Ethereum{
		accts: make(map[string]*wallet.Account),
		path: wallet.Path{
			hdkeychain.HardenedKeyStart + 44,
			hdkeychain.HardenedKeyStart + CoinType,
			hdkeychain.HardenedKeyStart,
			0,
			0,
		},
	}

	url := cfg.Str("coin.ethereum.ws")
	log.D("ethereum", "new", "url", url)

	if coin.cli, err = eth.NewClient(url); err != nil {
		log.E("ethereum", "new", "cli", err)
		return
	}
	return
}

func (coin *Ethereum) GetAccount() *wallet.Account {
	return coin.baseAcct
}

func (coin *Ethereum) GetAccounts() (accts []*wallet.Account) {
	accts = make([]*wallet.Account, len(coin.accts))
	i := 0
	for _, acct := range coin.accts {
		accts[i] = acct
	}
	return
}

func (coin *Ethereum) GetAddress(acct *wallet.Account) string {
	prv, _ := acct.PrivateKey()
	pub := prv.Public().(*ecdsa.PublicKey)
	return crypto.PubkeyToAddress(*pub).Hex()
}

func (coin *Ethereum) GetBalance() (total *big.Int, err error) {
	total = big.NewInt(0)

	var addr common.Address
	var prv *ecdsa.PrivateKey
	if prv, err = coin.baseAcct.PrivateKey(); err != nil {
		return
	}

	addr = crypto.PubkeyToAddress(prv.PublicKey)
	if total, err = coin.cli.Balance(addr); err != nil {
		return
	}
	return
}

func (coin *Ethereum) GetGas() (gas *big.Int, err error) {
	gas, err = coin.cli.GasPrice()
	return
}

func (coin *Ethereum) GetKey(w *wallet.Wallet) (acct *wallet.Account, err error) {
	path := coin.path[:]
	path[4]++

	if acct, err = w.Derive(path); err != nil {
		return
	}

	coin.accts[acct.Path.String()] = acct
	coin.path[4] = path[4]
	return
}

func (coin *Ethereum) GetName() string {
	return CoinName
}

func (coin *Ethereum) GetPath() string {
	return coin.path.String()
}

func (coin *Ethereum) GetSymbol() string {
	return CoinSymbol
}

func (coin *Ethereum) Load(w *wallet.Wallet) (err error) {
	coin.baseAcct, err = coin.GetKey(w)
	return
}

func (coin *Ethereum) SendTX(to string, amount *big.Int, data []byte) (tx string, err error) {
	log.D("coin", "ethereum", "send tx", to, amount, data)
	toAddr := common.HexToAddress(to)
	log.D("coin", "ethereum", "send tx", "to", toAddr)

	var prv *ecdsa.PrivateKey
	if prv, err = coin.baseAcct.PrivateKey(); err != nil {
		log.E("coin", "ethereum", "send tx", "private key", err)
		return
	}
	prv.Curve = crypto.S256()

	pub := prv.Public().(*ecdsa.PublicKey)
	fromAddr := crypto.PubkeyToAddress(*pub)
	log.D("coin", "ethereum", "send tx", "from", fromAddr)

	var nonce uint64
	if nonce, err = coin.cli.Nonce(fromAddr); err != nil {
		log.E("coin", "ethereum", "send tx", "pending nonce", err)
		return
	}
	log.D("coin", "ethereum", "send tx", "nonce", nonce)

	var gasPrice *big.Int
	if gasPrice, err = coin.GetGas(); err != nil {
		log.E("coin", "ethereum", "send tx", "gas price", err)
		return
	}
	log.D("coin", "ethereum", "send tx", "gas price", gasPrice)

	var gasLimit uint64
	if gasLimit, err = coin.cli.GasLimit(toAddr, data); err != nil {
		log.E("coin", "ethereum", "send tx", "gas limit", err)
		return
	}
	log.D("coin", "ethereum", "send tx", "gas limit", gasLimit)
	newTX := types.NewTransaction(nonce, toAddr, amount, gasLimit, gasPrice, data)

	var chainID *big.Int
	if chainID, err = coin.cli.ChainID(); err != nil {
		log.E("coin", "ethereum", "send tx", "chain ID", err)
		return
	}
	log.D("coin", "ethereum", "send tx", "chain ID", chainID)

	var signedTX *types.Transaction
	if signedTX, err = types.SignTx(newTX, types.NewEIP155Signer(chainID), prv); err != nil {
		log.E("coin", "ethereum", "send tx", "sign TX", err)
		return
	}

	if err = coin.cli.SendTX(signedTX); err != nil {
		log.E("coin", "ethereum", "send tx", "send TX", err)
		return
	}

	tx = signedTX.Hash().Hex()
	log.D("coin", "ethereum", "send tx", "tx", tx)
	return
}

func (coin *Ethereum) String() string {
	return coin.GetName()
}

func (coin *Ethereum) Subscribe(ch chan interface{}) (err error) {
	log.D("coin", "ethereum", "subscribe", "ch", ch)
	if coin.sub != nil {
		err = errors.New("already subscribed, unsubscribe first")
		return
	}

	if err = coin.cli.Subscribe(ch); err != nil {
		log.E("coin", "ethereum", "subscribe", err)
		return
	}
	return
}

func (coin *Ethereum) Unsubscribe() (err error) {
	log.D("coin", "ethereum", "unsubscribe")
	if err = coin.cli.Unsubscribe(); err != nil {
		log.E("coin", "ethereum", "unsubscribe", err)
		return
	}
	return
}
