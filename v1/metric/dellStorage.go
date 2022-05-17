package metric

// DellStorage start
type DellStorage struct {
	Drives      []Link
	DrivesCount int `json:"Drives@odata.count"`
}

// DellStorage end

// DellStorageDisk start
type DellStorageDisk struct {
	Id                            string  `json:Id`
	MediaType                     string  `json:MediaType`
	Name                          string  `json:Name`
	PartNumber                    string  `json:PartNumber`
	PredictedMediaLifeLeftPercent float64 `json:PredictedMediaLifeLeftPercent`
	SerialNumber                  string  `json:SerialNumber`
	CapacityBytes                 int     `json:CapacityBytes`
	Protocol                      string  `json:Protocol`
}

// DellStorageDisk end
