package hpe

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

func NewMetrics(s string) Metrics {
	return Metrics{server: s}
}

func (m Metrics) Describe(ch chan<- *prometheus.Desc) {
	ch <- base.SysState
	// ch <- base.SysStorageState
	ch <- base.SysStorageDisk
	ch <- base.SysEthernetInterface
	ch <- base.ChasPower
}

func (m Metrics) Collect(ch chan<- prometheus.Metric) {
	client := redfish.NewAPIClient("username", "password")
	// Get all Redfish Link
	hpeiLORD, err := m.getResourceDirectory(client)
	if err != nil {
		ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "")
		return
	}
	// System Metrics
	m.SetSystemHealthMetrics(ch, client, *hpeiLORD)
	m.SetDiskMetrics(ch, client, *hpeiLORD)
	m.SetEthernetMetrics(ch, client, *hpeiLORD)
	m.SetPowerMetrics(ch, client, *hpeiLORD)
}

func (m Metrics) getResourceDirectory(c redfish.APIClient) (*base.HpeiLOResourceDirectory, error) {
	redfishLinks := fmt.Sprintf("%s%s", m.server, "/redfish/v1/resourcedirectory")
	data, err := c.Get(redfishLinks)
	if err != nil {
		fmt.Println(redfishLinks)
		return nil, err
	}
	var rfLinks base.HpeiLOResourceDirectory

	err = json.Unmarshal(data, &rfLinks)
	// Data cannot convert ChassisCollection struct
	if err != nil {
		fmt.Println(redfishLinks)
		fmt.Println(err.Error())
		return nil, err
	}
	return &rfLinks, nil
}
func (m Metrics) SetSystemHealthMetrics(ch chan<- prometheus.Metric, c redfish.APIClient, HpeiLORD base.HpeiLOResourceDirectory) {
	// get System URL
	sysURL := findObject(HpeiLORD.Instances, "ComputerSystem.", m.server)

	for _, v := range sysURL {
		if v == "" {
			fmt.Println(v)
			fmt.Println("Not Found")
			ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return
		}
		sysData, err := c.Get(v)
		// Problem connect to server
		if err != nil {
			fmt.Println(v)
			fmt.Println(err.Error())
			ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return
		}
		var sys base.Systems
		err = json.Unmarshal(sysData, &sys)
		// Data cannot convert System struct
		if err != nil {
			fmt.Println(v)
			fmt.Println(err.Error())
			ch <- prometheus.MustNewConstMetric(base.SysState, prometheus.GaugeValue, 2.0, "", "")
			return
		}
		ch <- prometheus.MustNewConstMetric(
			base.SysState,
			prometheus.GaugeValue,
			sys.StatusToNumber(),
			sys.SKU,
			sys.SerialNumber,
		)
	}
	return

}

func (m Metrics) SetDiskMetrics(ch chan<- prometheus.Metric, c redfish.APIClient, HpeiLORD base.HpeiLOResourceDirectory) {
	// get System URL
	diskURL := findObject(HpeiLORD.Instances, "DiskDrive.", m.server)
	diskAlochym := make(chan base.StorageDisk)
	for _, url := range diskURL {
		go getDisk(c, url, diskAlochym)
	}
	for range diskURL {
		disk := <-diskAlochym
		if disk.SSDEnduranceUtilizationPercentage > 0 {
			ch <- prometheus.MustNewConstMetric(
				base.SysStorageDisk,
				prometheus.GaugeValue,
				(100.0 - disk.SSDEnduranceUtilizationPercentage),
				fmt.Sprintf("%s", disk.Id),
				fmt.Sprintf("%d", disk.CapacityMiB/1000),
				disk.InterfaceType,
				disk.MediaType,
			)
		}
	}
	return
}

func getDisk(c redfish.APIClient, url string, diskChan chan<- base.StorageDisk) {
	diskData, err := c.Get(url)
	// Problem connect to server
	if err != nil {
		return
	}
	var disk base.StorageDisk
	err = json.Unmarshal(diskData, &disk)
	// Data cannot convert StorageDisk struct
	if err != nil {
		return
	}
	diskChan <- disk
}

func (m Metrics) SetEthernetMetrics(ch chan<- prometheus.Metric, c redfish.APIClient, HpeiLORD base.HpeiLOResourceDirectory) {
	// Systems BaseNetworkAdapters
	ifURLs := findObject(HpeiLORD.Instances, "BaseNetworkAdapter.", m.server)
	// // Using goroutine
	// // TODO go routine start
	// Systems Ethernet Interfaces Collection end
	// Systems Ethernet Interfaces start
	ifAlochym := make(chan base.BaseNetworkAdapters)

	// Using go routine
	for _, url := range ifURLs {
		go getEthernet(c, url, ifAlochym)
		// ifData, err := c.Get(url)
		// if err != nil {
		// 	continue
		// }
		// var iface base.BaseNetworkAdapters
		// json.Unmarshal(ifData, &iface)
		// if err != nil {
		// 	continue
		// }
		// for _, v := range iface.PhysicalPorts {
		// 	ch <- prometheus.MustNewConstMetric(
		// 		base.SysEthernetInterface,
		// 		prometheus.GaugeValue,
		// 		v.PortStatus(),
		// 		iface.Id,
		// 		v.MacAddress,
		// 		fmt.Sprintf("%d", v.SpeedMbps),
		// 	)
		// }
	}

	// Get Ethernet Interfaces Data
	for range ifURLs {
		iface := <-ifAlochym
		for _, v := range iface.PhysicalPorts {
			ch <- prometheus.MustNewConstMetric(
				base.SysEthernetInterface,
				prometheus.GaugeValue,
				v.PortStatus(),
				iface.Id,
				v.MacAddress,
				fmt.Sprintf("%d", v.SpeedMbps),
			)
		}
	}
	// TODO go routine end
	return
}

func getEthernet(c redfish.APIClient, url string, ifaceChan chan<- base.BaseNetworkAdapters) {
	ifData, err := c.Get(url)
	if err != nil {
		return
	}
	var iface base.BaseNetworkAdapters
	json.Unmarshal(ifData, &iface)
	if err != nil {
		return
	}
	ifaceChan <- iface
}

func (m Metrics) SetPowerMetrics(ch chan<- prometheus.Metric, c redfish.APIClient, HpeiLORD base.HpeiLOResourceDirectory) {
	// Systems Ethernet Interfaces Collection
	ifPowers := findObject(HpeiLORD.Instances, "Power.", m.server)

	for _, url := range ifPowers {
		powerData, err := c.Get(url)
		if err != nil {
			return
		}
		var power base.PowerControl
		err = json.Unmarshal(powerData, &power)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		for _, v := range power.PowerControl {
			ch <- prometheus.MustNewConstMetric(
				base.ChasPower,
				prometheus.GaugeValue,
				v.PowerConsumedWatts,
			)
		}
	}

	// for range ifPowers {
	// 	var power base.PowerControl
	// 	data := <-ifpower
	// 	err := json.Unmarshal(data, &power)
	// 	// Data cannot convert PowerControl struct
	// 	if err != nil {
	// 		// fmt.Println(ifPowers[0])
	// 		fmt.Println(err.Error())
	// 		return
	// 	}
	// 	for _, v := range power.PowerControl {
	// 		ch <- prometheus.MustNewConstMetric(
	// 			base.ChasPower,
	// 			prometheus.GaugeValue,
	// 			v.PowerConsumedWatts,
	// 		)
	// 	}
	// }
}

func findObject(ob []base.RedfishLinksInstances, obType string, server string) []string {
	var links []string
	for i := range ob {
		if strings.Contains(ob[i].OdataType, obType) && obType == "ComputerSystem." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			links = append(links, url)
			return links
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "Power." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			links = append(links, url)
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "DiskDrive." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			links = append(links, url)
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "BaseNetworkAdapter." {
			url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
			// fmt.Println(url)
			links = append(links, url)
		} else if strings.Contains(ob[i].OdataType, obType) && obType == "EthernetInterface." {
			if strings.Contains(ob[i].ODataID, "Systems") {
				url := fmt.Sprintf("%s%s", server, ob[i].ODataID)
				// fmt.Println(url)
				links = append(links, url)
			}
		}
	}
	return links
}
