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

	SysStorageStatus = prometheus.NewDesc(
		"storage_status",
		"storage_status {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)

	SysPowerStatus = prometheus.NewDesc(
		"power_status",
		"power_status of server {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)

	SysFansStatus = prometheus.NewDesc(
		"fan_status",
		"fan_status of server {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)
	SysTemperatureStatus = prometheus.NewDesc(
		"temperature_status",
		"temperature_status of server {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)
)
