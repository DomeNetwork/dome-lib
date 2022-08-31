//go:build js
// +build js

package eth

import (
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"syscall/js"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

// TODO: add deadline/timeout for calls.

// Client provides a mechanism for requesting using WebSockets to
// an Ethereum RPC node or other service like Infura.  This provides
// direct support for Web (WASM & JS) usage.
// Note: Only a subset of RPC commands are supported at this time.
type Client struct {
	callID   uint64                 // The current call ID.
	callLock *sync.RWMutex          // Lock for managing the call map.
	callMap  map[uint64]chan string // The map of id -> channel calls.

	eth js.Value // Represents the JS WebSocket value.

	subID string           // The subscription id from the RPC node.
	subCH chan interface{} // The incoming subscription channel.
}

// NewClient will return a WASM JS compatible Client.
func NewClient(url string) (cli *Client, err error) {
	cli = &Client{
		callID:   1,
		callLock: new(sync.RWMutex),
		callMap:  make(map[uint64]chan string),
		eth:      js.Global().Get("WebSocket").New(url),
	}

	cli.callLock.Lock() // We lock the client until the WebSocket is open.

	// Attach event listeners to the JS WebSocket object.
	cli.eth.Call("addEventListener", "close", cli.close())
	cli.eth.Call("addEventListener", "error", cli.error())
	cli.eth.Call("addEventListener", "message", cli.message())
	cli.eth.Call("addEventListener", "open", cli.open())
	return
}

// Private

func (cli *Client) close() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		// TODO: cleanup state
		return nil
	})
}

func (cli *Client) error() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		// TODO: cleanup state
		return nil
	})
}

func (cli *Client) message() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		data := make(map[string]interface{})
		obj := args[0].Get("data").String()
		r := strings.NewReader(obj)
		if err := json.NewDecoder(r).Decode(&data); err != nil {
			return err
		}

		method := data["method"]
		if method == "eth_subscription" {
			params := data["params"].(map[string]interface{})
			result := params["result"]
			cli.subID = params["subscription"].(string) // TODO: move out to call response.

			cli.subCH <- result
		} else {
			var x string
			// Handle different types of result values.
			switch data["result"].(type) {
			case string:
				// TODO: move this out so subscription ID can be returned as hex.
				v, err := hexutil.DecodeBig(data["result"].(string))
				if err != nil {
					return err
				}

				x = v.String()
			default:
				x = "true"
			}

			// Get the call ID from the message.
			fid := data["id"].(float64)
			id := uint64(fid)

			// Send our data to the channel and cleanup.
			cli.callLock.RLock()
			defer cli.callLock.RUnlock()
			if ch, ok := cli.callMap[id]; ok {
				ch <- x
				close(cli.callMap[id])
				delete(cli.callMap, id)
			}
		}
		return nil
	})
}

func (cli *Client) open() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		// Once we have connected to the RPC node we should unlock
		// the client for processing of requests.
		cli.callLock.Unlock()
		return nil
	})
}

func (cli *Client) send(method, params string) (ch chan string) {
	// Put a read lock on the client state to update it.
	cli.callLock.RLock()
	defer cli.callLock.RUnlock()

	// We setup a channel here and add it to the call map so the consumer
	// can wait for a response from the server.
	ch = make(chan string, 1)
	cli.callMap[cli.callID] = ch

	// A manual string is defined here instead of another type because we need
	// the parameters to have quotes escaped and simple json.Encoding will not
	// do that for us.
	// TODO: is there a way to do that with the json encoder?
	s := fmt.Sprintf(
		"{ \"jsonrpc\": \"2.0\", \"id\": %d, \"method\": \"%s\", \"params\": %s }",
		cli.callID, method, params,
	)
	cli.eth.Call("send", s)
	cli.callID++
	return
}

// Public

// Balance will return the balance calculated for the provided address against
// the latest block header.
func (cli *Client) Balance(addr common.Address) (balance *big.Int, err error) {
	s := fmt.Sprintf("[\"%s\", \"%s\"]", addr.Hex(), "latest")
	cb := cli.send("eth_getBalance", s)

	v := <-cb // Wait for a response in the channel.

	var x int64
	if x, err = strconv.ParseInt(v, 10, 64); err != nil {
		return
	}

	balance = big.NewInt(x)
	return
}

// ChainID return the integer representation of the current connected chain.
// The chain ID should match the chain connected to during initialization.
func (cli *Client) ChainID() (chain *big.Int, err error) {
	cb := cli.send("eth_chainId", "[]")

	v := <-cb // Wait for a response in the channel.

	var x int64
	if x, err = strconv.ParseInt(v, 10, 64); err != nil {
		return
	}

	chain = big.NewInt(x)
	return
}

// GasLimit returns the estimate for the address and data of the a transaction.
func (cli *Client) GasLimit(addr common.Address, data []byte) (gasLimit uint64, err error) {
	s := fmt.Sprintf("[{\"from\":, \"%s\", \"data\":, \"%s\"}]", addr.Hex(), hexutil.Encode(data))
	cb := cli.send("eth_estimateGas", s)

	v := <-cb // Wait for a response in the channel.

	gasLimit, err = strconv.ParseUint(v, 10, 64)
	return
}

// GapPrice will return the suggested gas price to use for a transaction on the network.
func (cli *Client) GasPrice() (gasPrice *big.Int, err error) {
	cb := cli.send("eth_gasPrice", "[]")

	v := <-cb // Wait for a response in the channel.

	var x int64
	if x, err = strconv.ParseInt(v, 10, 64); err != nil {
		return
	}

	gasPrice = big.NewInt(x)
	return
}

// Nonce returns the provided address's nonce value for usage in sending transactions.
func (cli *Client) Nonce(addr common.Address) (nonce uint64, err error) {
	s := fmt.Sprintf("[\"%s\"]", addr.Hex())
	cb := cli.send("eth_getTransactionCount", s)

	v := <-cb // Wait for a response in the channel.

	nonce, err = strconv.ParseUint(v, 10, 64)
	return
}

// SendTX will send a signed TX to the RPC node for submission to the network.
func (cli *Client) SendTX(signedTX *types.Transaction) (err error) {
	var b []byte
	if b, err = signedTX.MarshalJSON(); err != nil {
		return
	}

	s := fmt.Sprintf("[\"%s\"]", b)
	cli.send("eth_sendTransaction", s)
	return
}

// Subscribe to new block headers.
func (cli *Client) Subscribe(ch chan interface{}) (err error) {
	cli.subCH = ch
	cli.send("eth_subscribe", "[\"newHeads\"]")
	return
}

// Unsubscribe from the current subscription.
func (cli *Client) Unsubscribe() (err error) {
	s := fmt.Sprintf("[\"%s\"]", cli.subID)
	cli.send("eth_unsubscribe", s)

	close(cli.subCH)
	cli.subCH = nil
	cli.subID = ""
	return
}
