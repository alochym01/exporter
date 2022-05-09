package dell

import "github.com/alochym01/exporter/v1/storage"

type Service struct {
	store storage.RedFishClient
}

// NewService return a DELLHandler struct
func NewService(s storage.RedFishClient) Service {
	return Service{
		store: s,
	}
}
func (s Service) SystemStatus(url string) ([]byte, error) {
	data, err := s.store.Get(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
