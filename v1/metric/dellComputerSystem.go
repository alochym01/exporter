package metric

// DellComputerSystem start
type DellComputerSystem struct {
	Oem          DellSystemComputerOEM
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

// DellComputerSystem end

// DellComputerSystemOME start
type DellSystemComputerOEM struct {
	Dell struct {
		DellSystem struct {
			BatteryRollupStatus string `json:"BatteryRollupStatus"`
			FanRollupStatus     string `json:"FanRollupStatus"`
			StorageRollupStatus string `json:"StorageRollupStatus"`
			TempRollupStatus    string `json:"TempRollupStatus"`
			PSRollupStatus      string `json:"PSRollupStatus"`
		}
	}
}

func (c DellSystemComputerOEM) StorageStatus() float64 {
	switch c.Dell.DellSystem.StorageRollupStatus {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	case "Degraded":
		return 1.0
	default:
		return 3.0
	}
}

func (c DellSystemComputerOEM) FansStatus() float64 {
	switch c.Dell.DellSystem.FanRollupStatus {
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
func (c DellSystemComputerOEM) PowerSupplyStatus() float64 {
	switch c.Dell.DellSystem.PSRollupStatus {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	case "Error":
		return 2.0
	default:
		return 3.0
	}
}
func (c DellSystemComputerOEM) TemperatureStatus() float64 {
	switch c.Dell.DellSystem.TempRollupStatus {
	case "OK":
		return 0.0
	case "Warning":
		return 1.0
	case "Critical":
		return 2.0
	case "Degraded":
		return 1.0
	default:
		return 3.0
	}
}

// DellComputerSystemOME end
