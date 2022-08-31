package fetch

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/domenetwork/dome-lib/pkg/cfg"
	"github.com/domenetwork/dome-lib/pkg/common"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/domenetwork/dome-lib/pkg/valid"
)

// Client is the internal HTTP client for interacting with the RESTful endpoints provided by
// each of the DOME services.
type Client struct {
	host   string
	http   *http.Client
	signer func([]byte) ([]byte, error)
	token  string
}

// NewClient will return a new Client instance.
func NewClient() (cli *Client) {
	cli = &Client{
		host: cfg.Str("api.host"),
		http: http.DefaultClient,
	}
	return
}

// Provides the mechanisms for authorization of requests when interacting
// with the backend API services.
func (cli *Client) auth(req *http.Request, data []byte) (err error) {
	// Setup the etag and time for signing.
	hash := sha256.Sum256(data)
	// log.D("fetch", "do", "hash", hash)
	eTag := hex.EncodeToString(hash[:])
	// log.D("fetch", "do", "eTag", eTag)
	unix := common.Unix()
	// log.D("fetch", "do", "unix", unix)

	unixBytes := []byte(fmt.Sprintf("%d", unix))
	log.D("fetch", "do", "unix", unixBytes)

	signable := append(unixBytes, hash[:]...)
	// log.D("fetch", "do", "singable", signable)
	hash = sha256.Sum256(signable)
	log.D("fetch", "do", "hash", hash[:])

	var sig []byte
	if sig, err = cli.signer(hash[:]); err != nil {
		log.E("fetch", "do", "signer", err)
		return
	}
	log.D("fetch", "do", "sig", sig)
	sigHex := hex.EncodeToString(sig)
	// log.D("fetch", "do", "sig hex", sigHex)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("ETag", eTag)
	req.Header.Set("Last-Modified", fmt.Sprintf("%d", unix))
	req.Header.Set("User-Agent", "DOME SDK WASM v1")
	req.Header.Set("X-Auth-Sig", sigHex)

	// If the token is a user's handle, same format as email, then set the From
	// header otherwise set the Bearer token in the Authorization header.
	if err = valid.UUIDv4(cli.token); err != nil {
		handle := fmt.Sprintf(
			"%s@%s",
			cfg.Str("user.username"), cfg.Str("user.domain"),
		)
		req.Header.Set("From", handle)
	} else {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cli.token))
	}
	err = nil
	// log.D("fetch", "auth", "headers", req.Header)
	return
}

func (cli *Client) do(method, route string, body interface{}) (result interface{}, err error) {
	log.D("fetch", "do", method, route, body)
	buf := new(bytes.Buffer)
	if body != nil {
		if err = json.NewEncoder(buf).Encode(body); err != nil {
			log.E("fetch", "do", "body json", buf.String(), err)
			return
		}
	}
	// log.D("fetch", "do", "json", buf.String())

	timeout := 5 * time.Second
	// log.D("fetch", "do", "timeout", timeout)
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var req *http.Request
	url := fmt.Sprintf("%s%s", cli.host, route)
	// log.D("fetch", "do", "url", url)
	if req, err = http.NewRequestWithContext(ctx, method, url, buf); err != nil {
		return
	}

	// log.D("fetch", "do", "auth")
	if err = cli.auth(req, buf.Bytes()); err != nil {
		log.E("fetch", "do", "auth", err)
		return
	}

	// log.D("fetch", "do", "do")
	var res *http.Response
	if res, err = cli.http.Do(req); err != nil {
		return
	}
	defer res.Body.Close()

	// log.D("fetch", "do", "response", "json")
	tmp := map[string]interface{}{}
	if err = json.NewDecoder(res.Body).Decode(&tmp); err != nil {
		return
	}
	log.D("fetch", "do", "response", tmp)
	if tmp["result"] != nil {
		result = tmp["result"]
	}
	return
}

// Auth will set the provided token for future request usage.  The token will be attached to
// Authorization header as a Bearer token.
func (cli *Client) Auth(token string) {
	cli.token = token
}

// Delete make a POST request with the provided body to the provided host and route.
func (cli *Client) Delete(host, route string, body interface{}) (err error) {
	url := fmt.Sprintf("%s%s", host, route)
	_, err = cli.do("DELETE", url, body)
	return
}

// Get makes a GET request to the provided host and route.
func (cli *Client) Get(host, route string) (result interface{}, err error) {
	url := fmt.Sprintf("%s%s", host, route)
	result, err = cli.do("GET", url, nil)
	return
}

// Post make a POST request with the provided body to the provided host and route.
func (cli *Client) Post(host, route string, body interface{}) (result interface{}, err error) {
	url := fmt.Sprintf("%s%s", host, route)
	result, err = cli.do("POST", url, body)
	return
}

// Put make a PUT request with the provided body to the provided host and route.
func (cli *Client) Put(host, route string, body interface{}) (result interface{}, err error) {
	url := fmt.Sprintf("%s%s", host, route)
	result, err = cli.do("PUT", url, body)
	return
}

// Signer will attach a callback to be used by the client for signing or data.
func (cli *Client) Signer(signer func([]byte) ([]byte, error)) {
	cli.signer = signer
}
