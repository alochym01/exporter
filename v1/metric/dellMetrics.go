package metric

type Link struct {
	ODataID string `json:"@odata.id"`
}

type DellHealthStatus struct {
	Health string `json:"Health"`
}

type DellStateStatus struct {
	State string `json:"State"`
}

type DellHealthRollupStatus struct {
	HealthRollup string `json:"HealthRollup"`
}

type DellComputerSystemStatus struct {
	DellHealthStatus
	DellStateStatus
	DellHealthRollupStatus
}

type DellPowerStatus struct {
	DellHealthStatus
	DellStateStatus
}
