package router

import (
	"strconv"
	"strings"
	"time"

	dellHandler "github.com/alochym01/exporter/v1/handler/dell"
	dellService "github.com/alochym01/exporter/v1/service/dell"
	"github.com/alochym01/exporter/v1/storage"
	"github.com/gin-gonic/gin"
)

// Router ...
// func Router(db *sql.DB) *gin.Engine {
func Router(ginMode string, timeOut string) *gin.Engine {
	// redfish.ClientDELL = redfish.NewAPIClient("root", "calvin")
	// redfish.ClientHPE = redfish.NewAPIClient("username", "password")

	router := gin.Default()

	if strings.Contains(ginMode, "release") {
		gin.SetMode(gin.ReleaseMode)
	}

	// Dell Handler Metrics
	t, _ := strconv.ParseInt(timeOut, 10, 64)
	dStore := storage.NewClient("root", "calvin", time.Duration(t))
	dService := dellService.NewService(dStore)
	dHandler := dellHandler.NewHandler(dService)
	router.GET("/metrics/dell", dHandler.Metrics)

	return router
}
