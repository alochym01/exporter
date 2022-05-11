# Server Exporter using iLO/iDrac

1. Call flow

    ```bash
    main.go > router/router.go
        router.go
                // Create Storage
                dStore := storage.NewClient("root", "calvin", time.Duration(2))
                // Create Service
                dService := dellService.NewService(dStore)
                // Create Handler
                dHandler := dellHandler.NewHandler(dService)
                // Define function to handle request
                router.GET("/metrics/dell", dHandler.Metrics)

    handler/dell/metrics.go
        - define Custom Prometheus Interface.
            - Implementation Describe Function.
            - Implementation Collect Function.
        - This is a main file to code.

    handler/dell/metricsurl.go
        - define dell metrics URLs.

    handler/dell/dell.go
        - Register Prometheus Metrics.

    metric/dellMetrics.go
        - Define all struct to marshal from iDrac data to Golang Object

    metric/promMetrics.go
        - Define Exporter Metric in Prometheus Format

    service/dell/dell.go
        - Convert data from iDrac into Golang Object.

    storage/redfish.go
        - Getting data from iDrac/iLO
    ```
