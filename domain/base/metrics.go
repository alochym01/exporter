package base

import "github.com/prometheus/client_golang/prometheus"

var (
	// SysState => System Health Metric
	SysState = prometheus.NewDesc(
		"system_status",
		"system_status {0: OK, 1: Warning, 2: Critical}",
		[]string{"sku", "serialnumber"},
		nil,
	)

	// SysStorageDisk => System Storage Metric
	SysStorageDisk = prometheus.NewDesc(
		"storage_drive_ssd_endurance",
		"storage_drive_ssd_endurance {100: OK, 50: Warning, 20: Critical}",
		[]string{"id", "capacity", "interface_type", "media_type"},
		nil,
	)
	// SysStorageDisk => System Storage Metric
	SysStorageStatus = prometheus.NewDesc(
		"storage_status",
		"storage_status {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)
	// SysEthernetInterface => System Storage Metric
	SysEthernetInterface = prometheus.NewDesc(
		"ethernet_port",
		"ethernet_port {0: LinkUp, 2: LinkDown}",
		[]string{"id", "mac", "speed"},
		nil,
	)
	// ChasPower => Chassis Power Metric
	ChasPower = prometheus.NewDesc(
		"power_consumed",
		"power_consumed of server",
		[]string{},
		nil,
	)
	ChasPowerStatus = prometheus.NewDesc(
		"power_status",
		"power_status of server {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)

	ChasFansStatus = prometheus.NewDesc(
		"fan_status",
		"fan_status of server {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)
	ChasTemperatureStatus = prometheus.NewDesc(
		"temperature_status",
		"temperature_status of server {0: OK, 1: Warning, 2: Critical}",
		[]string{},
		nil,
	)
	// ChasNetworkStatus = prometheus.NewDesc(
	// 	"Temperature",
	// 	"Temperature of server",
	// 	[]string{},
	// 	nil,
	// )
)
