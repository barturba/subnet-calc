package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/c-robinson/iplib"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "subnet-calc",
		Usage:       "calculate subnet information",
		ArgsUsage:   "IP_ADDRESS NETWORK",
		Description: "Calculate network address, host range, broadcast address, and subnet information",
		Action: func(cCtx *cli.Context) error {
			if cCtx.NArg() != 2 {
				return fmt.Errorf("expected 2 arguments, got %d", cCtx.NArg())
			}
			if err := printNetworkAddress(cCtx.Args().Get(0), cCtx.Args().Get(1)); err != nil {
				return fmt.Errorf("calculation error: %w", err)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func parseNetwork(address, network string) (iplib.Net4, error) {
	networkString, err := strconv.Atoi(strings.TrimPrefix(network, "/"))
	if err != nil {
		return iplib.Net4{}, fmt.Errorf("invalid network mask: %w", err)
	}

	// Validate network mask range
	if networkString < 0 || networkString > 32 {
		return iplib.Net4{}, fmt.Errorf("invalid network mask: must be between 0 and 32")
	}

	ip := net.ParseIP(address)
	if ip == nil {
		return iplib.Net4{}, fmt.Errorf("invalid IP address")
	}

	return iplib.NewNet4(ip, networkString), nil
}

func printNetworkAddress(address, network string) error {
	net4, err := parseNetwork(address, network)
	if err != nil {
		return err
	}

	_, ipv4Net, err := net.ParseCIDR(fmt.Sprintf("%s%s", address, network))
	if err != nil {
		return err
	}

	fmt.Printf("Network address: %s\n", net4.IP())
	fmt.Printf("Usable host IP range: %s - %s\n", net4.FirstAddress(), net4.LastAddress())
	fmt.Printf("Broadcast address: %s\n", net4.BroadcastAddress())
	fmt.Printf("Total number of hosts: %d\n", net4.Count()+2)
	fmt.Printf("Number of useable hosts: %d\n", net4.Count())
	fmt.Printf("Subnet mask: %s\n", ipv4MaskString(ipv4Net.Mask))
	fmt.Printf("Wildcard mask: %s\n", wildcard(net.IP(ipv4Net.Mask)))

	return nil
}

func ipv4MaskString(m []byte) string {
	if len(m) != 4 {
		return "invalid mask"
	}
	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}

func wildcard(mask net.IP) string {
	ipVal := make(net.IP, len(mask))
	for i, octet := range mask {
		ipVal[i] = ^octet
	}
	return ipVal.String()
}
