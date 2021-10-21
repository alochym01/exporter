# exporter
1. reference redfish link for dell server - <https://github.com/dell/iDRAC-Redfish-Scripting>
1. reference redfish link for hpe server - [hpe-redfish-resorucedirectory.json](https://github.com/alochym01/hardware-exporter/blob/master/resourcedirectory.json)
## Review get link for HPE server and DellEMC server
1. review get redfish link for HPE server
    1. https://github.com/alochym01/hardware-exporter/blob/master/domain/server/hpe/exporter.py
    1. https://github.com/alochym01/hardware-exporter/blob/master/domain/server/hpe/redfishObject.py
1. review get redfish link for Dell server
    1. https://api-marketplace.dell.com/#/overview-doc/2978/6818/docs/0WhatsNew.md
## How to code
1. Create route - using router/router.go
1. Create handler for each route
    1. using server/hpe.go HPE server
    1. using server/dell.go DellEMC server
1. Define metrics - using domain/base/metrics.go
1. Add HPE metrics
    1. domain/hpe/metrics.go
1. Add Dell metrics
    1. domain/hpe/metrics.go
1. Interactive with iDrac or iLO - using storage/redfish/client.go
## Using Exporter
1. GIN_MODE=release ./exporter -addr=IP-Address:Port => Production Mode
1. ./exporter => Debug Mode and default PORT=4000
1. curl localhost:4000/metrics/dell?host=https://192.168.2.171
1. curl localhost:4000/metrics/hpe?host=https://192.168.2.156
## Sample metrics
```bash
# HELP ethernet_port ethernet_port {0: LinkUp, 2: LinkDown}
# TYPE ethernet_port gauge
ethernet_port{id="NIC.Embedded.1-1-1",mac="D0:54:66:7A:DC:10",speed="1000"} 0
ethernet_port{id="NIC.Embedded.2-1-1",mac="D0:54:66:7A:DC:11",speed="1000"} 0
ethernet_port{id="NIC.Slot.3-1-1",mac="B4:56:51:33:3A:84",speed="10240"} 0
ethernet_port{id="NIC.Slot.3-2-1",mac="B4:56:51:33:3A:86",speed="10240"} 0
# HELP power_consumed power_consumed of server
# TYPE power_consumed gauge
power_consumed 174
# HELP storage_drive_ssd_endurance storage_drive_ssd_endurance {100: OK, 50: Warning, 20: Critical}
# TYPE storage_drive_ssd_endurance gauge
storage_drive_ssd_endurance{capacity="755",id="0",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="755",id="1",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="755",id="2",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="755",id="3",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="755",id="4",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="755",id="6",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="755",id="7",interface_type="SATA",media_type="SSD"} 100
storage_drive_ssd_endurance{capacity="555",id="5",interface_type="SATA",media_type="SSD"} 100
# HELP system_state system_state {0: OK, 1: Warning, 2: Critical}
# TYPE system_state gauge
system_state{serialnumber="xxxxxxx",sku="xxxxxx"} 0
```
