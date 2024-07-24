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
