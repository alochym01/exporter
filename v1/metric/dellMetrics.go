package metric

// type SystemStatus struct {

// }

type DellSystems struct {
	Meta
	Actions                DellActions
	AssetTag               string `json:AssetTag`
	Bios                   Link
	BiosVersion            string `json:BiosVersion`
	Boot                   DellSystemsBoot
	Description            string `json:Description`
	EthernetInterfaces     Link
	HostName               string `json:HostName`
	HostWatchdogTimer      DellSystemHostWatchdogTimer
	HostingRoles           []string `json:HostingRoles`
	HostingRolesOdataCount int      `json:HostingRoles@odata.count`
	Id                     string   `json:Id`
	IndicatorLED           string   `json:IndicatorLED`
	Links                  DellSystemsLinks
	Manufacturer           string `json:Manufacturer`
	Memory                 Link
	MemorySummary          DellSystemsMemorySummary
	Model                  string `json:Model`
	Name                   string `json:Name`
	NetworkInterfaces      Link
	Oem                    DellSystemsOEM
	PCIeDevices            []Link
	PCIeDevicesOdataCount  int    `json:PCIeDevices@odata.count`
	PartNumber             string `json:PartNumber`
	PowerState             string `json:PowerState`
	ProcessorSummary       DellSystemsProcessorSummary
	Processors             Link
	SKU                    string `json:SKU`
	SecureBoot             Link
	SerialNumber           string `json:SerialNumber`
	SimpleStorage          Link
	Status                 Status
	Storage                Link
	SystemType             string `json:SystemType`
	TrustedModules         []DellSystemTrustedModules
	UUID                   string `json:UUID`
}
