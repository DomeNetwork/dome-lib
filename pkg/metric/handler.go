package metric

import (
	"github.com/VictoriaMetrics/metrics"
	"github.com/domenetwork/dome-lib/pkg/log"
	"github.com/gin-gonic/gin"
)

// ReportMetrics godoc
// @Summary Prometheus metrics.
// @Schemes
// @Description Capture route for Prometheus metrics provided by the internal metric client.
// @Tags depot
// @Accept json
// @Produce json,plain
// @Success 200 {object} interface{}
// @Router /metrics [get]
func ReportMetrics(c *gin.Context) {
	log.D("metric", "get")
	metrics.WritePrometheus(c.Writer, false)
}
