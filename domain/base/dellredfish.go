package base

type SystemsCollection struct {
	Meta
	Description       string `json:Description`
	Members           []Link
	MembersOdataCount int    `json:Members@odata.count`
	Name              string `json:Name`
}

type DellSystemTrustedModules struct {
}
type DellSystemsProcessorSummary struct {
}
type DellChassisOEMDellDellSystem struct {
	BatteryRollupStatus string `json:"BatteryRollupStatus"`
	FanRollupStatus     string `json:"FanRollupStatus"`
	StorageRollupStatus string `json:"StorageRollupStatus"`
	TempRollupStatus    string `json:"TempRollupStatus"`
	PSRollupStatus      string `json:"PSRollupStatus"`
	// CPURollupStatus     string `json:"CPURollupStatus"`
	// BatteryRollupStatus string `json:"BatteryRollupStatus"`
	// BatteryRollupStatus string `json:"BatteryRollupStatus"`
	// BatteryRollupStatus string `json:"BatteryRollupStatus"`
	// BatteryRollupStatus string `json:"BatteryRollupStatus"`
	// BatteryRollupStatus string `json:"BatteryRollupStatus"`
	// BatteryRollupStatus string `json:"BatteryRollupStatus"`
}

func (c DellChassisOEMDellDellSystem) StorageStatus() float64 {
	switch c.StorageRollupStatus {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	case "Degraded":
		return 2.0
	default:
		return 3.0
	}
}
func (c DellChassisOEMDellDellSystem) FansStatus() float64 {
	switch c.FanRollupStatus {
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
func (c DellChassisOEMDellDellSystem) PowerSupplyStatus() float64 {
	switch c.PSRollupStatus {
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
func (c DellChassisOEMDellDellSystem) TemperatureStatus() float64 {
	switch c.TempRollupStatus {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	case "Degraded":
		return 2.0
	default:
		return 3.0
	}
}

type DellSystemsOEMDell struct {
	DellSystem DellChassisOEMDellDellSystem
}
type DellSystemsOEM struct {
	Dell DellSystemsOEMDell
}
type DellSystemsMemorySummary struct {
}
type DellSystemsLinks struct {
}
type DellSystemHostWatchdogTimer struct {
}
type DellSystemsBoot struct {
}
type DellActions struct {
}
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

func (s DellSystems) StatusToNumber() float64 {

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

type DellStorageCollection struct {
	Meta
	Description       string `json:Description`
	Members           []Link
	MembersOdataCount int    `json:Members@odata.count`
	Name              string `json:Name`
}

//
// type DellSystemsStorageControllers struct {
// }
//
type DellStorage struct {
	Meta
	Description                  string `json:Description`
	Drives                       []Link
	DrivesOdataCount             int    `json:Drives@odata.count`
	Id                           string `json:Id`
	Links                        DellSystemsStorageLinks
	Name                         string `json:Name`
	OEM                          DellSystemsStorageOEM
	Status                       Status
	StorageControllers           []DellSystemsStorageControllers
	StorageControllersOdataCount int `json:StorageControllers@odata.count`
	Volumes                      Link
}

func (s DellStorage) StatusToNumber() float64 {
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

type DellSystemsStorageLinks struct {
}
type DellSystemsStorageOEM struct {
}
type DellSystemsStorageControllers struct {
}

type DellSystemsStorageDiskIdentifiers struct {
}
type DellSystemsStorageDiskLinks struct {
}
type DellSystemsStorageDiskPhysicalLocation struct {
}
type DellStorageDisk struct {
	Meta
	BlockSizeBytes                int    `json:BlockSizeBytes`
	CapableSpeedGbs               int    `json:CapableSpeedGbs`
	CapacityBytes                 int    `json:CapacityBytes`
	Description                   string `json:Description`
	EncryptionAbility             string `json:EncryptionAbility`
	EncryptionStatus              string `json:EncryptionStatus`
	FailurePredicted              bool   `json:FailurePredicted`
	Id                            string `json:Id`
	Identifiers                   []DellSystemsStorageDiskIdentifiers
	Links                         DellSystemsStorageDiskLinks
	Location                      []string `json:Location`
	Manufacturer                  string   `json:Manufacturer`
	MediaType                     string   `json:MediaType`
	Name                          string   `json:Name`
	NegotiatedSpeedGbs            int      `json:NegotiatedSpeedGbs`
	PartNumber                    string   `json:PartNumber`
	PhysicalLocation              DellSystemsStorageDiskPhysicalLocation
	PredictedMediaLifeLeftPercent float64 `json:PredictedMediaLifeLeftPercent`
	Protocol                      string  `json:Protocol`
	Revision                      string  `json:Revision`
	SerialNumber                  string  `json:SerialNumber`
	Status                        Status
	// RotationSpeedRPM              string  `json:RotationSpeedRPM`
	// Actions SystemsStorageDiskAction
	// Assembly                      base.Link
	// IdentifiersOdataCount         int `json:Identifiers@odata.count`
	// OEM                           SystemsStorageDiskOEM
	// Operations                    []string `json:Operations`
	// OperationsOdataCount          int      `json:Operations@odata.count`
}

type DellEthernetInterfaceCollection struct {
	Meta
	Description       string `json:Description`
	Members           []Link
	MembersOdataCount int    `json:Members@odata.count`
	Name              string `json:Name`
}

type DellEthernetInterface struct {
	Meta
	AutoNeg     bool   `json:AutoNeg`
	Description string `json:Description`
	FQDN        string `json:FQDN`
	FullDuplex  bool   `json:FullDuplex`
	HostName    string `json:HostName`
	Id          string `json:Id`
	LinkStatus  string `json:LinkStatus`
	MACAddress  string `json:MACAddress`
	MTUSize     int    `json:MTUSize`
	Name        string `json:Name`
	SpeedMbps   int    `json:SpeedMbps`
	Status      Status
	// Links       SystemsEthernetInterfaceLinks
	// InterfaceEnabled bool   `json:InterfaceEnabled`
	// UefiDevicePath   string `json:UefiDevicePath`
	// IPv4Addresses       []string `json:IPv4Addresses`
	// IPv6Addresses       []string `json:IPv6Addresses`
	// PermanentMACAddress string `json:PermanentMACAddress`
}

func (e DellEthernetInterface) PortStatus() float64 {
	switch e.SpeedMbps {
	case 0:
		return 2.0
	default:
		return 0.0
	}
	// switch e.LinkStatus {
	// case LinkUp:
	// 	return 0.0
	// case LinkDown:
	// 	return 2.0
	// default:
	// 	return 3.0
	// }
}

type DellChassisCollection struct {
	Meta
	Description       string `json:Description`
	Members           []Link
	MembersOdataCount int    `json:Members@odata.count`
	Name              string `json:Name`
}

type DellChassisPhysicalSecurity struct {
}

// type ChassisOEM struct {
// 	Dell ChassisOEMDell
// }
type ChassisOEMDell struct {
	DellChassis ChassisOEMDellDellChassis
}

type ChassisOEMDellDellChassis struct {
	Meta
	CanBeFRUed bool `json:CanBeFRUed`
	Links      ChassisOEMDellDellChassisLinks
	SystemID   int `json:SystemID`
}

type ChassisOEMDellDellChassisLinks struct {
	ComputerSystem Link
}
type DellChassisOEM struct {
	Dell ChassisOEMDell
}
type DellChassisLocationPostalAddress struct {
	Building string `json:Building`
	Room     string `json:Room`
}
type DellChassisLocationPlacement struct {
	Rack string `json:Rack`
	Row  string `json:Row`
}
type DellChassisLocation struct {
	Info          string `json:Info`
	InfoFormat    string `json:InfoFormat`
	Placement     DellChassisLocationPlacement
	PostalAddress DellChassisLocationPostalAddress
}
type DellChassisLinks struct {
	ComputerSystems             []Link `json:ComputerSystems`
	ManagedBy                   []Link
	ComputerSystemsOdataCount   int    `json:ComputerSystems@odata.count`
	Contains                    []Link `json:Contains`
	ContainsOdataCount          int    `json:Contains@odata.count`
	CooledBy                    []Link
	CooledByOdataCount          int      `json:CooledBy@odata.count`
	Drives                      []string `json:Drives`
	DrivesOdataCount            int      `json:Drives@odata.count`
	ManagedByOdataCount         int      `json:ManagedBy@odata.count`
	ManagersInChassis           []Link
	ManagersInChassisOdataCount int `json:ManagersInChassis@odata.count`
	PCIeDevices                 []Link
	PCIeDevicesOdataCount       int `json:PCIeDevices@odata.count`
	Processors                  []Link
	ProcessorsOdataCount        int `json:Processors@odata.count`
	PoweredBy                   []Link
	PoweredByOdataCount         int `json:PoweredBy@odata.count`
	Storage                     []Link
	StorageOdataCount           int `json:Storage@odata.count`
}
type DellChassis struct {
	// base.Chassis
	Meta
	AssetTag         string `json:AssetTag`
	ChassisType      string `json:ChassisType`
	Id               string `json:Id`
	IndicatorLED     string `json:IndicatorLED`
	Manufacturer     string `json:Manufacturer`
	Model            string `json:Model`
	Name             string `json:Name`
	NetworkAdapters  Link
	Power            Link
	PowerState       string `json:PowerState`
	SKU              string `json:SKU`
	SerialNumber     string `json:SerialNumber`
	Thermal          Link
	Assembly         Link
	Description      string `json:Description`
	Links            DellChassisLinks
	Location         DellChassisLocation
	Oem              DellChassisOEM
	PCIeSlots        Link
	PartNumber       string `json:PartNumber`
	PhysicalSecurity DellChassisPhysicalSecurity
	Sensors          Link
	Status           Status
	UUID             string `json:UUID`
	// base.Actions
}

func (c DellChassis) StatusToNumber() float64 {
	switch c.Status.Health {
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

type ChasPowerControlPowerLimit struct {
	CorrectionInMs int    `json:CorrectionInMs`
	LimitException string `json:LimitException`
	LimitInWatts   int    `json:LimitInWatts`
}
type ChasPowerControlPowerMetrics struct {
	AverageConsumedWatts int `json:AverageConsumedWatts`
	IntervalInMin        int `json:IntervalInMin`
	MaxConsumedWatts     int `json:MaxConsumedWatts`
	MinConsumedWatts     int `json:MinConsumedWatts`
}
type DellChasPowerControl struct {
	Meta
	MemberId            string  `json:MemberId`
	Name                string  `json:Name`
	PowerAllocatedWatts int     `json:PowerAllocatedWatts`
	PowerAvailableWatts int     `json:PowerAvailableWatts`
	PowerCapacityWatts  int     `json:PowerCapacityWatts`
	PowerConsumedWatts  float64 `json:PowerConsumedWatts`
	PowerLimit          ChasPowerControlPowerLimit
	PowerMetrics        ChasPowerControlPowerMetrics
	PowerRequestedWatts int `json:PowerRequestedWatts`
	RelatedItem         []Link
}
type ChasPowerSuppliesInputRanges struct {
	InputType          string `json:InputType`
	MaximumFrequencyHz int    `json:MaximumFrequencyHz`
	MaximumVoltage     int    `json:MaximumVoltage`
	MinimumFrequencyHz int    `json:MinimumFrequencyHz`
	MinimumVoltage     int    `json:MinimumVoltage`
	OutputWattage      int    `json:OutputWattage`
}
type ChasPowerSuppliesOem struct {
}
type DellChasPowerSupplies struct {
	Meta
	Assembly             Link
	EfficiencyPercent    float64 `json:EfficiencyPercent`
	FirmwareVersion      string  `json:FirmwareVersion`
	HotPluggable         bool    `json:HotPluggable`
	InputRanges          []ChasPowerSuppliesInputRanges
	LastPowerOutputWatts float64 `json:LastPowerOutputWatts`
	LineInputVoltage     float64 `json:LineInputVoltage`
	LineInputVoltageType string  `json:LineInputVoltageType`
	Manufacturer         string  `json:Manufacturer`
	MemberId             string  `json:MemberId`
	Model                string  `json:Model`
	Name                 string  `json:MemberId`
	OEM                  ChasPowerSuppliesOem
	PartNumber           string  `json:PartNumber`
	PowerCapacityWatts   float64 `json:PowerCapacityWatts`
	PowerInputWatts      float64 `json:PowerInputWatts`
	PowerOutputWatts     float64 `json:PowerOutputWatts`
	PowerSupplyType      string  `json:PowerSupplyType`
	Redundancy           []DellChasPowerRedundancy
	RedundancyOdataCount int `json:Redundancy@odata.count`
	RelatedItem          []Link
	SerialNumber         string `json:SerialNumber`
	SparePartNumber      string `json:SparePartNumber`
	Status               Status
}
type DellChasPowerRedundancy struct {
	Meta
	MaxNumSupported         int    `json:MaxNumSupported`
	MemberId                string `json:MemberId`
	MinNumNeeded            int    `json:MinNumNeeded`
	Mode                    string `json:Mode`
	Name                    string `json:Name`
	RedundancySet           []Link
	RedundancySetOdataCount int `json:RedundancySet@odata.count`
	Status                  Status
}
type DellChasVoltages struct {
}
type DellPowerControl struct {
	Meta
	Description             string                    `json:Description`
	Id                      string                    `json:Id`
	Name                    string                    `json:Name`
	PowerControl            []DellChasPowerControl    // Chassis
	PowerControlOdataCount  int                       `json:PowerControl@odata.count`
	PowerSupplies           []DellChasPowerSupplies   // Chassis
	PowerSuppliesOdataCount int                       `json:PowerSupplies@odata.count`
	Redundancy              []DellChasPowerRedundancy // Chassis
	RedundancyOdataCount    int                       `json:Redundancy@odata.count`
	Voltages                []DellChasVoltages        // Chassis
	VoltagesOdataCount      int                       `json:Voltages@odata.count`
}
