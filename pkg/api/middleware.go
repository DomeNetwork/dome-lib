package api

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"

	"github.com/domenetwork/dome-lib/pkg/common"
	"github.com/domenetwork/dome-lib/pkg/db/kv"
	"github.com/domenetwork/dome-lib/pkg/db/sql"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// AuthGate will short-circuit the request if it is middle
// authorization from the platform.
func AuthGate() func(*gin.Context) {
	return func(c *gin.Context) {
		xauth := c.Request.Header.Get("X-Auth-Check")
		if xauth != "" && xauth == "YES" {
			OK(c, "OK")
			return
		}
	}
}

// AuthSig verifies the provided signature matches the
// users current identity root key.
func AuthSig(db sql.SQL) func(*gin.Context) {
	return func(c *gin.Context) {
		sigHex := c.Request.Header.Get("X-Auth-Sig")
		sig, err := hex.DecodeString(sigHex)
		if err != nil {
			ErrBadRequest(c, err)
			return
		}

		etag := c.Request.Header.Get("ETag")
		etagBytes, err := hex.DecodeString(etag)
		if err != nil {
			ErrBadRequest(c, err)
			return
		}

		user := &common.User{}
		userID, found := c.Get("userID")
		// log.D("api", "middleware", "AuthSig", "userID", userID, found)
		if !found {
			from := c.Request.Header.Get("From")
			// log.D("api", "middleware", "AuthSig", "from", from)
			parts := strings.Split(from, "@")
			// log.D("api", "middleware", "AuthSig", "from parts", parts)

			user.Domain = parts[1]
			user.Username = parts[0]
		} else {
			user.ID = userID.(common.ID)
		}
		if err := db.Get(user); err != nil {
			ErrUnauthorized(c, err)
			return
		}

		pubHex := user.PublicKey
		pubBytes, err := hex.DecodeString(pubHex)
		if err != nil {
			ErrBadGateway(c, err)
			return
		}
		pub, err := crypto.UnmarshalPubkey(pubBytes)
		if err != nil {
			ErrForbidden(c, err)
			return
		}

		unix := c.Request.Header.Get("Last-Modified")
		signable := append([]byte(unix), etagBytes...)
		hash := sha256.Sum256(signable)
		log.D("api", "middleware", "AuthSig", "unix", []byte(unix))
		log.D("api", "middleware", "AuthSig", "hash", hash)
		log.D("api", "middleware", "AuthSig", "sig", sig)

		if !ecdsa.VerifyASN1(pub, hash[:], sig) {
			ErrUnauthorized(c, err)
			return
		}

		c.Next()
	}
}

// AuthToken will verify that the token provided in the Authorization
// header is valid and found in the token store.
func AuthToken(cache kv.KV) func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") == "" {
			c.Next()
			return
		}

		bearer := extractBearerFromHeader(c)

		v, err := cache.Get(bearer)
		if err != nil {
			ErrUnauthorized(c, err)
			return
		}

		userID := v.(string)
		c.Set("userID", common.ID(userID))

		c.Next()
	}
}

// Defaults will install all of the default middleware like CORS, etc.
func Defaults(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"DELETE",
			"GET",
			"OPTIONS",
			"POST",
			"PUT",
		},
		AllowHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
			"ETag",
			"From",
			"Last-Modified",
			"Origin",
			"User-Agent",
			"X-Auth-Sig",
		},
		ExposeHeaders: []string{
			"Content-Length",
			"Content-Type",
			"ETag",
			"From",
			"Last-Modified",
		},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))
}
