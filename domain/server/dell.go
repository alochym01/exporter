package server

import (
	"github.com/alochym01/exporter/domain/dell"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// DellHandler ...
type DELLHandler struct{}

// Metric ...
func (handler DELLHandler) Metric(c *gin.Context) {
	// Set Host get from Request
	// redfish.ClientHPE.Server = c.Query("host")

	// Register Server Dell Metrics
	// Using custom registry
	registry := prometheus.NewRegistry()

	server := dell.NewMetrics(c.Query("host"))
	// server := hpe.NewMetricsV1()

	registry.MustRegister(server)

	// Make promhttp response to Request
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(c.Writer, c.Request)
}

// NewDELLHandler return a DELLHandler struct
func NewDELLHandler() DELLHandler {
	return DELLHandler{}
}
