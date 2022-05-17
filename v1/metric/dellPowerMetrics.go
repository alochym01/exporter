package metric

// DellChasPowerControl start
type DellPower struct {
	Description             string                    `json:Description`
	Id                      string                    `json:Id`
	Name                    string                    `json:Name`
	PowerControl            []DellChasPowerControl    // Chassis
	PowerControlOdataCount  int                       `json:PowerControl@odata.count`
	PowerSupplies           []DellChasPowerSupplies   // Chassis
	PowerSuppliesOdataCount int                       `json:PowerSupplies@odata.count`
	Redundancy              []DellChasPowerRedundancy // Chassis
	RedundancyOdataCount    int                       `json:Redundancy@odata.count`
	// Voltages                []DellChasVoltages        // Chassis
	// VoltagesOdataCount      int                       `json:Voltages@odata.count`
}

type DellChasPowerControlPowerLimit struct {
	CorrectionInMs int    `json:CorrectionInMs`
	LimitException string `json:LimitException`
	LimitInWatts   int    `json:LimitInWatts`
}
type DellChasPowerControlPowerMetrics struct {
	AverageConsumedWatts int `json:AverageConsumedWatts`
	IntervalInMin        int `json:IntervalInMin`
	MaxConsumedWatts     int `json:MaxConsumedWatts`
	MinConsumedWatts     int `json:MinConsumedWatts`
}
type DellChasPowerControl struct {
	MemberId            string  `json:MemberId`
	Name                string  `json:Name`
	PowerAllocatedWatts int     `json:PowerAllocatedWatts`
	PowerAvailableWatts int     `json:PowerAvailableWatts`
	PowerCapacityWatts  int     `json:PowerCapacityWatts`
	PowerConsumedWatts  float64 `json:PowerConsumedWatts`
	PowerLimit          DellChasPowerControlPowerLimit
	PowerMetrics        DellChasPowerControlPowerMetrics
	PowerRequestedWatts int `json:PowerRequestedWatts`
	RelatedItem         []Link
}
type DellChasPowerSuppliesInputRanges struct {
	InputType          string `json:InputType`
	MaximumFrequencyHz int    `json:MaximumFrequencyHz`
	MaximumVoltage     int    `json:MaximumVoltage`
	MinimumFrequencyHz int    `json:MinimumFrequencyHz`
	MinimumVoltage     int    `json:MinimumVoltage`
	OutputWattage      int    `json:OutputWattage`
}

type DellChasPowerSupplies struct {
	Assembly             Link
	EfficiencyPercent    float64 `json:EfficiencyPercent`
	FirmwareVersion      string  `json:FirmwareVersion`
	HotPluggable         bool    `json:HotPluggable`
	InputRanges          []DellChasPowerSuppliesInputRanges
	LastPowerOutputWatts float64 `json:LastPowerOutputWatts`
	LineInputVoltage     float64 `json:LineInputVoltage`
	LineInputVoltageType string  `json:LineInputVoltageType`
	Manufacturer         string  `json:Manufacturer`
	MemberId             string  `json:MemberId`
	Model                string  `json:Model`
	Name                 string  `json:MemberId`
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
	Status               DellPowerStatus
}
type DellChasPowerRedundancy struct {
	MaxNumSupported         int    `json:MaxNumSupported`
	MemberId                string `json:MemberId`
	MinNumNeeded            int    `json:MinNumNeeded`
	Mode                    string `json:Mode`
	Name                    string `json:Name`
	RedundancySet           []Link
	RedundancySetOdataCount int `json:RedundancySet@odata.count`
	Status                  DellPowerStatus
}
type DellChasVoltages struct {
}

// DellChasPowerControl end
