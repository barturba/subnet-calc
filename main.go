package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "subnet-calc",
		Usage: "calculated subnets",

		Action: func(cCtx *cli.Context) error {
			printNetworkAddress(cCtx.Args().Get(0), cCtx.Args().Get(1))

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func printNetworkAddress(address, network string) {
	fmt.Printf("printNetworkAddress\n")
	fmt.Printf("Network address: %s\n", networkAddress(address, network))
	fmt.Printf("Usable host IP range: %s\n", usableHostIPRange(address, network))
	fmt.Printf("Broadcast address: %s\n", broadcastAddress(address, network))
	fmt.Printf("Total number of hosts: %d\n", totalNumberOfHosts(address, network))
	fmt.Printf("Subnet mask: %s\n", subnetMask(network))
	fmt.Printf("Wildcard mask: %s\n", wildcardMask(network))
	fmt.Printf("Binary subnet mask: %s\n", binarySubnetMask(network))
	fmt.Printf("IP class: %s\n", ipClass(address))
	fmt.Printf("IP type: %s\n", ipType(address))
	fmt.Printf("Short: %s\n", short(address))
	fmt.Printf("Binary ID: %s\n", binaryID(address))
	fmt.Printf("Integer ID: %d\n", integerID(address))
	fmt.Printf("Hex ID: %s\n", hexID(address))
	fmt.Printf("in-addr.arpa: %s\n", inAddrArpa(address))
	fmt.Printf("ipv4 mapped address: %s\n", ipv4MappedAddress(address))
	fmt.Printf("6to4 address: %s\n", sixToFourAddress(address))
}

func sixToFourAddress(address string) string {
	panic("unimplemented")
}

func binaryID(address string) string {
	panic("unimplemented")
}

func binarySubnetMask(network string) string {
	panic("unimplemented")
}

func ipv4MappedAddress(address string) string {
	panic("unimplemented")
}

func inAddrArpa(address string) string {
	panic("unimplemented")
}

func hexID(address string) string {
	panic("unimplemented")
}

func integerID(address string) int {
	panic("unimplemented")
}

func short(address string) string {
	panic("unimplemented")
}

func ipType(address string) string {
	panic("unimplemented")
}

func ipClass(address string) string {
	panic("unimplemented")
}

func wildcardMask(network string) string {
	panic("unimplemented")
}

func subnetMask(network string) string {
	panic("unimplemented")
}

func totalNumberOfHosts(address, network string) int {
	panic("unimplemented")
}

func broadcastAddress(address, network string) string {
	panic("unimplemented")
}

func usableHostIPRange(address, network string) string {
	panic("unimplemented")
}

func networkAddress(address, network string) string {
	panic("unimplemented")
}
