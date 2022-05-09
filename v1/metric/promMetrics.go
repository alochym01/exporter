package metric

import "github.com/prometheus/client_golang/prometheus"

var (
	// SysState => System Health Metric
	SysState = prometheus.NewDesc(
		"system_status",
		"system_status {0: OK, 1: Warning, 2: Critical}",
		[]string{"sku", "serialnumber", "model"},
		nil,
	)
)
