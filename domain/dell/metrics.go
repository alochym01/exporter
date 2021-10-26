package dell

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/alochym01/exporter/domain/base"
	"github.com/alochym01/exporter/storage/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

type Metrics struct {
	server string
}

// Describe a description of metrics
func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
	ch <- base.SysStorageStatus
	ch <- base.SysStorageDisk
	ch <- base.SysEthernetInterface
	ch <- base.ChasPower
	ch <- base.ChasFansStatus
	ch <- base.ChasPowerStatus
	ch <- base.ChasTemperatureStatus
}

// Collect return a metric with all desc value and metric value
func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	client := redfish.NewAPIClient("root", "calvin")
	// System Metrics
	system, err := m.SetSystemHealthMetrics(ch, client)
	if err != nil {
		return
	}
	m.SetDiskMetrics(ch, client, system)
	m.SetEthernetMetric(ch, client, system.EthernetInterfaces.ODataID)
	m.ChassisCollector(ch, client)
}

// NewMetrics return a Metrics struct
func NewMetrics(s string) Metrics {
	return Metrics{
		server: s,
	}
}

func (m Metrics) SetSystemHealthMetrics(chSys chan<- prometheus.Metric, c redfish.APIClient) (*base.DellSystems, error) {
	sysCollectionURL := fmt.Sprintf("%s%s", m.server, "/redfish/v1/Systems")
	sysCollectionData, err := c.Get(sysCollectionURL)
	// Cannot connect to Server
	if err != nil {
		fmt.Println(sysCollectionURL)
		fmt.Println("Not Found")
		chSys <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return nil, err
	}
	var sysCollection base.SystemsCollection
	err = json.Unmarshal(sysCollectionData, &sysCollection)
	// Data cannot convert System struct
	if err != nil {
		fmt.Println(sysCollectionURL)
		fmt.Println(err.Error())
		chSys <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
		return nil, err
	}
	var sys base.DellSystems
	for _, v := range sysCollection.Members {
		sysURL := fmt.Sprintf("%s%s", m.server, v.ODataID)
		sysData, err := c.Get(sysURL)
		// Cannot connect to Server
		if err != nil {
			fmt.Println(sysURL)
			fmt.Println("Not Found")
			chSys <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return nil, err
		}

		err = json.Unmarshal(sysData, &sys)
		// Data cannot convert System struct
		if err != nil {
			fmt.Println(sysURL)
			fmt.Println(err.Error())
			chSys <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return nil, err
		}
	}
	// b, err := json.MarshalIndent(sys, "", "   ")
	// fmt.Println(string(b))
	// fmt.Println(sys.Oem.Dell.DellSystem.StorageStatus())
	chSys <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, sys.StatusToNumber(), sys.SKU, sys.SerialNumber)
	chSys <- prometheus.MustNewConstMetric(base.SysStorageStatus, prometheus.GaugeValue, sys.Oem.Dell.DellSystem.StorageStatus())
	chSys <- prometheus.MustNewConstMetric(base.ChasFansStatus, prometheus.GaugeValue, sys.Oem.Dell.DellSystem.FansStatus())
	chSys <- prometheus.MustNewConstMetric(base.ChasPowerStatus, prometheus.GaugeValue, sys.Oem.Dell.DellSystem.PowerSupplyStatus())
	chSys <- prometheus.MustNewConstMetric(base.ChasTemperatureStatus, prometheus.GaugeValue, sys.Oem.Dell.DellSystem.TemperatureStatus())
	return &sys, nil
}

func (m Metrics) SetDiskMetrics(chDisk chan<- prometheus.Metric, c redfish.APIClient, system *base.DellSystems) {
	url := fmt.Sprintf("%s%s", m.server, system.Storage.ODataID)
	data, err := c.Get(url)
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	var storeCollection base.DellStorageCollection
	err = json.Unmarshal(data, &storeCollection)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	// Systems Storage Data start
	// Set Systems Storage URL
	var storeURL string
	for _, v := range storeCollection.Members {
		if strings.Contains(v.ODataID, "RAID") {
			// storeURL = fmt.Sprintf("%s%s", redfish.ClientDELL.Host, v.ODataID)
			storeURL = fmt.Sprintf("%s%s", m.server, v.ODataID)
			break
		}
	}
	// Get Systems Storage Data
	storeData, err := c.Get(storeURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(storeURL)
		fmt.Println(err.Error())
		return
	}
	var store base.DellStorage
	err = json.Unmarshal(storeData, &store)
	// Data cannot convert Storage struct
	if err != nil {
		fmt.Println(storeURL)
		fmt.Println(err.Error())
		return
	}

	// Set Disk metrics
	diskAlochym := make(chan base.DellStorageDisk)
	// TODO go routine start
	for _, v := range store.Drives {
		diskURL := fmt.Sprintf("%s%s", m.server, v.ODataID)
		go getDisk(c, diskURL, diskAlochym)
	}
	for range store.Drives {
		disk := <-diskAlochym
		// Check Disk is SSD
		if disk.PredictedMediaLifeLeftPercent > 0 {
			// m.sysStorageDisk(ch, disk)
			enclosure := strings.Split(disk.Id, ".")[2]
			id := strings.Split(enclosure, ":")[0]
			chDisk <- prometheus.MustNewConstMetric(
				base.SysStorageDisk,
				prometheus.GaugeValue,
				disk.PredictedMediaLifeLeftPercent,
				fmt.Sprintf("%s", id),
				fmt.Sprintf("%d", disk.CapacityBytes/1000000000),
				disk.Protocol,
				disk.MediaType,
			)
		}
	}
	// TODO go routine end
}

func getDisk(c redfish.APIClient, url string, diskChan chan<- base.DellStorageDisk) {
	diskData, err := c.Get(url)
	// Problem connect to server
	if err != nil {
		return
	}
	var disk base.DellStorageDisk
	err = json.Unmarshal(diskData, &disk)
	// Data cannot convert StorageDisk struct
	if err != nil {
		return
	}
	diskChan <- disk
}

func (m Metrics) SetEthernetMetric(chIface chan<- prometheus.Metric, c redfish.APIClient, url string) {
	ifurl := fmt.Sprintf("%s%s", m.server, url)
	data, err := c.Get(ifurl)
	if err != nil {
		fmt.Println(ifurl)
		fmt.Println(err.Error())
		return
	}
	var ethCollectionIf base.DellEthernetInterfaceCollection
	err = json.Unmarshal(data, &ethCollectionIf)
	// Data cannot convert StorageCollection struct
	if err != nil {
		fmt.Println(ifurl)
		fmt.Println(err.Error())
		return
	}
	// Systems Ethernet Interfaces Collection end

	// Systems Ethernet Interfaces start
	// TODO go routine start
	ifAlochym := make(chan base.DellEthernetInterface)
	for _, v := range ethCollectionIf.Members {
		ifURL := fmt.Sprintf("%s%s", m.server, v.ODataID)
		go getIface(c, ifURL, ifAlochym)
	}
	for range ethCollectionIf.Members {
		iface := <-ifAlochym
		chIface <- prometheus.MustNewConstMetric(
			base.SysEthernetInterface,
			prometheus.GaugeValue,
			iface.PortStatus(),
			iface.Id,
			iface.MACAddress,
			fmt.Sprintf("%d", iface.SpeedMbps),
		)
	}
	// TODO go routine end
	return
}

func getIface(c redfish.APIClient, url string, ifChan chan<- base.DellEthernetInterface) {
	// TODO go routine start
	// ifURL := fmt.Sprintf("%s%s", m.server, url)
	// Get Ethernet Interfaces Data
	ifData, err := c.Get(url)
	// Problem connect to server
	if err != nil {
		fmt.Println("Ethernets Interface", url)
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	var iface base.DellEthernetInterface
	err = json.Unmarshal(ifData, &iface)
	// Data cannot convert EthernetInterface struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	ifChan <- iface
}

func (m Metrics) ChassisCollector(ch chan<- prometheus.Metric, c redfish.APIClient) {
	// Chassis Collection start
	chasCollectionURL := fmt.Sprintf("%s%s", m.server, "/redfish/v1/Chassis")
	chas, err := m.GetChassis(chasCollectionURL, c)
	if err != nil {
		return
	}
	// Chassis Collection end

	// Set PowerControl Link
	chasPowerLink := fmt.Sprintf("%s%s", m.server, chas.Power.ODataID)
	m.SetPowerMetrics(ch, c, chasPowerLink)

	// Set Thermal Link
	// chasThermalLink := fmt.Sprintf("%s%s", Host, chas.Thermal.ODataID)
}

func (m Metrics) GetChassis(url string, c redfish.APIClient) (*base.DellChassis, error) {
	var chasCollection base.DellChassisCollection
	dataCollection, err := c.Get(url)
	// Problem connect to server
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return nil, err
	}
	err = json.Unmarshal(dataCollection, &chasCollection)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return nil, err
	}
	// b, _ := json.MarshalIndent(chas, "", "    ")
	// fmt.Println(string(b))

	// Chassis start
	// Set a chassis url
	var chasURL string
	for _, v := range chasCollection.Members {
		if strings.Contains(v.ODataID, "System") {
			chasURL = fmt.Sprintf("%s%s", m.server, v.ODataID)
			break
		}
	}

	// get Chassis Data
	dataChassis, err := c.Get(chasURL)
	// Problem connect to server
	if err != nil {
		fmt.Println(chasURL)
		fmt.Println(err.Error())
		return nil, err
	}

	var chas base.DellChassis
	// Data cannot convert Chassis struct
	err = json.Unmarshal(dataChassis, &chas)
	if err != nil {
		fmt.Println(chasURL)
		fmt.Println(err.Error())
		return nil, err
	}

	// b, _ := json.MarshalIndent(chas, "", "    ")
	// fmt.Println(string(b))
	return &chas, nil
}

func (m Metrics) SetPowerMetrics(ch chan<- prometheus.Metric, c redfish.APIClient, url string) {
	data, err := c.Get(url)
	// Problem connect to server
	if err != nil {
		return
	}
	var power base.DellPowerControl
	err = json.Unmarshal(data, &power)
	// Data cannot convert PowerControl struct
	if err != nil {
		fmt.Println(url)
		fmt.Println(err.Error())
		return
	}
	// b, _ := json.MarshalIndent(power, "", "    ")
	// fmt.Println(string(b))
	// Everything is ok
	for _, v := range power.PowerControl {
		ch <- prometheus.MustNewConstMetric(
			base.ChasPower,
			prometheus.GaugeValue,
			v.PowerConsumedWatts,
		)
	}
	return
}
