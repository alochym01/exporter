package dell

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alochym01/exporter/v1/metric"
	"github.com/alochym01/exporter/v1/storage"
	"github.com/prometheus/client_golang/prometheus"
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

// Storage handle error should return Golang Object type with having Error
func (s Service) Storage(url string) (*metric.DellStorage, error) {
	data, err := s.store.Get(url)
	if err != nil {
		return nil, err
	}
	var computerStorage metric.DellStorage

	err = json.Unmarshal(data, &computerStorage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &computerStorage, nil
}

// Storage handle error should return Golang Object type with having Error
func (s Service) StorageWithChannel(url, server string, ch chan<- prometheus.Metric) {
	data, err := s.store.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	var computerStorage metric.DellStorage

	err = json.Unmarshal(data, &computerStorage)
	if err != nil {
		fmt.Println(err)
		return
	}

	// get Disk info
	// disks := make(chan metric.DellStorageDisk, len(computerStorage.Drives))
	for _, v := range computerStorage.Drives {
		// fmt.Println("index ", i)
		// fmt.Println("value ", v.ODataID)
		diskURL := server + v.ODataID
		// go s.StorageDiskWithChannel(diskURL, disks)
		disk, err := s.StorageDisk(diskURL)
		if err == nil {
			if disk.PredictedMediaLifeLeftPercent > 0 {
				// string{"id", "capacity", "interface_type", "media_type"},
				// storage_drive_ssd_endurance{capacity="959",id="18",interface_type="SAS",media_type="SSD"} 100
				enclosure := strings.Split(disk.Id, ".")[2]
				id := strings.Split(enclosure, ":")[0]
				ch <- prometheus.MustNewConstMetric(
					metric.SysStorageDisk,
					prometheus.GaugeValue,
					disk.PredictedMediaLifeLeftPercent,
					fmt.Sprintf("%s", id),
					fmt.Sprintf("%d", disk.CapacityBytes/1000000000),
					disk.Protocol,
					disk.MediaType,
				)
			}
		}
	}

	// for range computerStorage.Drives {
	// 	disk := <-disks
	// 	if disk.PredictedMediaLifeLeftPercent > 0 {
	// 		// string{"id", "capacity", "interface_type", "media_type"},
	// 		// storage_drive_ssd_endurance{capacity="959",id="18",interface_type="SAS",media_type="SSD"} 100
	// 		enclosure := strings.Split(disk.Id, ".")[2]
	// 		id := strings.Split(enclosure, ":")[0]
	// 		ch <- prometheus.MustNewConstMetric(
	// 			metric.SysStorageDisk,
	// 			prometheus.GaugeValue,
	// 			disk.PredictedMediaLifeLeftPercent,
	// 			fmt.Sprintf("%s", id),
	// 			fmt.Sprintf("%d", disk.CapacityBytes/1000000000),
	// 			disk.Protocol,
	// 			disk.MediaType,
	// 		)
	// 	}
	// }
}

// StorageDisk handle error handle error should return Golang Object type with having Error
func (s Service) StorageDisk(url string) (*metric.DellStorageDisk, error) {
	data, err := s.store.Get(url)
	if err != nil {
		return nil, err
	}
	var storageDisk metric.DellStorageDisk

	err = json.Unmarshal(data, &storageDisk)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	b, err := json.MarshalIndent(storageDisk, "", "   ")
	fmt.Println(string(b))
	return &storageDisk, nil
}

// StorageDiskWithChannel
func (s Service) StorageDiskWithChannel(url string, disk chan<- metric.DellStorageDisk) {
	data, err := s.store.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	var storageDisk metric.DellStorageDisk

	err = json.Unmarshal(data, &storageDisk)
	if err != nil {
		fmt.Println(err)
		return
	}

	// b, err := json.MarshalIndent(storageDisk, "", "   ")
	// fmt.Println(string(b))
	disk <- storageDisk
}
