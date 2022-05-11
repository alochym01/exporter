package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/alochym01/exporter/v1/metric"
)

func main() {
	// file, _ := ioutil.ReadFile("ComputerSystem.json")
	file, _ := ioutil.ReadFile("Storage.json")
	var chy metric.DellStorage

	err := json.Unmarshal(file, &chy)
	if err != nil {
		fmt.Println(err)
		return
	}
	// b1, _ := json.MarshalIndent(chy, "", "   ")
	// fmt.Println(string(file))
	// fmt.Println(string(b1))
	fmt.Println(chy.Drives[0].ODataID)
	fmt.Println(chy.DrivesCount)

	// fmt.Println(chy.Oem.FansStatus())
	// fmt.Println(chy.Oem.PowerSupplyStatus())
	// fmt.Println(chy.Oem.StorageStatus())
	// fmt.Println(chy.Oem.TemperatureStatus())

}
