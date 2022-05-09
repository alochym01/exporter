package dell

import (
	"fmt"

	"github.com/alochym01/exporter/v1/service/dell"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Handler struct {
	svc dell.Service
}

func (d Handler) Metrics(c *gin.Context) {
	// Using custom registry
	registry := prometheus.NewRegistry()

	host := c.Query("host")
	fmt.Println(host)
	server := NewMetrics(host, d.svc)

	registry.MustRegister(server)

	// Make promhttp response to Request
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(c.Writer, c.Request)
}

// NewHandler return a DELLHandler struct
func NewHandler(s dell.Service) Handler {
	return Handler{
		svc: s,
	}
}
