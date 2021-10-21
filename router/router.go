package router

import (
	"strings"

	"github.com/alochym01/exporter/domain/server"
	"github.com/gin-gonic/gin"
)

// Router ...
// func Router(db *sql.DB) *gin.Engine {
func Router(ginMode string) *gin.Engine {
	// redfish.ClientDELL = redfish.NewAPIClient("root", "calvin")
	// redfish.ClientHPE = redfish.NewAPIClient("username", "password")

	router := gin.Default()

	if strings.Contains(ginMode, "release") {
		gin.SetMode(gin.ReleaseMode)
	}

	hpeHandler := server.NewHPEHandler()
	router.GET("/metrics/hpe", hpeHandler.Metric)

	dellHandler := server.NewDELLHandler()
	router.GET("/metrics/dell", dellHandler.Metric)

	return router
}
