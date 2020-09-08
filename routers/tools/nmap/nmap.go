package nmap

import (
	"context"
	"github.com/Ullaakut/nmap"
	"log"
	"strconv"
	"time"
)

func NmapScan(ip string, port string) (resip, resport, resprotocol string) {
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(ip),
		nmap.WithPorts(port),
		nmap.WithContext(ctx),
		nmap.WithSkipHostDiscovery(), // s.args = append(s.args, "-Pn") 加上 -Pn 就不去ping主机，因为有的主机防止ping,增加准确度
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
		return
	}

	result, warnings, err := scanner.Run()
	if err != nil {
		log.Fatalf("Unable to run nmap scan: %v", err)
		return
	}

	if warnings != nil {
		log.Printf("Warnings: \n %v", warnings)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			if port.State.State == "open" {
				if port.Service.Name == "microsoft-ds" {
					port.Service.Name = "SMB"
				}

				b := strconv.Itoa(int(port.ID))
				c := string(b)
				return host.Addresses[0].String(), c , port.Service.Name
			}
			return
		}
		return
	}
	return

}
