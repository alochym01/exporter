package dell

import (
	"fmt"

	"github.com/alochym01/exporter/v1/metric"
	"github.com/alochym01/exporter/v1/service/dell"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	server string
	svc    dell.Service
}

// Describe a description of metrics
func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	// fmt.Println(m.server)
	ch <- metric.SysState
	ch <- metric.ChasPower
}

// Collect return a metric with all desc value and metric value
func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	// ComputerSystem Start
	url := m.server + computerSystemURL
	sys, err := m.svc.ComputerSystem(url)
	if err != nil {
		fmt.Println(err)
		// ComputerSystem set Warning for Timeout
		// ch <- prometheus.MustNewConstMetric(metric.SysState, prometheus.GaugeValue, 1, "", "", "")
		return
	}
	ch <- prometheus.MustNewConstMetric(metric.SysState, prometheus.GaugeValue, sys.StatusToNumber(), sys.SKU, sys.SerialNumber, sys.Model)
	// b, err := json.MarshalIndent(sys, "", "   ")
	// fmt.Println(string(b))
	// ComputerSystem End

	// // ComputerSystem Storage Status start
	// ch <- prometheus.MustNewConstMetric(metric.SysStorageStatus, prometheus.GaugeValue, sys.Oem.StorageStatus())
	// // ComputerSystem Storage Status end

	// // ComputerSystem Storage Status start
	// ch <- prometheus.MustNewConstMetric(metric.SysFansStatus, prometheus.GaugeValue, sys.Oem.FansStatus())
	// // ComputerSystem Storage Status end

	// // ComputerSystem Storage Status start
	// ch <- prometheus.MustNewConstMetric(metric.SysPowerStatus, prometheus.GaugeValue, sys.Oem.PowerSupplyStatus())
	// // ComputerSystem Storage Status end

	// // ComputerSystem Storage Status start
	// ch <- prometheus.MustNewConstMetric(metric.SysTemperatureStatus, prometheus.GaugeValue, sys.Oem.TemperatureStatus())
	// // ComputerSystem Storage Status end

	// Storage Disk start
	storageurl := m.server + storageURL
	m.svc.StorageWithChannel(storageurl, m.server, ch)
	// Storage Disk end

	// Power start
	powerurl := m.server + powerURL
	m.svc.PowerWithChannel(powerurl, ch)
	// Power end
}

// NewMetrics return a Metrics struct
func NewMetrics(svr string, s dell.Service) Metrics {
	return Metrics{
		server: svr,
		svc:    s,
	}
}
