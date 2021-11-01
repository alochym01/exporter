package base

import "strings"

type RedfishLinksInstances struct {
	ODataID     string   `json:"@odata.id"`
	OdataType   string   `json:"@odata.type"`
	ETag        string   `json:"ETag"`
	HttpMethods []string `json:"HttpMethods"`
	/*{
	    "@odata.id": "/redfish/v1/Systems",
	    "@odata.type": "#ComputerSystemCollection.ComputerSystemCollection",
	    "ETag": "W/\"AA6D42B0\"",
	    "HttpMethods": [
	        "GET",
	        "HEAD"
	    ]
	}*/
}

type HpeiLOResourceDirectory struct {
	Meta
	Id        string `json:"Id"`
	OdataType string `json:"@odata.type"`
	Instances []RedfishLinksInstances
}

type StorageDisk struct {
	Meta
	BlockSizeBytes                    int      `json:"BlockSizeBytes"`
	CapacityGB                        float64  `json:"CapacityGB"`
	CapacityLogicalBlocks             int      `json:"CapacityLogicalBlocks"`
	CapacityMiB                       int      `json:"CapacityMiB"`
	CarrierApplicationVersion         string   `json:"CarrierApplicationVersion"`
	CarrierAuthenticationStatus       string   `json:"CarrierAuthenticationStatus"`
	CurrentTemperatureCelsius         int      `json:"CurrentTemperatureCelsius"`
	Description                       string   `json:"Description"`
	DiskDriveStatusReasons            []string `json:"DiskDriveStatusReasons"`
	DiskDriveUse                      string   `json:"DiskDriveUse"`
	EncryptedDrive                    bool     `json:"EncryptedDrive"`
	Id                                string   `json:"Id"`
	InterfaceSpeedMbps                int      `json:"InterfaceSpeedMbps"`
	InterfaceType                     string   `json:"InterfaceType"`
	LegacyBootPriority                string   `json:"LegacyBootPriority"`
	Location                          string   `json:"Location"`
	LocationFormat                    string   `json:"LocationFormat"`
	MaximumTemperatureCelsius         int      `json:"MaximumTemperatureCelsius"`
	MediaType                         string   `json:"MediaType"`
	Model                             string   `json:"Model"`
	Name                              string   `json:"Name"`
	PowerOnHours                      int      `json:"PowerOnHours"`
	RotationalSpeedRpm                int      `json:"RotationalSpeedRpm"`
	SSDEnduranceUtilizationPercentage float64  `json:"SSDEnduranceUtilizationPercentage"`
	SerialNumber                      string   `json:"SerialNumber"`
	Status                            Status
	UncorrectedReadErrors             int `json:"UncorrectedReadErrors"`
	UncorrectedWriteErrors            int `json:"UncorrectedWriteErrors"`
	// FirmwareVersion                   DiskDrivesFirmwareVersion
}

type HPESystemsOEMHpeAggregateHealthStatusStatus struct {
	Status HealthStatus
}

func (s HPESystemsOEMHpeAggregateHealthStatusStatus) StatusToNumber() float64 {
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

type HPESystemsOEMHpeAggregateHealthStatus struct {
	AgentlessManagementService string `json:"AgentlessManagementService"`
	BiosOrHardwareHealth       HPESystemsOEMHpeAggregateHealthStatusStatus
	FanRedundancy              string `json:"FanRedundancy"`
	Fans                       HPESystemsOEMHpeAggregateHealthStatusStatus
	Memory                     HPESystemsOEMHpeAggregateHealthStatusStatus
	Network                    HPESystemsOEMHpeAggregateHealthStatusStatus
	PowerSupplies              HPESystemsOEMHpeAggregateHealthStatusStatus
	Processors                 HPESystemsOEMHpeAggregateHealthStatusStatus
	SmartStorageBattery        HPESystemsOEMHpeAggregateHealthStatusStatus
	Storage                    HPESystemsOEMHpeAggregateHealthStatusStatus
	Temperatures               HPESystemsOEMHpeAggregateHealthStatusStatus
}
type HPESystemsOEMHpe struct {
	AggregateHealthStatus HPESystemsOEMHpeAggregateHealthStatus
}
type HPESystemsOEM struct {
	Hpe HPESystemsOEMHpe
}
type Systems struct {
	Meta
	// Actions
	AssetTag string `json:"AssetTag"`
	// Bios                   base.Link
	BiosVersion string `json:"BiosVersion"`
	// Boot                   SystemsBoot
	Description string `json:"Description"`
	// EthernetInterfaces     base.Link
	HostName string `json:"HostName"`
	// HostWatchdogTimer      SystemHostWatchdogTimer
	// HostingRoles           []string `json:"HostingRoles"`
	// HostingRolesOdataCount int      `json:"HostingRoles@odata.count"`
	Id           string `json:"Id"`
	IndicatorLED string `json:"IndicatorLED"`
	// Links                  SystemsLinks
	Manufacturer string `json:"Manufacturer"`
	// Memory                 base.Link
	// MemorySummary          SystemsMemorySummary
	Model string `json:"Model"`
	Name  string `json:"Name"`
	// NetworkInterfaces      base.Link
	Oem HPESystemsOEM
	// PCIeDevices            []base.Link
	// PCIeDevicesOdataCount  int    `json:"PCIeDevices@odata.count"`
	PartNumber string `json:"PartNumber"`
	PowerState string `json:"PowerState"`
	// ProcessorSummary       SystemsProcessorSummary
	// Processors             base.Link
	SKU string `json:"SKU"`
	// SecureBoot             base.Link
	SerialNumber string `json:"SerialNumber"`
	// SimpleStorage          base.Link
	Status Status
	// Storage                base.Link
	SystemType string `json:"SystemType"`
	// TrustedModules         []SystemTrustedModules
	UUID string `json:"UUID"`
}

func (s Systems) StatusToNumber() float64 {
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

type BaseNetworkAdaptersPhysicalPorts struct {
	FullDuplex bool   `json:"FullDuplex"`
	LinkStatus string `json:"LinkStatus"`
	MacAddress string `json:"MacAddress"`
	Name       string `json:"Name"`
	SpeedMbps  int    `json:"SpeedMbps"`
	Status     HealthStatus
}
type BaseNetworkAdapters struct {
	Meta
	Id             string `json:"Id"`
	Name           string `json:"Name"`
	PartNumber     string `json:"PartNumber"`
	PhysicalPorts  []BaseNetworkAdaptersPhysicalPorts
	SerialNumber   string `json:"SerialNumber"`
	Status         StateStatus
	StructuredName string `json:"StructuredName"`
}

func (b BaseNetworkAdaptersPhysicalPorts) PortStatus() float64 {
	switch strings.ToLower(b.Status.Health) {
	case "ok":
		return 0.0
	case "warning":
		return 1.0
	default:
		return 2.0
	}
	// switch b.SpeedMbps {
	// case 0:
	// 	return 2.0
	// default:
	// 	return 0.0
	// }
}

type PowerControlOem struct {
}

type ChassisPowerControlChassisPowerControl struct {
	AverageConsumedWatts int `json:"AverageConsumedWatts"`
	IntervalInMin        int `json:"IntervalInMin"`
	MaxConsumedWatts     int `json:"MaxConsumedWatts"`
	MinConsumedWatts     int `json:"MinConsumedWatts"`
}

type ChassisPowerControl struct {
	OdataID            string  `json:"@odata.id"`
	MemberId           string  `json:"MemberId"`
	PowerCapacityWatts int     `json:"PowerCapacityWatts"`
	PowerConsumedWatts float64 `json:"PowerConsumedWatts"`
	PowerMetrics       ChassisPowerControlChassisPowerControl
	/*
		{
			"@odata.id": "/redfish/v1/Chassis/1/Power#PowerControl/0",
			"MemberId": "0",
			"PowerCapacityWatts": 1600,
			"PowerConsumedWatts": 412,
			"PowerMetrics": {
				"AverageConsumedWatts": 408,
				"IntervalInMin": 20,
				"MaxConsumedWatts": 478,
				"MinConsumedWatts": 404
			}
		}
	*/
}
type ChassisPowerSupplies struct {
	OdataID  string `json:"@odata.id"`
	MemberId string `json:"MemberId"`

	/*
		{
			"@odata.id": "/redfish/v1/Chassis/1/Power#PowerSupplies/0",
			"FirmwareVersion": "1.02",
			"LastPowerOutputWatts": 192,
			"LineInputVoltage": 228,
			"LineInputVoltageType": "ACHighLine",
			"Manufacturer": "CHCNY",
			"MemberId": "0",
			"Model": "865414-B21",
			"Name": "HpeServerPowerSupply",
			"Oem": {
				"Hpe": {
					"@odata.context": "/redfish/v1/$metadata#HpeServerPowerSupply.HpeServerPowerSupply",
					"@odata.type": "#HpeServerPowerSupply.v2_0_0.HpeServerPowerSupply",
					"AveragePowerOutputWatts": 192,
					"BayNumber": 1,
					"HotplugCapable": true,
					"MaxPowerOutputWatts": 219,
					"Mismatched": false,
					"PowerSupplyStatus": {
						"State": "Ok"
					},
					"iPDUCapable": false
				}
			}
		}
	*/
}
type ChassisRedundancy struct {
}
type PowerControl struct {
	Meta
	Id            string `json:"Id"`
	Name          string `json:"Name"`
	Oem           PowerControlOem
	PowerControl  []ChassisPowerControl
	PowerSupplies []ChassisPowerSupplies
	Redundancy    []ChassisRedundancy
}
