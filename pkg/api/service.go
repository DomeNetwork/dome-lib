package api

import (
	"github.com/domenetwork/dome-lib/pkg/db/kv"
	"github.com/domenetwork/dome-lib/pkg/db/sql"
	"github.com/domenetwork/dome-lib/pkg/metric"
	"github.com/gin-gonic/gin"
)

// Service defines the interface for internal services
// to use.  These are the backend APIs that power the platform.
type Service interface {
	Addr() string
	API() *gin.Engine
	Cache() kv.KV
	Close() error
	DB() sql.SQL
	Metric() *metric.Metric
	Run() error
}
