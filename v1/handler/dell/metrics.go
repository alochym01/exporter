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
}

// Collect return a metric with all desc value and metric value
func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	url := m.server + computerSystemURL
	sys, err := m.svc.ComputerSystem(url)
	if err != nil {
		fmt.Println(err)
		ch <- prometheus.MustNewConstMetric(metric.SysState, prometheus.GaugeValue, 1, "", "", "")
		return
	}
	ch <- prometheus.MustNewConstMetric(metric.SysState, prometheus.GaugeValue, sys.StatusToNumber(), sys.SKU, sys.SerialNumber, sys.Model)
	// fmt.Println(data)
}

// NewMetrics return a Metrics struct
func NewMetrics(svr string, s dell.Service) Metrics {
	return Metrics{
		server: svr,
		svc:    s,
	}
}
