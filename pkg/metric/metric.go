package metric

import (
	"fmt"
	"strings"

	"github.com/VictoriaMetrics/metrics"
	"github.com/domenetwork/dome-lib/pkg/log"
)

// Metric wrapper to record metrics.
type Metric struct {
	groups []string
}

// New returns a new Metric reference.
func New() (m *Metric) {
	m = &Metric{
		groups: make([]string, 0),
	}
	return
}

func (m *Metric) prefix() string {
	return strings.Join(m.groups, "_")
}

// Call will increment a method call in the platform.
func (m *Metric) Call(method, key, value string) {
	log.D("metric", "call", method, key, value)
	name := fmt.Sprintf("%s_%s{%s=\"%s\"}", m.prefix(), method, key, value)
	metrics.GetOrCreateCounter(name).Inc()
}

// Inc increments the label of the provided label.  If the counter does
// not exist in the context then it will be created for future reuse.
func (m *Metric) Inc(label string) {
	log.D("metric", "inc", label)
	name := fmt.Sprintf("%s_%s", m.prefix(), label)
	metrics.GetOrCreateCounter(name).Inc()
}

// SubGroup returns an instance of the metric object that prefixes
// the supplied group to the method label names.
func (m *Metric) Sub(group string) *Metric {
	mm := &Metric{
		groups: m.groups,
	}
	mm.groups = append(mm.groups, group)
	return mm
}
