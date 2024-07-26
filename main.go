package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/c-robinson/iplib"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "subnet-calc",
		Usage: "calculated subnets",

		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() != 2 {
				return fmt.Errorf("expected 2 arguments, got %d", cCtx.NArg())
			}

			printNetworkAddress(cCtx.Args().Get(0), cCtx.Args().Get(1))

			return nil
		},
		ArgsUsage:   "IP_ADDRESS NETWORK",
		Description: `This tool calculates the network address, usable host IP range, broadcast address, total number of hosts, number of usable hosts, subnet mask, and wildcard mask for a given IP address and network.`,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printNetworkAddress(address, network string) {
	fmt.Printf("Network address: %s\n", networkAddress(address, network))
	fmt.Printf("Usable host IP range: %s\n", usableHostIPRange(address, network))
	fmt.Printf("Broadcast address: %s\n", broadcastAddress(address, network))
	fmt.Printf("Total number of hosts: %d\n", totalNumberOfHosts(address, network))
	fmt.Printf("Number of useable hosts: %d\n", numberOfUsableHosts(address, network))
	fmt.Printf("Subnet mask: %s\n", subnetMask(address, network))
	fmt.Printf("Wildcard mask: %s\n", wildcardMask(address, network))
}

func networkAddress(address, network string) string {
	networkString, err := strconv.Atoi(strings.Replace(network, "/", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	ipa := iplib.NewNet4(net.ParseIP(address), networkString)
	output := ipa.IP().String()
	return output
}

func usableHostIPRange(address, network string) string {
	networkString, err := strconv.Atoi(strings.Replace(network, "/", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	ipa := iplib.NewNet4(net.ParseIP(address), networkString)
	output := fmt.Sprintf("%v - %v", ipa.FirstAddress().String(), ipa.LastAddress().String())
	return output
}

func broadcastAddress(address, network string) string {
	networkString, err := strconv.Atoi(strings.Replace(network, "/", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	ipa := iplib.NewNet4(net.ParseIP(address), networkString)
	output := ipa.BroadcastAddress().String()
	return output
}

func totalNumberOfHosts(address, network string) uint32 {
	networkString, err := strconv.Atoi(strings.Replace(network, "/", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	ipa := iplib.NewNet4(net.ParseIP(address), networkString)
	output := ipa.Count() + 2
	return output
}

func numberOfUsableHosts(address, network string) uint32 {
	networkString, err := strconv.Atoi(strings.Replace(network, "/", "", 1))
	if err != nil {
		log.Fatal(err)
	}
	ipa := iplib.NewNet4(net.ParseIP(address), networkString)
	output := ipa.Count()
	return output
}

func subnetMask(address, network string) string {
	_, ipv4Net, err := net.ParseCIDR(fmt.Sprintf("%s%s", address, network))
	if err != nil {
		log.Fatal(err)
	}
	return ipv4MaskString(ipv4Net.Mask)
}

func ipv4MaskString(m []byte) string {
	if len(m) != 4 {
		panic("ipv4Mask: len must be 4 bytes")
	}

	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}

func wildcardMask(address, network string) string {
	_, ipv4Net, err := net.ParseCIDR(fmt.Sprintf("%s%s", address, network))
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", wildcard(net.IP(ipv4Net.Mask)))
}

// wildcard returns the opposite of the
// the netmask for the network.
func wildcard(mask net.IP) net.IP {
	var ipVal net.IP
	for _, octet := range mask {
		ipVal = append(ipVal, ^octet)
	}
	return ipVal
}
