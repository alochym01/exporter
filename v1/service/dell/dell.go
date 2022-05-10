package dell

import (
	"encoding/json"

	"github.com/alochym01/exporter/v1/metric"
	"github.com/alochym01/exporter/v1/storage"
)

type Service struct {
	store storage.RedFishClient
}

// NewService return a DELLHandler struct
func NewService(s storage.RedFishClient) Service {
	return Service{
		store: s,
	}
}
func (s Service) ComputerSystem(url string) (*metric.DellComputerSystem, error) {
	data, err := s.store.Get(url)
	if err != nil {
		return nil, err
	}
	var computerSystem metric.DellComputerSystem

	err = json.Unmarshal(data, &computerSystem)
	return &computerSystem, nil
}
