package fingerprint

import (
	"fmt"
	"net"
	"strings"

	"github.com/jaypipes/ghw"
	"github.com/lunarhue/libs-go/log"
)

type NetworkInterfaceInfo struct {
	InterfaceName string `json:"interface_name"`
	MacAddress    string `json:"mac_address"`
	Vendor        string `json:"vendor"`
	CurrentIp     string `json:"current_ip"`
}

func GetNetworkInterfaces() ([]NetworkInterfaceInfo, error) {
	netInfo, err := ghw.Network()
	if err != nil {
		return nil, fmt.Errorf("error getting ghw network info: %v", err)
	}

	osInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("error getting OS network interfaces: %v", err)
	}

	var interfaces []NetworkInterfaceInfo

	for _, nic := range netInfo.NICs {
		if nic.IsVirtual {
			continue
		}

		vendor, err := GetVendor(nic.MACAddress)
		if err != nil {
			vendor = "Unknown"
			log.Warnf("Unable to get Vendor for MAC Address: %s, %v", nic.MACAddress, err)
		}

		var currentIp string
		for _, iface := range osInterfaces {
			if strings.EqualFold(iface.HardwareAddr.String(), nic.MACAddress) {
				addrs, err := iface.Addrs()
				if err == nil {
					for _, addr := range addrs {
						if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
							if ipNet.IP.To4() != nil {
								currentIp = ipNet.IP.String()
								break
							}
						}
					}
				}
				break
			}
		}

		interfaces = append(interfaces, NetworkInterfaceInfo{
			InterfaceName: nic.Name,
			MacAddress:    nic.MACAddress,
			Vendor:        vendor,
			CurrentIp:     currentIp,
		})
	}

	return interfaces, nil
}
