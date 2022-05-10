package metric

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

type DellComputerSystem struct {
	Model        string `json:Model`
	PartNumber   string `json:PartNumber`
	SerialNumber string `json:SerialNumber`
	SKU          string `json:SKU`
	Status       DellComputerSystemStatus
}

func (s DellComputerSystem) StatusToNumber() float64 {
	switch s.Status.Health {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	default:
		return 3.0
	}
}

// type DellSystems struct {
// 	Meta
// 	Actions                DellActions
// 	AssetTag               string `json:AssetTag`
// 	Bios                   Link
// 	BiosVersion            string `json:BiosVersion`
// 	Boot                   DellSystemsBoot
// 	Description            string `json:Description`
// 	EthernetInterfaces     Link
// 	HostName               string `json:HostName`
// 	HostWatchdogTimer      DellSystemHostWatchdogTimer
// 	HostingRoles           []string `json:HostingRoles`
// 	HostingRolesOdataCount int      `json:HostingRoles@odata.count`
// 	Id                     string   `json:Id`
// 	IndicatorLED           string   `json:IndicatorLED`
// 	Links                  DellSystemsLinks
// 	Manufacturer           string `json:Manufacturer`
// 	Memory                 Link
// 	MemorySummary          DellSystemsMemorySummary
// 	Model                  string `json:Model`
// 	Name                   string `json:Name`
// 	NetworkInterfaces      Link
// 	Oem                    DellSystemsOEM
// 	PCIeDevices            []Link
// 	PCIeDevicesOdataCount  int    `json:PCIeDevices@odata.count`
// 	PartNumber             string `json:PartNumber`
// 	PowerState             string `json:PowerState`
// 	ProcessorSummary       DellSystemsProcessorSummary
// 	Processors             Link
// 	SKU                    string `json:SKU`
// 	SecureBoot             Link
// 	SerialNumber           string `json:SerialNumber`
// 	SimpleStorage          Link
// 	Status                 Status
// 	Storage                Link
// 	SystemType             string `json:SystemType`
// 	TrustedModules         []DellSystemTrustedModules
// 	UUID                   string `json:UUID`
// }
