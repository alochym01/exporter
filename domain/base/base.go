package base

type Link struct {
	ODataID string `json:"@odata.id"`
}

type Meta struct {
	ODataContext string `json:"@odata.context"`
	ODataID      string `json:"@odata.id"`
	ODataType    string `json:"@odata.type"`
}

type HealthStatus struct {
	Health string `json:"Health"`
}
type Status struct {
	HealthStatus
	StateStatus
	HealthRollupStatus
}
type StateStatus struct {
	State string `json:"State"`
}

type HealthRollupStatus struct {
	HealthRollup string `json:"HealthRollup"`
}
